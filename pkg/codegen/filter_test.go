package codegen

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/assert"
)

func TestFilterOperationsByTag(t *testing.T) {
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true

	// Get a spec from the test definition in this file:
	swagger, err := loader.LoadFromData([]byte(testOpenAPIDefinition))
	assert.NoError(t, err)

	packageName := "testswagger"
	t.Run("include tags", func(t *testing.T) {
		opts := Configuration{
			PackageName: packageName,
			Generate: GenerateOptions{
				EchoServer:   true,
				Client:       true,
				Models:       true,
				EmbeddedSpec: true,
			},
			OutputOptions: OutputOptions{
				IncludeTags: []string{"hippo", "giraffe", "cat"},
			},
		}

		loader := openapi3.NewLoader()
		loader.IsExternalRefsAllowed = true

		// Get a spec from the test definition in this file:
		swagger, err := loader.LoadFromData([]byte(testOpenAPIDefinition))
		assert.NoError(t, err)

		// Run our code generation:
		gen, err := NewGenerator(swagger, opts)
		assert.NoError(t, err)
		assert.NotEmpty(t, gen)
		code, err := gen.Generate()
		assert.NoError(t, err)
		assert.NotEmpty(t, code)
		assert.NotContains(t, code, `"/test/:name"`)
		assert.Contains(t, code, `"/cat"`)
	})

	t.Run("exclude tags", func(t *testing.T) {
		opts := Configuration{
			PackageName: packageName,
			Generate: GenerateOptions{
				EchoServer:   true,
				Client:       true,
				Models:       true,
				EmbeddedSpec: true,
			},
			OutputOptions: OutputOptions{
				ExcludeTags: []string{"hippo", "giraffe", "cat"},
			},
		}

		// Run our code generation:
		gen, err := NewGenerator(swagger, opts)
		assert.NoError(t, err)
		assert.NotEmpty(t, gen)
		code, err := gen.Generate()
		assert.NoError(t, err)
		assert.NotEmpty(t, code)
		assert.Contains(t, code, `"/test/:name"`)
		assert.NotContains(t, code, `"/cat"`)
	})
}

func TestFilterOperationsByOperationID(t *testing.T) {
	t.Run("include operation ids", func(t *testing.T) {
		loader := openapi3.NewLoader()
		loader.IsExternalRefsAllowed = true

		// Get a spec from the test definition in this file:
		swagger, err := loader.LoadFromData([]byte(testOpenAPIDefinition))
		assert.NoError(t, err)

		packageName := "testswagger"

		opts := Configuration{
			PackageName: packageName,
			Generate: GenerateOptions{
				StdHTTPServer: true,
				Client:        true,
				Models:        true,
				EmbeddedSpec:  true,
			},
			Compatibility: CompatibilityOptions{
				PreserveOriginalOperationIdCasingInEmbeddedSpec: true,
			},
			OutputOptions: OutputOptions{
				IncludeOperationIDs: []string{"getCatStatus"},
			},
		}

		filterOperationsByOperationID(swagger, opts)

		testPath := swagger.Paths.Value("/test/{name}")
		if assert.NotNil(t, testPath) {
			assert.Nil(t, testPath.Get)
		}
		catPath := swagger.Paths.Value("/cat")
		if assert.NotNil(t, catPath, "/cat") {
			assert.NotNil(t, catPath.Get)
		}
	})

	t.Run("exclude operation ids", func(t *testing.T) {
		loader := openapi3.NewLoader()
		loader.IsExternalRefsAllowed = true

		// Get a spec from the test definition in this file:
		swagger, err := loader.LoadFromData([]byte(testOpenAPIDefinition))
		assert.NoError(t, err)

		packageName := "testswagger"

		opts := Configuration{
			PackageName: packageName,
			Generate: GenerateOptions{
				EchoServer:   true,
				Client:       true,
				Models:       true,
				EmbeddedSpec: true,
			},
			Compatibility: CompatibilityOptions{
				PreserveOriginalOperationIdCasingInEmbeddedSpec: true,
			},
			OutputOptions: OutputOptions{
				ExcludeOperationIDs: []string{"getCatStatus"},
			},
		}

		filterOperationsByOperationID(swagger, opts)

		testPath := swagger.Paths.Value("/test/{name}")
		if assert.NotNil(t, testPath) {
			assert.NotNil(t, testPath.Get)
		}
		catPath := swagger.Paths.Value("/cat")
		if assert.NotNil(t, catPath, "/cat") {
			assert.Nil(t, catPath.Get)
		}
	})
}
