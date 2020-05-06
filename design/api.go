package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// API describes the global properties of the API server.
var _ = API("kreditbureau", func() {
	Title("Credit Reference Bureau Service")

	Description("HTTP service implementing the Kenyan CRB regulations")

	Version("1.0")

	Docs(func() {
		Description("Additional documentation")
		URL("https://github.com/wondenge/kreditbureau/wiki")
	})

	License(func() {
		Name("Mozilla Public License 2.0")
		URL("https://github.com/wondenge/kreditbureau/blob/master/LICENSE")
	})

	TermsOfService("https://help.github.com/articles/github-terms-of-API/")

	Contact(func() {
		Name("William Ondenge")
		Email("ondengew@gmail.com")
		URL("https://github.com/wondenge")
	})

	Server("kreditbureau", func() {
		Host("localhost", func() {
			Description("default host")
			URI("http://localhost:8000/kreditbureau")
			URI("grpc://localhost:8080/kreditbureau")
		})
	})

	HTTP(func() {
		// Prefix to HTTP path of all requests.
		Path("/api")
	})
})




