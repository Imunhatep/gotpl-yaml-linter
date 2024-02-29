package main

import (
	"flag"
	"fmt"
	"github.com/imunhatep/gotpl-yaml-linter/internal"
	yaml "gopkg.in/yaml.v3"
	"os"
)

func toYaml(data interface{}) string {
	b, err := yaml.Marshal(data)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return ""
	}

	return string(b)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <path_to_gotpl_file>")
		return
	}

	root := flag.String("path", ".", "path for listing")
	filter := flag.String("filter", "*.yaml", "pattern for matching in directories")
	format := flag.Bool("fmt", false, "Format files inplace")
	flag.Parse()

	fileList, err := internal.ListFilesInDir(root, filter)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	fmt.Printf("Found files: \n%s\n\n", toYaml(fileList))

	// Print matched files
	allTrue := true
	for _, match := range fileList {
		valid, err := internal.FormatFilesInPath(match, *format)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}

		allTrue = allTrue && valid
	}

	if !allTrue {
		os.Exit(1)
	}

	os.Exit(0)
}
