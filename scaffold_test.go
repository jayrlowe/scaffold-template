package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

func TestGenerateScaffold(t *testing.T) {
	tempDir := t.TempDir()
	testDir := filepath.Join(tempDir, "test_webapp")
	expectedDir := "./projects/web"

	s := Scaffold{
		Name:          "web",
		Location:      testDir,
		GitRepository: "github.com/username/web",
		Role:          true,
	}

	expectedOutput := fmt.Sprintf("Generating scaffold for project web in %s", testDir)
	expectedFileContents := map[string]string{
		filepath.Join(testDir, "go.mod"):                      filepath.Join(expectedDir, "go.mod"),
		filepath.Join(testDir, "server.go"):                   filepath.Join(expectedDir, "server.go"),
		filepath.Join(testDir, "handlers", "api.go"):          filepath.Join(expectedDir, "handlers/api.go"),
		filepath.Join(testDir, "handlers", "index.go"):        filepath.Join(expectedDir, "handlers/index.go"),
		filepath.Join(testDir, "handlers", "setup.go"):        filepath.Join(expectedDir, "handlers/setup.go"),
		filepath.Join(testDir, "handlers", "healthcheck.go"):  filepath.Join(expectedDir, "handlers/healthcheck.go"),
		filepath.Join(testDir, "static", "css", "styles.css"): filepath.Join(expectedDir, "static/css/styles.css"),
		filepath.Join(testDir, "static", "js", "index.js"):    filepath.Join(expectedDir, "static/js/index.js"),
	}

	byteBuf := new(bytes.Buffer)
	err := generateScaffold(byteBuf, s)
	if err != nil {
		t.Fatalf("Error generating scaffold: %v", err)
	}

	if output := byteBuf.String(); !strings.Contains(output, expectedOutput) {
		t.Errorf("Expected output does not match. Expected substring:\n%s\nActual output:\n%s", expectedOutput, output)
	}

	for filePath, expectedPath := range expectedFileContents {
		verifyFileContents(t, filePath, expectedPath)
	}
}

func verifyFileContents(t *testing.T, filePath, expectedPath string) {
	actualContents, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("Error reading file %s: %v", filePath, err)
		return
	}

	expectedContents, err := os.ReadFile(expectedPath)
	if err != nil {
		t.Errorf("Error reading expected file %s: %v", expectedPath, err)
		return
	}

	if !bytes.Equal(actualContents, expectedContents) {
		t.Errorf("File contents mismatch for %s. Expected:\n%s\nActual:\n%s", filePath, expectedContents, actualContents)
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
