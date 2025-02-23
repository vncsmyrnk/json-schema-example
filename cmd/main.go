package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xeipuuv/gojsonschema"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	schemaPath := filepath.Join(cwd, "schemas/schema.json")
	documentPath := filepath.Join(cwd, "documents/example.json")

	schemaLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://%s", schemaPath))
	documentLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://%s", documentPath))

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}
