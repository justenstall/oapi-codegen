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
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/codegen/generator"
)

type NameNormalizerFunction = generator.NameNormalizerFunction

const (
	// NameNormalizerFunctionUnset is the default case, where the `name-normalizer` option hasn't been set. This will use the `ToCamelCase` function.
	//
	// See the docs for `NameNormalizerFunctionToCamelCase` for more details.
	NameNormalizerFunctionUnset = generator.NameNormalizerFunctionUnset
	// NameNormalizerFunctionToCamelCase will use the `ToCamelCase` function.
	//
	// For instance:
	//
	// - `getHttpPet`   => `GetHttpPet`
	// - `OneOf2things` => `OneOf2things`
	NameNormalizerFunctionToCamelCase = generator.NameNormalizerFunctionToCamelCase
	// NameNormalizerFunctionToCamelCaseWithDigits will use the `NameNormalizerFunctionToCamelCaseWithDigits` function.
	//
	// For instance:
	//
	// - `getHttpPet`   => `GetHttpPet`
	// - `OneOf2things` => `OneOf2Things`
	NameNormalizerFunctionToCamelCaseWithDigits = generator.NameNormalizerFunctionToCamelCaseWithDigits
	// NameNormalizerFunctionToCamelCaseWithInitialisms will use the `NameNormalizerFunctionToCamelCaseWithInitialisms` function.
	//
	// For instance:
	//
	// - `getHttpPet`   => `GetHTTPPet`
	// - `OneOf2things` => `OneOf2things`
	NameNormalizerFunctionToCamelCaseWithInitialisms = generator.NameNormalizerFunctionToCamelCaseWithInitialisms
)

// NameNormalizer is a function that takes a type name, and returns that type name converted into a different format.
//
// This may be an Operation ID i.e. `retrieveUserRequests` or a Schema name i.e. `BigBlockOfCheese`
//
// NOTE: this must return a string that can be used as a valid Go type name
type NameNormalizer = generator.NameNormalizer

type NameNormalizerMap = generator.NameNormalizerMap

// NameNormalizers contains the valid options for `NameNormalizerFunction`s that `oapi-codegen` supports.
//
// If you are calling `oapi-codegen` as a library, this allows you to specify your own normalisation types before generating code.
var NameNormalizers = generator.NameNormalizers

// UppercaseFirstCharacter Uppercases the first character in a string. This assumes UTF-8, so we have
// to be careful with unicode, don't treat it as a byte array.
func UppercaseFirstCharacter(str string) string {
	return generator.UppercaseFirstCharacter(str)
}

// Uppercase the first character in a identifier with pkg name. This assumes UTF-8, so we have
// to be careful with unicode, don't treat it as a byte array.
func UppercaseFirstCharacterWithPkgName(str string) string {
	return generator.UppercaseFirstCharacterWithPkgName(str)
}

// LowercaseFirstCharacter Lowercases the first character in a string. This assumes UTF-8, so we have
// to be careful with unicode, don't treat it as a byte array.
func LowercaseFirstCharacter(str string) string {
	return generator.LowercaseFirstCharacter(str)
}

// Lowercase the first upper characters in a string for case of abbreviation.
// This assumes UTF-8, so we have to be careful with unicode, don't treat it as a byte array.
func LowercaseFirstCharacters(str string) string {
	return generator.LowercaseFirstCharacter(str)
}

// ToCamelCase will convert query-arg style strings to CamelCase. We will
// use `., -, +, :, ;, _, ~, ' ', (, ), {, }, [, ]` as valid delimiters for words.
// So, "word.word-word+word:word;word_word~word word(word)word{word}[word]"
// would be converted to WordWordWordWordWordWordWordWordWordWordWordWordWord
func ToCamelCase(str string) string {
	return generator.ToCamelCase(str)
}

// ToCamelCaseWithDigits function will convert query-arg style strings to CamelCase. We will
// use `., -, +, :, ;, _, ~, ' ', (, ), {, }, [, ]` as valid delimiters for words.
// The difference of ToCamelCase that letter after a number becomes capitalized.
// So, "word.word-word+word:word;word_word~word word(word)word{word}[word]3word"
// would be converted to WordWordWordWordWordWordWordWordWordWordWordWordWord3Word
func ToCamelCaseWithDigits(s string) string {
	return generator.ToCamelCaseWithDigits(s)
}

// ToCamelCaseWithInitialisms function will convert query-arg style strings to CamelCase with initialisms in uppercase.
// So, httpOperationId would be converted to HTTPOperationID
func ToCamelCaseWithInitialisms(s string) string {
	return generator.ToCamelCaseWithInitialisms(s)
}

func ToCamelCaseWithInitialism(str string) string {
	return generator.ToCamelCaseWithInitialism(str)
}

// SortedMapKeys takes a map with keys of type string and returns a slice of those
// keys sorted lexicographically.
func SortedMapKeys[T any](m map[string]T) []string {
	return generator.SortedMapKeys(m)
}

// SortedSchemaKeys returns the keys of the given SchemaRef dictionary in sorted
// order, since Golang scrambles dictionary keys. This isn't a generic key sort, because
// we support an extension to grant specific orders to schemas to help control output
// ordering.
func SortedSchemaKeys(dict map[string]*openapi3.SchemaRef) []string {
	return generator.SortedSchemaKeys(dict)
}

// StringInArray checks whether the specified string is present in an array
// of strings
func StringInArray(str string, array []string) bool {
	return generator.StringInArray(str, array)
}

// RefPathToObjName returns the name of referenced object without changes.
//
//	#/components/schemas/Foo -> Foo
//	#/components/parameters/Bar -> Bar
//	#/components/responses/baz_baz -> baz_baz
//	document.json#/Foo -> Foo
//	http://deepmap.com/schemas/document.json#/objObj -> objObj
//
// Does not check refPath correctness.
func RefPathToObjName(refPath string) string {
	return generator.RefPathToObjName(refPath)
}

// RefPathToGoType takes a $ref value and converts it to a Go typename.
// #/components/schemas/Foo -> Foo
// #/components/parameters/Bar -> Bar
// #/components/responses/Baz -> Baz
// Remote components (document.json#/Foo) are supported if they present in --import-mapping
// URL components (http://deepmap.com/schemas/document.json#/Foo) are supported if they present in --import-mapping
// Remote and URL also support standard local paths even though the spec doesn't mention them.
func RefPathToGoType(refPath string) (string, error) {
	return globalState.RefPathToGoType(refPath)
}

// IsGoTypeReference takes a $ref value and checks if it has link to go type.
// #/components/schemas/Foo                     -> true
// ./local/file.yml#/components/parameters/Bar  -> true
// ./local/file.yml                             -> false
// IsGoTypeReference can be used to check whether RefPathToGoType($ref) is possible.
func IsGoTypeReference(ref string) bool {
	return generator.IsGoTypeReference(ref)
}

// IsWholeDocumentReference takes a $ref value and checks if it is whole document reference.
// #/components/schemas/Foo                             -> false
// ./local/file.yml#/components/parameters/Bar          -> false
// ./local/file.yml                                     -> true
// http://deepmap.com/schemas/document.json             -> true
// http://deepmap.com/schemas/document.json#/Foo        -> false
func IsWholeDocumentReference(ref string) bool {
	return generator.IsWholeDocumentReference(ref)
}

// SwaggerUriToIrisUri converts a OpenAPI style path URI with parameters to an
// Iris compatible path URI. We need to replace all of OpenAPI parameters with
//
//	{param}
//	{param*}
//	{.param}
//	{.param*}
//	{;param}
//	{;param*}
//	{?param}
//	{?param*}
func SwaggerUriToIrisUri(uri string) string {
	return generator.SwaggerUriToIrisUri(uri)
}

// SwaggerUriToEchoUri converts a OpenAPI style path URI with parameters to an
// Echo compatible path URI. We need to replace all of OpenAPI parameters with
// ":param". Valid input parameters are:
//
//	{param}
//	{param*}
//	{.param}
//	{.param*}
//	{;param}
//	{;param*}
//	{?param}
//	{?param*}
func SwaggerUriToEchoUri(uri string) string {
	return generator.SwaggerUriToEchoUri(uri)
}

// SwaggerUriToFiberUri converts a OpenAPI style path URI with parameters to a
// Fiber compatible path URI. We need to replace all of OpenAPI parameters with
// ":param". Valid input parameters are:
//
//	{param}
//	{param*}
//	{.param}
//	{.param*}
//	{;param}
//	{;param*}
//	{?param}
//	{?param*}
func SwaggerUriToFiberUri(uri string) string {
	return generator.SwaggerUriToFiberUri(uri)
}

// SwaggerUriToChiUri converts a swagger style path URI with parameters to a
// Chi compatible path URI. We need to replace all Swagger parameters with
// "{param}". Valid input parameters are:
//
//	{param}
//	{param*}
//	{.param}
//	{.param*}
//	{;param}
//	{;param*}
//	{?param}
//	{?param*}
func SwaggerUriToChiUri(uri string) string {
	return generator.SwaggerUriToChiUri(uri)
}

// SwaggerUriToGinUri converts a swagger style path URI with parameters to a
// Gin compatible path URI. We need to replace all Swagger parameters with
// ":param". Valid input parameters are:
//
//	{param}
//	{param*}
//	{.param}
//	{.param*}
//	{;param}
//	{;param*}
//	{?param}
//	{?param*}
func SwaggerUriToGinUri(uri string) string {
	return generator.SwaggerUriToGinUri(uri)
}

// SwaggerUriToGorillaUri converts a swagger style path URI with parameters to a
// Gorilla compatible path URI. We need to replace all Swagger parameters with
// ":param". Valid input parameters are:
//
//	{param}
//	{param*}
//	{.param}
//	{.param*}
//	{;param}
//	{;param*}
//	{?param}
//	{?param*}
func SwaggerUriToGorillaUri(uri string) string {
	return generator.SwaggerUriToGorillaUri(uri)
}

// SwaggerUriToStdHttpUri converts a swagger style path URI with parameters to a
// Chi compatible path URI. We need to replace all Swagger parameters with
// "{param}". Valid input parameters are:
//
//	{param}
//	{param*}
//	{.param}
//	{.param*}
//	{;param}
//	{;param*}
//	{?param}
//	{?param*}
func SwaggerUriToStdHttpUri(uri string) string {
	return generator.SwaggerUriToStdHttpUri(uri)
}

// OrderedParamsFromUri returns the argument names, in order, in a given URI string, so for
// /path/{param1}/{.param2*}/{?param3}, it would return param1, param2, param3
func OrderedParamsFromUri(uri string) []string {
	return generator.OrderedParamsFromUri(uri)
}

// ReplacePathParamsWithStr replaces path parameters of the form {param} with %s
func ReplacePathParamsWithStr(uri string) string {
	return generator.ReplacePathParamsWithStr(uri)
}

// SortParamsByPath reorders the given parameter definitions to match those in the path URI.
func SortParamsByPath(path string, in []generator.ParameterDefinition) ([]generator.ParameterDefinition, error) {
	return generator.SortParamsByPath(path, in)
}

// IsGoKeyword returns whether the given string is a go keyword
func IsGoKeyword(str string) bool {
	return generator.IsGoKeyword(str)
}

// IsPredeclaredGoIdentifier returns whether the given string
// is a predefined go identifier.
//
// See https://golang.org/ref/spec#Predeclared_identifiers
func IsPredeclaredGoIdentifier(str string) bool {
	return generator.IsPredeclaredGoIdentifier(str)
}

// IsGoIdentity checks if the given string can be used as an identity
// in the generated code like a type name or constant name.
//
// See https://golang.org/ref/spec#Identifiers
func IsGoIdentity(str string) bool {
	return generator.IsGoIdentity(str)
}

// IsValidGoIdentity checks if the given string can be used as a
// name of variable, constant, or type.
func IsValidGoIdentity(str string) bool {
	return generator.IsValidGoIdentity(str)
}

// SanitizeGoIdentity deletes and replaces the illegal runes in the given
// string to use the string as a valid identity.
func SanitizeGoIdentity(str string) string {
	return generator.SanitizeGoIdentity(str)
}

// SanitizeEnumNames fixes illegal chars in the enum names
// and removes duplicates
func SanitizeEnumNames(enumNames, enumValues []string) map[string]string {
	return globalState.SanitizeEnumNames(enumNames, enumValues)
}

// SchemaNameToTypeName converts a Schema name to a valid Go type name. It converts to camel case, and makes sure the name is
// valid in Go
func SchemaNameToTypeName(name string) string {
	return generator.SchemaNameToTypeName(name, globalState.NameNormalizer)
}

// According to the spec, additionalProperties may be true, false, or a
// schema. If not present, true is implied. If it's a schema, true is implied.
// If it's false, no additional properties are allowed. We're going to act a little
// differently, in that if you want additionalProperties code to be generated,
// you must specify an additionalProperties type
// If additionalProperties it true/false, this field will be non-nil.
func SchemaHasAdditionalProperties(schema *openapi3.Schema) bool {
	return generator.SchemaHasAdditionalProperties(schema)
}

// PathToTypeName converts a path, like Object/field1/nestedField into a go
// type name.
func PathToTypeName(path []string) string {
	return generator.PathToTypeName(path, globalState.NameNormalizer)
}

// StringToGoComment renders a possible multi-line string as a valid Go-Comment.
// Each line is prefixed as a comment.
func StringToGoComment(in string) string {
	return generator.StringToGoComment(in)
}

// StringWithTypeNameToGoComment renders a possible multi-line string as a
// valid Go-Comment, including the name of the type being referenced. Each line
// is prefixed as a comment.
func StringWithTypeNameToGoComment(in, typeName string) string {
	return generator.StringWithTypeNameToGoComment(in, typeName)
}

func DeprecationComment(reason string) string {
	return generator.DeprecationComment(reason)
}

// EscapePathElements breaks apart a path, and looks at each element. If it's
// not a path parameter, eg, {param}, it will URL-escape the element.
func EscapePathElements(path string) string {
	return generator.EscapePathElements(path)
}

func ParseGoImportExtension(v *openapi3.SchemaRef) (*generator.GoImport, error) {
	return generator.ParseGoImportExtension(v)
}

func MergeImports(dst, src map[string]generator.GoImport) {
	generator.MergeImports(dst, src)
}

// TypeDefinitionsEquivalent checks for equality between two type definitions, but
// not every field is considered. We only want to know if they are fundamentally
// the same type.
func TypeDefinitionsEquivalent(t1, t2 generator.TypeDefinition) bool {
	return generator.TypeDefinitionsEquivalent(t1, t2)
}
