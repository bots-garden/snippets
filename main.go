package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	_ "embed"

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

func generateJsonSnippets(yamlSnippets map[string]YamlSnippet) ([]byte, error) {
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

func parse(command string, args []string) error {
	switch command {

	case "generate", "gen":

		flagSet := flag.NewFlagSet("generate", flag.ExitOnError)

		input := flagSet.String("input", "", "input yaml file path")
		output := flagSet.String("output", "", "output json file path")

		flagSet.Parse(args[0:])

		yamlData, err := readYamlFile(*input)
		if err != nil {
			fmt.Println("😡", err.Error())
			os.Exit(1)
		}
		jsonData, err := generateJsonSnippets(yamlData)
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
