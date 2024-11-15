package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	_ "embed"

	"github.com/go-resty/resty/v2"
	"gopkg.in/yaml.v3"
)

type YamlSnippet struct {
	Name        string
	Description string
	Prefix      string
	Scope       string
	Body        string
}

//go:embed version.txt
var version []byte

//go:embed help.md
var help []byte

func readYamlFile(yamlFilePath string) (map[string]YamlSnippet, error) {

	yamlFile, err := os.ReadFile(yamlFilePath)

	if err != nil {
		return nil, err
	}

	data := make(map[string]YamlSnippet)

	err = yaml.Unmarshal(yamlFile, &data)

	if err != nil {
		return nil, err
	}
	return data, nil
}

func generateVSCodeSnippets(yamlSnippets map[string]YamlSnippet) ([]byte, error) {
	VSCodeSnippets := make(map[string]interface{})

	for _, snippet := range yamlSnippets {
		body := strings.Split(snippet.Body, "\n")

		VSCodeSnippets[snippet.Name] = map[string]interface{}{
			"prefix":      snippet.Prefix,
			"description": snippet.Description,
			"scope":       snippet.Scope,
			"body":        body,
		}

	}

	bSnippets, err := json.MarshalIndent(VSCodeSnippets, "", "  ")

	if err != nil {
		return nil, err
	}
	return bSnippets, nil
}

func writeJsonFile(jsonFilePath string, jsonData []byte) error {

	f, err := os.Create(jsonFilePath)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(string(jsonData))

	if err != nil {
		return err
	}
	return nil
}

func downloadYamlFile(yamlFileUrl, authHeaderName, authHeaderValue, saveFilePath string) error {
	// authenticationHeader:
	// Example: "PRIVATE-TOKEN: ${GITLAB_WASM_TOKEN}"
	client := resty.New()

	if authHeaderName != "" {
		client.SetHeader(authHeaderName, authHeaderValue)
	}

	resp, err := client.R().
		SetOutput(saveFilePath).
		Get(yamlFileUrl)

	if resp.IsError() {
		return errors.New("error while downloading the yaml file")
	}

	if err != nil {
		return err
	}
	return nil
}

func parse(command string, args []string) error {
	switch command {

	case "generate", "gen":

		flagSet := flag.NewFlagSet("generate", flag.ExitOnError)

		input := flagSet.String("input", "", "input yaml file path")
		output := flagSet.String("output", "", "output json file path")

		yamFileUrl := flagSet.String("url", "", "Url to download the yaml file")
		authHeaderName := flagSet.String("auth-header-name", "", "Authentication header name, ex: PRIVATE-TOKEN")
		authHeaderValue := flagSet.String("auth-header-value", "", "Value of the authentication header, ex: IlovePandas")

		flagSet.Parse(args[0:])

		if *yamFileUrl != "" { // we need to download the yaml file
			fmt.Println("🌍 downloading ", *yamFileUrl, "...")
			err := downloadYamlFile(*yamFileUrl, *authHeaderName, *authHeaderValue, *input)
			if err != nil {
				fmt.Println("😡", err.Error())
				os.Exit(1)
			}
		}

		yamlData, err := readYamlFile(*input)
		if err != nil {
			fmt.Println("😡", err.Error())
			os.Exit(1)
		}
		jsonData, err := generateVSCodeSnippets(yamlData)
		if err != nil {
			fmt.Println("😡", err.Error())
			os.Exit(1)
		}

		err = writeJsonFile(*output, jsonData)

		if err != nil {
			fmt.Println("😡", err.Error())
			os.Exit(1)
		}
		fmt.Println("🙂", *output, "generated")
		//os.Exit(0)
		return nil

	case "version":
		fmt.Println(string(version))
		//os.Exit(0)
		return nil

	case "help":
		fmt.Println(string(help))
		//os.Exit(0)
		return nil

	default:
		return fmt.Errorf("🔴 invalid command")
	}
}

func main() {

	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("🔴 invalid command")
		os.Exit(0)
	}

	command := flag.Args()[0]

	parse(command, flag.Args()[1:])

}
