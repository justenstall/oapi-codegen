package codegen

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/codegen/generator"
)

func PropertiesEqual(a, b generator.Property) bool {
	return generator.PropertiesEqual(a, b)
}

func GenerateGoSchema(sref *openapi3.SchemaRef, path []string) (generator.Schema, error) {
	return globalState.GenerateGoSchema(sref, path)
}

// GenFieldsFromProperties produce corresponding field names with JSON annotations,
// given a list of schema descriptors
func GenFieldsFromProperties(props []generator.Property) []string {
	return globalState.GenFieldsFromProperties(props)
}

func GenStructFromSchema(schema generator.Schema) string {
	return globalState.GenStructFromSchema(schema)
}
