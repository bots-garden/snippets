package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"snippet/internal"
)

//go:embed version.txt
var version []byte

//go:embed help.md
var help []byte

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
			err := internal.DownloadYamlFile(*yamFileUrl, *authHeaderName, *authHeaderValue, *input)
			if err != nil {
				fmt.Println("😡", err.Error())
				os.Exit(1)
			}
		}

		yamlData, err := internal.ReadYamlFile(*input)
		if err != nil {
			fmt.Println("😡", err.Error())
			os.Exit(1)
		}
		jsonData, err := internal.GenerateVSCodeSnippets(yamlData)
		if err != nil {
			fmt.Println("😡", err.Error())
			os.Exit(1)
		}

		err = internal.WriteJsonFile(*output, jsonData)

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
		fmt.Println(string(help))
		fmt.Println("🔴 invalid command")
		os.Exit(0)
	}

	command := flag.Args()[0]

	parse(command, flag.Args()[1:])

}
