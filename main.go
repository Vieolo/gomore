package main

import (
	_ "embed"

	"github.com/vieolo/gomore/cmd"
)

// The project's go.yaml is embedded into the cli
// and then injected downward to the cmd module

//go:embed go.yaml
var thisGyByte []byte

func main() {
	cmd.ThisGyByte = thisGyByte
	cmd.Execute()
}
