package main

import (
	"os"
	"testing"
)

func TestGenerateVSCodeSnippets(t *testing.T) {
	yamlFilePath := "samples/python01.yml"
	yamlData, err := readYamlFile(yamlFilePath)
	if err != nil {
		t.Errorf("failed to read YAML file: %v", err)
		return
	}

	jsonSnippets, err := generateVSCodeSnippets(yamlData)
	if err != nil {
		t.Errorf("failed to generate JSON snippets: %v", err)
		return
	}

	codeSnippetsFilePath := "samples/python01.code-snippets"
	expectedJsonSnippets, err := os.ReadFile(codeSnippetsFilePath)
	if err != nil {
		t.Errorf("failed to read code snippets file: %v", err)
		return
	}

	computedSnippet := string(jsonSnippets)
	expectedSnippet := string(expectedJsonSnippets)
	if computedSnippet != expectedSnippet {
		t.Errorf("generated JSON snippets do not match expected output: \n%v\nvs\n%v", computedSnippet, expectedSnippet)
	}
}
