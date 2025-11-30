package goyaml

import (
	"errors"
	"fmt"
	"os"

	"github.com/vieolo/termange"
	"gopkg.in/yaml.v2"
)

type GoYAML struct {
	// The name of the project. The name should be the human-readable name and not
	// the git path. e.g., gomore
	Name string `yaml:"name"`
	// [Optional] description of the project
	Description string `yaml:"description"`
	// version of the project
	Version string `yaml:"version"`
	// Pre-defined commands in a key-value pair format
	Commands map[string]string `yaml:"commands"`
}

// Prints the list of the pre-defined commands, if any
func (y GoYAML) PrintCommandList(title string) {
	if len(y.Commands) == 0 {
		return
	}

	termange.PrintInfoln(title)
	for k := range y.Commands {
		fmt.Printf(" - %s\n", k)
	}
}

// Reads and parses the go.yaml file
func ReadGoYAML() (GoYAML, error) {
	b, readErr := os.ReadFile("go.yaml")
	if readErr != nil {
		return GoYAML{}, errors.New("go.yaml file could not found")
	}

	var raw GoYAML
	parseErr := yaml.Unmarshal(b, &raw)
	if parseErr != nil {
		return GoYAML{}, parseErr
	}

	return raw, nil
}
