package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func createScaffoldDirs(s Scaffold) error {
	_, err := os.Stat(s.Location)
	if err == nil {
		return errors.New(fmt.Sprintf("Local path: %s already exists", s.Location))
	} else if !os.IsNotExist(err) {
		return err
	}

	err = os.MkdirAll(filepath.Join(s.Location, "handlers"), 0755)
	if err != nil {
		return err
	}

	if s.Role {
		err := os.MkdirAll(filepath.Join(s.Location, "static", "js"), 0755)
		if err != nil {
			return err
		}

		err = os.MkdirAll(filepath.Join(s.Location, "static", "css"), 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
