package main

import (
	"os"
	"path/filepath"
	"text/template"
)

func writeToFile(s Scaffold, fileMap map[string][]byte) error {
	for fileName, data := range fileMap {
		f, err := os.Create(filepath.Join(s.Location, fileName))
		if err != nil {
			return err
		}
		f.Write(data)
		f.Close()
	}
	return nil
}

func renderToFile(s Scaffold, tmplMap map[string]string) error {
	tmpl := template.New("scaffold")

	for fileName, t := range tmplMap {
		f, err := os.Create(filepath.Join(s.Location, fileName))
		if err != nil {
			return err
		}

		tmpl, err = tmpl.Parse(t)
		if err != nil {
			return err
		}
		err = tmpl.Execute(f, s)
		if err != nil {
			return err
		}
		f.Close()
	}
	return nil
}
