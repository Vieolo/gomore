package goyaml

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type GoYAML struct {
	Name        string
	Description string
	Version     string
	Commands    map[string]string
	Config      map[string]any
}

func ReadGoYAML() (GoYAML, error) {
	b, readErr := os.ReadFile("go.yaml")
	if readErr != nil {
		return GoYAML{}, errors.New("go.yaml file could not found")
	}

	var raw map[string]any
	parseErr := yaml.Unmarshal(b, &raw)
	if parseErr != nil {
		return GoYAML{}, errors.New("go.yaml could not be parsed")
	}

	name, nameErr := getString(raw, "name", false)
	if nameErr != nil {
		return GoYAML{}, nameErr
	}

	description, descriptionErr := getString(raw, "description", true)
	if descriptionErr != nil {
		return GoYAML{}, descriptionErr
	}

	version, versionErr := getString(raw, "version", false)
	if versionErr != nil {
		return GoYAML{}, versionErr
	}

	goyaml := GoYAML{
		Name:        name,
		Description: description,
		Version:     version,
	}
	for k, v := range raw {
		fmt.Println(k, v)
	}

	return goyaml, nil
}

func getString(raw map[string]any, key string, optional bool) (string, error) {
	val, exist := raw[key]
	if !exist {
		if optional {
			return "", nil
		}
		return "", fmt.Errorf("%s is missing", key)
	}

	s, isString := val.(string)
	if !isString {
		return "", fmt.Errorf("%s is not a string", key)
	}

	return s, nil
}
