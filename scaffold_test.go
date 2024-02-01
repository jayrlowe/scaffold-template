package main

import (
	"bytes"
	"testing"
)

func TestScaffoldCreate(t *testing.T) {
	s := Scaffold{
		Name:     "MyProject",
		Location: "/path/to/project",
	}

	var buf bytes.Buffer
	generateScaffold(&buf, s)

	expected := "Generating scaffold for project MyProject in /path/to/project\n"
	if ok := buf.String(); ok != expected {
		t.Errorf("Expected: %s, Got: %s", expected, ok)
	}
}

func TestScaffoldSetupFlags(t *testing.T) {
	s := Scaffold{}

	args := []string{"-n", "TestProject", "-d", "/test/path", "-r", "https://github.com/test/test.git", "-s"}

	err := s.setupFlags(args)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if s.Name != "TestProject" || s.Location != "/test/path" || s.GitRepository != "https://github.com/test/test.git" || !s.Role {
		t.Errorf("Unexpected Scaffold configuration after setupFlags: %+v", s)
	}
}

func TestScaffoldValidate(t *testing.T) {
	tests := []struct {
		name          string
		scaffold      Scaffold
		expectedError string
	}{
		{
			name: "Valid Scaffold",
			scaffold: Scaffold{
				Name:          "ValidProject",
				Location:      "/valid/path",
				GitRepository: "https://github.com/valid/valid.git",
				Role:          true,
			},
			expectedError: "",
		},
		{
			name: "Missing Project Name",
			scaffold: Scaffold{
				Location:      "/valid/path",
				GitRepository: "https://github.com/valid/valid.git",
				Role:          true,
			},
			expectedError: "project name cannot be empty",
		},
		{
			name: "Missing Project Path",
			scaffold: Scaffold{
				Name:          "InvalidProject",
				GitRepository: "https://github.com/invalid/invalid.git",
				Role:          true,
			},
			expectedError: "project path cannot be empty",
		},
		{
			name: "Missing Repository URL",
			scaffold: Scaffold{
				Name:     "InvalidProject",
				Location: "/invalid/path",
				Role:     true,
			},
			expectedError: "project repository URL cannot be empty",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.scaffold.validate()

			if test.expectedError == "" && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if test.expectedError != "" && err == nil {
				t.Error("Expected error, but got nil")
			}

			if err != nil && test.expectedError != err.Error() {
				t.Errorf("Expected error: %s, Got: %s", test.expectedError, err.Error())
			}
		})
	}
}
