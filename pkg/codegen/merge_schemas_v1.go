package codegen

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// GenStructFromAllOf generates an object that is the union of the objects in the
// input array. In the case of Ref objects, we use an embedded struct, otherwise,
// we inline the fields.
func GenStructFromAllOf(allOf []*openapi3.SchemaRef, path []string) (string, error) {
	return globalState.GenStructFromAllOf(allOf, path)
}
