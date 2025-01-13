package codegen

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/codegen/generator"
)

// MergeSchemas merges all the fields in the schemas supplied into one giant schema.
// The idea is that we merge all fields together into one schema.
func MergeSchemas(allOf []*openapi3.SchemaRef, path []string) (generator.Schema, error) {
	return globalState.MergeSchemas(allOf, path)
}
