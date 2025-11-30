package main

import (
	_ "embed"

	"github.com/vieolo/gomore/cmd"
)

//go:embed go.yaml
var thisGyByte []byte

func main() {
	cmd.ThisGyByte = thisGyByte
	cmd.Execute()
}
