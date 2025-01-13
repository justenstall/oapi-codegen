package codegen

// // TODO: uncomment
// var (
// 	// globalState stores all global state. Please don't put global state anywhere
// 	// else so that we can easily track it.
// 	globalState State
// 	// globalStateError stores any errors from producing the global state.
// 	// TODO: where to surface this?
// 	globalStateError error
// )

// // TODO: uncomment
// // Generate uses the Go templating engine to generate all of our server wrappers from
// // the descriptions we've built up above from the schema objects.
// // opts defines
// func Generate(spec *openapi3.T, opts Configuration) (string, error) {
// 	// Create new state as a local variable
// 	state, err := NewGenerator(spec, opts)
// 	if err != nil {
// 		return "", err
// 	}
// 	// Set the global state to the locally-scoped generator AFTER generating code
// 	// This preserves the existing global state behavior without race conditions
// 	defer func() { globalState = *state }()
// 	// Run the locally-scoped generator
// 	return state.Generate()
// }

// // TODO: uncomment
// func SetGlobalStateSpec(spec *openapi3.T) {
// 	globalStateError = errors.Join(globalStateError,
// 		globalState.SetSpec(spec))
// }

// // TODO: uncomment
// func SetGlobalStateOptions(opts Configuration) {
// 	globalStateError = errors.Join(globalStateError,
// 		globalState.SetOptions(opts))
// }
