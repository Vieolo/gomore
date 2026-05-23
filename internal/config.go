package internal

import (
	"fmt"

	"github.com/vieolo/termange"
)

type Config struct {
	// Pre-defined commands in a key-value pair format
	Commands map[string]string `yaml:"commands"`
}

// Prints the list of the pre-defined commands, if any
func (y Config) PrintCommandList(title string) {
	if len(y.Commands) == 0 {
		return
	}

	termange.PrintInfoln(title)
	for k := range y.Commands {
		fmt.Printf(" - %s\n", k)
	}
}
