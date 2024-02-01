package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Scaffold struct {
	Name          string
	Location      string
	GitRepository string
	Role          bool
}

func (s *Scaffold) setupFlags(args []string) error {
	fs := flag.NewFlagSet("cli", flag.ExitOnError)
	fs.StringVar(&s.Name, "n", "", "Project name")
	fs.StringVar(&s.Location, "d", "", "Project location on disk")
	fs.StringVar(&s.GitRepository, "r", "", "Project remote repository URL")
	fs.BoolVar(&s.Role, "s", false, "Project will have static assets or not")

	err := fs.Parse(args)
	if err != nil {
		return fmt.Errorf("error parsing flags: %v", err)
	}

	return nil
}

func (s *Scaffold) validate() error {
	var validationErrors []string

	if s.Name == "" {
		validationErrors = append(validationErrors, "project name cannot be empty")
	}

	if s.Location == "" {
		validationErrors = append(validationErrors, "project path cannot be empty")
	}

	if s.GitRepository == "" {
		validationErrors = append(validationErrors, "project repository URL cannot be empty")
	}

	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, "\n"))
	}

	return nil
}

func generateScaffold(w io.Writer, s Scaffold) error {
	fmt.Fprintf(w, "Generating scaffold for project %s in %s\n", s.Name, s.Location)
	err := createScaffoldDirs(s)
	if err != nil {
		return err
	}
	err = createScaffoldFiles(s)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	s := Scaffold{}

	err := s.setupFlags(os.Args[1:])
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	err = s.validate()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	err = generateScaffold(os.Stdout, s)
	if err != nil {
		fmt.Printf("Error generating scaffold: %v\n", err)
	}
}
