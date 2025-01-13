// Copyright 2019 DeepMap, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codegen

import (
	"embed"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/oapi-codegen/oapi-codegen/v2/pkg/codegen/generator"
)

// globalState stores all global state. Please don't put global state anywhere
// else so that we can easily track it.
var globalState generator.State

// Generate uses the Go templating engine to generate all of our server wrappers from
// the descriptions we've built up above from the schema objects.
// opts defines
func Generate(spec *openapi3.T, opts Configuration) (string, error) {
	state, err := generator.NewGenerator(spec, opts)
	if err != nil {
		return "", err
	}
	globalState = *state
	return globalState.Generate()
}

func GenerateTypeDefinitions(t *template.Template, swagger *openapi3.T, ops []generator.OperationDefinition, excludeSchemas []string) (string, error) {
	return globalState.GenerateTypeDefinitions(t, swagger, ops, excludeSchemas)
}

// GenerateConstants generates operation ids, context keys, paths, etc. to be exported as constants
func GenerateConstants(t *template.Template, ops []generator.OperationDefinition) (string, error) {
	return generator.GenerateConstants(t, ops)
}

// GenerateTypesForSchemas generates type definitions for any custom types defined in the
// components/schemas section of the Swagger spec.
func GenerateTypesForSchemas(t *template.Template, schemas map[string]*openapi3.SchemaRef, excludeSchemas []string) ([]generator.TypeDefinition, error) {
	return globalState.GenerateTypesForSchemas(t, schemas, excludeSchemas)
}

// GenerateTypesForParameters generates type definitions for any custom types defined in the
// components/parameters section of the Swagger spec.
func GenerateTypesForParameters(t *template.Template, params map[string]*openapi3.ParameterRef) ([]generator.TypeDefinition, error) {
	return globalState.GenerateTypesForParameters(t, params)
}

// GenerateTypesForResponses generates type definitions for any custom types defined in the
// components/responses section of the Swagger spec.
func GenerateTypesForResponses(t *template.Template, responses openapi3.ResponseBodies) ([]generator.TypeDefinition, error) {
	return globalState.GenerateTypesForResponses(t, responses)
}

// GenerateTypesForRequestBodies generates type definitions for any custom types defined in the
// components/requestBodies section of the Swagger spec.
func GenerateTypesForRequestBodies(t *template.Template, bodies map[string]*openapi3.RequestBodyRef) ([]generator.TypeDefinition, error) {
	return globalState.GenerateTypesForRequestBodies(t, bodies)
}

// GenerateTypes passes a bunch of types to the template engine, and buffers
// its output into a string.
func GenerateTypes(t *template.Template, types []generator.TypeDefinition) (string, error) {
	return generator.GenerateTypes(t, types)
}

func GenerateEnums(t *template.Template, types []generator.TypeDefinition) (string, error) {
	return globalState.GenerateEnums(t, types)
}

// GenerateImports generates our import statements and package definition.
func GenerateImports(t *template.Template, externalImports []string, packageName string, versionOverride *string) (string, error) {
	return globalState.GenerateImports(t, externalImports, packageName, versionOverride)
}

// GenerateAdditionalPropertyBoilerplate generates all the glue code which provides
// the API for interacting with additional properties and JSON-ification
func GenerateAdditionalPropertyBoilerplate(t *template.Template, typeDefs []generator.TypeDefinition) (string, error) {
	return generator.GenerateAdditionalPropertyBoilerplate(t, typeDefs)
}

func GenerateUnionBoilerplate(t *template.Template, typeDefs []generator.TypeDefinition) (string, error) {
	return generator.GenerateUnionBoilerplate(t, typeDefs)
}

func GenerateUnionAndAdditionalProopertiesBoilerplate(t *template.Template, typeDefs []generator.TypeDefinition) (string, error) {
	return generator.GenerateUnionAndAdditionalProopertiesBoilerplate(t, typeDefs)
}

// SanitizeCode runs sanitizers across the generated Go code to ensure the
// generated code will be able to compile.
func SanitizeCode(goCode string) string {
	return generator.SanitizeCode(goCode)
}

// GetUserTemplateText attempts to retrieve the template text from a passed in URL or file
// path when inputData is more than one line.
// This function will attempt to load a file first, and if it fails, will try to get the
// data from the remote endpoint.
// The timeout for remote download file is 30 seconds.
func GetUserTemplateText(inputData string) (template string, err error) {
	return generator.GetUserTemplateText(inputData)
}

// LoadTemplates loads all of our template files into a text/template. The
// path of template is relative to the templates directory.
func LoadTemplates(src embed.FS, t *template.Template) error {
	return generator.LoadTemplates(src, t)
}

func OperationSchemaImports(s *generator.Schema) (map[string]generator.GoImport, error) {
	return generator.OperationSchemaImports(s)
}

func OperationImports(ops []generator.OperationDefinition) (map[string]generator.GoImport, error) {
	return generator.OperationImports(ops)
}

func GetTypeDefinitionsImports(swagger *openapi3.T, excludeSchemas []string) (map[string]generator.GoImport, error) {
	return generator.GetTypeDefinitionsImports(swagger, excludeSchemas)
}

func GoSchemaImports(schemas ...*openapi3.SchemaRef) (map[string]generator.GoImport, error) {
	return generator.GoSchemaImports(schemas...)
}

func GetSchemaImports(schemas map[string]*openapi3.SchemaRef, excludeSchemas []string) (map[string]generator.GoImport, error) {
	return generator.GetSchemaImports(schemas, excludeSchemas)
}

func GetRequestBodiesImports(bodies map[string]*openapi3.RequestBodyRef) (map[string]generator.GoImport, error) {
	return generator.GetRequestBodiesImports(bodies)
}

func GetResponsesImports(responses map[string]*openapi3.ResponseRef) (map[string]generator.GoImport, error) {
	return generator.GetResponsesImports(responses)
}

func GetParametersImports(params map[string]*openapi3.ParameterRef) (map[string]generator.GoImport, error) {
	return generator.GetParametersImports(params)
}

func SetGlobalStateSpec(spec *openapi3.T) {
	globalState.SetSpec(spec)
}

func SetGlobalStateOptions(opts Configuration) {
	_ = globalState.SetOptions(opts)
}
