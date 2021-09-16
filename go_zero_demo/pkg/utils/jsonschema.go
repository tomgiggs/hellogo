package utils

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

func JsonschemaValid(schemaContent,jsonContent string)(valid bool,err error) {
	loader1 := gojsonschema.NewStringLoader(schemaContent)
	schema, err := gojsonschema.NewSchema(loader1)
	if err != nil {
		return false,err
	}

	documentLoader := gojsonschema.NewStringLoader(jsonContent)
	result, err := schema.Validate(documentLoader)
	if err != nil {
		return false,err
	}

	if result.Valid() {
		return true,nil
	} else {
		errInfo := ""
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			errInfo += desc.String()+"\n"
		}
		return false,fmt.Errorf(errInfo)
	}
	return true,nil
}
