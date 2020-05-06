package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// The directory gen/openapi and gen/http/openapi contain the
// implementation for a HTTP handler that serves a openapi.json file.
var _ = Service("openapi", func() {
	// Serve the file with relative path ../../gen/http/openapi.json
	// for requests sent to /swagger.json.
	Files("/swagger.json", "../../gen/http/openapi.json")
})

var _ = Service("swagger", func() {
	Description("The swagger service serves the API swagger definition.")
	HTTP(func() {
		Path("/swagger")
	})
	Files("/swagger.json", "gen/http/openapi.json", func() {
		Description("JSON document containing the API swagger definition")
	})
})
