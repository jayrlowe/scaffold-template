package main

import (
	_ "embed"
)

//go:embed scaffold-template/go.mod.tmpl
var goModTmpl string

//go:embed scaffold-template/server.go.tmpl
var serverGoTmpl string

//go:embed scaffold-template/handlers/api.go.tmpl
var apiGoTmpl string

//go:embed scaffold-template/handlers/healthcheck.go
var healthcheckGo []byte

//go:embed scaffold-template/handlers/index.go.tmpl
var indexGoTmpl string

//go:embed scaffold-template/handlers/setup.go.tmpl
var setupGoTmpl string

//go:embed scaffold-template/static/css/styles.css
var stylesCss []byte

//go:embed scaffold-template/static/js/index.js
var indexJs []byte

func createScaffoldFiles(s Scaffold) error {
	tmplMap := map[string]string{
		"go.mod":            goModTmpl,
		"server.go":         serverGoTmpl,
		"handlers/api.go":   apiGoTmpl,
		"handlers/setup.go": setupGoTmpl,
	}
	if s.Role {
		tmplMap["handlers/index.go"] = indexGoTmpl
	}
	err := renderToFile(s, tmplMap)
	if err != nil {
		return err
	}

	fileMap := map[string][]byte{
		"handlers/healthcheck.go": healthcheckGo,
	}

	if s.Role {
		fileMap["static/js/index.js"] = indexJs
		fileMap["static/css/styles.css"] = stylesCss
	}

	return writeToFile(s, fileMap)
}
