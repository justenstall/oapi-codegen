package global

import (
	"errors"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/codegen"
)

var (
	// globalState stores all global state. Please don't put global state anywhere
	// else so that we can easily track it.
	globalState codegen.State

	// globalStateError stores any errors from producing the global state.
	// TODO: where to surface this?
	globalStateError error
)

// Generate uses the Go templating engine to generate all of our server wrappers from
// the descriptions we've built up above from the schema objects.
// opts defines
func Generate(spec *openapi3.T, opts codegen.Configuration) (string, error) {
	state, err := codegen.NewGenerator(spec, opts)
	if err != nil {
		globalStateError = err
		return "", err
	}
	globalState = *state
	return globalState.Generate()
}

func SetGlobalStateSpec(spec *openapi3.T) {
	globalStateError = errors.Join(globalStateError,
		globalState.SetSpec(spec))
}

func SetGlobalStateOptions(opts codegen.Configuration) {
	globalStateError = errors.Join(globalStateError,
		globalState.SetOptions(opts))
}

var (
	// codegen.go
	// Generate                      = globalState.Generate
	GenerateTypeDefinitions       = globalState.GenerateTypeDefinitions
	GenerateTypesForSchemas       = globalState.GenerateTypesForSchemas
	GenerateTypesForParameters    = globalState.GenerateTypesForParameters
	GenerateTypesForResponses     = globalState.GenerateTypesForResponses
	GenerateTypesForRequestBodies = globalState.GenerateTypesForRequestBodies
	GenerateEnums                 = globalState.GenerateEnums
	GenerateImports               = globalState.GenerateImports

	// merge_schemas_v1.go
	GenStructFromAllOf = globalState.GenStructFromAllOf

	// merge_schemas.go
	MergeSchemas = globalState.MergeSchemas

	// operations.go
	DescribeParameters           = globalState.DescribeParameters
	OperationDefinitions         = globalState.OperationDefinitions
	GenerateBodyDefinitions      = globalState.GenerateBodyDefinitions
	GenerateResponseDefinitions  = globalState.GenerateResponseDefinitions
	GenerateTypeDefsForOperation = globalState.GenerateTypeDefsForOperation
	GenerateParamsTypes          = globalState.GenerateParamsTypes

	// schema.go
	GenerateGoSchema        = globalState.GenerateGoSchema
	GenFieldsFromProperties = globalState.GenFieldsFromProperties
	GenStructFromSchema     = globalState.GenStructFromSchema

	// template_helpers.go
	TemplateFunctions = globalState.TemplateFunctions

	// utils.go
	RefPathToGoType      = globalState.RefPathToGoType
	SanitizeEnumNames    = globalState.SanitizeEnumNames
	SchemaNameToTypeName = globalState.SchemaNameToTypeName
	PathToTypeName       = globalState.PathToTypeName
)
