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
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/oapi-codegen/oapi-codegen/v2/pkg/codegen/generator"
)

// DescribeParameters walks the given parameters dictionary, and generates the above
// descriptors into a flat list. This makes it a lot easier to traverse the
// data in the template engine.
func DescribeParameters(params openapi3.Parameters, path []string) ([]generator.ParameterDefinition, error) {
	return globalState.DescribeParameters(params, path)
}

func DescribeSecurityDefinition(securityRequirements openapi3.SecurityRequirements) []generator.SecurityDefinition {
	return generator.DescribeSecurityDefinition(securityRequirements)
}

// FilterParameterDefinitionByType returns the subset of the specified parameters which are of the
// specified type.
func FilterParameterDefinitionByType(params []generator.ParameterDefinition, in string) []generator.ParameterDefinition {
	return generator.FilterParameterDefinitionByType(params, in)
}

// OperationDefinitions returns all operations for a swagger definition.
func OperationDefinitions(swagger *openapi3.T, initialismOverrides bool) ([]generator.OperationDefinition, error) {
	return globalState.OperationDefinitions(swagger, initialismOverrides)
}

// GenerateBodyDefinitions turns the Swagger body definitions into a list of our body
// definitions which will be used for code generation.
func GenerateBodyDefinitions(operationID string, bodyOrRef *openapi3.RequestBodyRef) ([]generator.RequestBodyDefinition, []generator.TypeDefinition, error) {
	return globalState.GenerateBodyDefinitions(operationID, bodyOrRef)
}

func GenerateResponseDefinitions(operationID string, responses map[string]*openapi3.ResponseRef) ([]generator.ResponseDefinition, error) {
	return globalState.GenerateResponseDefinitions(operationID, responses)
}

func GenerateTypeDefsForOperation(op generator.OperationDefinition) []generator.TypeDefinition {
	return globalState.GenerateTypeDefsForOperation(op)
}

// GenerateParamsTypes defines the schema for a parameters definition object
// which encapsulates all the query, header and cookie parameters for an operation.
func GenerateParamsTypes(op generator.OperationDefinition) []generator.TypeDefinition {
	return globalState.GenerateParamsTypes(op)
}

// GenerateTypesForOperations generates code for all types produced within operations
func GenerateTypesForOperations(t *template.Template, ops []generator.OperationDefinition) (string, error) {
	return generator.GenerateTypesForOperations(t, ops)
}

// GenerateIrisServer generates all the go code for the ServerInterface as well as
// all the wrapper functions around our handlers.
func GenerateIrisServer(t *template.Template, operations []generator.OperationDefinition) (string, error) {
	return generator.GenerateIrisServer(t, operations)
}

// GenerateChiServer generates all the go code for the ServerInterface as well as
// all the wrapper functions around our handlers.
func GenerateChiServer(t *template.Template, operations []generator.OperationDefinition) (string, error) {
	return generator.GenerateChiServer(t, operations)
}

// GenerateFiberServer generates all the go code for the ServerInterface as well as
// all the wrapper functions around our handlers.
func GenerateFiberServer(t *template.Template, operations []generator.OperationDefinition) (string, error) {
	return generator.GenerateFiberServer(t, operations)
}

// GenerateEchoServer generates all the go code for the ServerInterface as well as
// all the wrapper functions around our handlers.
func GenerateEchoServer(t *template.Template, operations []generator.OperationDefinition) (string, error) {
	return generator.GenerateEchoServer(t, operations)
}

// GenerateGinServer generates all the go code for the ServerInterface as well as
// all the wrapper functions around our handlers.
func GenerateGinServer(t *template.Template, operations []generator.OperationDefinition) (string, error) {
	return generator.GenerateGinServer(t, operations)
}

// GenerateGorillaServer generates all the go code for the ServerInterface as well as
// all the wrapper functions around our handlers.
func GenerateGorillaServer(t *template.Template, operations []generator.OperationDefinition) (string, error) {
	return generator.GenerateGorillaServer(t, operations)
}

// GenerateStdHTTPServer generates all the go code for the ServerInterface as well as
// all the wrapper functions around our handlers.
func GenerateStdHTTPServer(t *template.Template, operations []generator.OperationDefinition) (string, error) {
	return generator.GenerateStdHTTPServer(t, operations)
}

func GenerateStrictServer(t *template.Template, operations []generator.OperationDefinition, opts Configuration) (string, error) {
	return generator.GenerateStrictServer(t, operations, opts)
}

func GenerateStrictResponses(t *template.Template, responses []generator.ResponseDefinition) (string, error) {
	return generator.GenerateStrictResponses(t, responses)
}

// GenerateClient uses the template engine to generate the function which registers our wrappers
// as Echo path handlers.
func GenerateClient(t *template.Template, ops []generator.OperationDefinition) (string, error) {
	return generator.GenerateClient(t, ops)
}

// GenerateClientWithResponses generates a client which extends the basic client which does response
// unmarshaling.
func GenerateClientWithResponses(t *template.Template, ops []generator.OperationDefinition) (string, error) {
	return generator.GenerateClientWithResponses(t, ops)
}

// GenerateTemplates used to generate templates
func GenerateTemplates(templates []string, t *template.Template, ops interface{}) (string, error) {
	return generator.GenerateTemplates(templates, t, ops)
}

// CombineOperationParameters combines the Parameters defined at a global level (Parameters defined for all methods on a given path) with the Parameters defined at a local level (Parameters defined for a specific path), preferring the locally defined parameter over the global one
func CombineOperationParameters(globalParams []generator.ParameterDefinition, localParams []generator.ParameterDefinition) ([]generator.ParameterDefinition, error) {
	return generator.CombineOperationParameters(globalParams, localParams)
}
