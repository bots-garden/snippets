package internal

import (
	"encoding/json"
	"errors"
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

func ReadYamlFile(yamlFilePath string) (map[string]YamlSnippet, error) {

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

func GenerateVSCodeSnippets(yamlSnippets map[string]YamlSnippet) ([]byte, error) {
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

func WriteJsonFile(jsonFilePath string, jsonData []byte) error {

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

func DownloadYamlFile(yamlFileUrl, authHeaderName, authHeaderValue, saveFilePath string) error {
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
