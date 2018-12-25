package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/ast"
)

type bnetField struct {
	name  string
	items []string
}

// mapFieldToBnetField maps supplied field to a field that is required for the supplied field to be fetched
func mapFieldToBnetField(field string) string {
	bnetFields := []bnetField{
		bnetField{
			name:  "items",
			items: []string{"averageItemLevel", "averageItemLevelEquipped", "items"},
		},
	}

	for _, bnetField := range bnetFields {
		if bnetField.name == "items" {

			for _, item := range bnetField.items {
				if item == field {
					return bnetField.name
				}
			}

		}
	}

	return field
}

func getFieldsFromSelectionSet(selectionSet ast.SelectionSet) []ast.Field {
	fields := []ast.Field{}

	for _, selection := range selectionSet {
		switch sel := selection.(type) {
		case *ast.Field:
			fields = append(fields, *sel)
		}
	}

	return fields
}

func fieldNamesToStrings(fields []ast.Field) []string {
	strings := []string{}

	for _, field := range fields {
		strings = append(strings, field.Name)
	}

	return strings
}

func MapFieldsToBnetFields(ctx context.Context) []string {
	requestContext := graphql.GetRequestContext(ctx)

	stringFields := []string{}

	for _, operation := range requestContext.Doc.Operations {
		// character
		queryFields := getFieldsFromSelectionSet(operation.SelectionSet)
		// { id name ... }
		for _, qf := range queryFields {
			typeFields := getFieldsFromSelectionSet(qf.SelectionSet)
			stringFields = fieldNamesToStrings(typeFields)
		}
	}

	bnetFields := []string{}

	for _, field := range stringFields {
		bnf := mapFieldToBnetField(field)
		if !stringContains(bnetFields, bnf) {
			bnetFields = append(bnetFields, bnf)
		}
	}

	return bnetFields
}

// WriteResponseBodyToJSONFile writes to a JSON file in the project root
func WriteResponseBodyToJSONFile(body []byte, filename string) {
	var prettyJSON bytes.Buffer

	err := json.Indent(&prettyJSON, body, "", "\t")
	if err != nil {
		log.Println("PrettyPrint error:", err)
	}

	projectPath := "/home/smappa/go/src/github.com/zrcni/go-bnet-graphql-api"

	// if err := os.Mkdir(fmt.Sprintf("%s/example-responses/%s", projectPath, filename), 0644); err != nil {
	// 	log.Println(err)
	// }

	// fname := fmt.Sprintf("%s/example-responses/%s/response.json", projectPath, filename)
	fname := fmt.Sprintf("%s/example-responses/%s.json", projectPath, filename)
	er := ioutil.WriteFile(fname, prettyJSON.Bytes(), 0644)

	if er != nil {
		log.Println(er)
	}
}

func stringContains(slice []string, element string) bool {
	for _, el := range slice {
		if el == element {
			return true
		}
	}
	return false
}
