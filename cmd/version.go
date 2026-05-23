package cmd

import (
	"bytes"

	"github.com/spf13/cobra"
	"github.com/vieolo/godotyaml"
	"github.com/vieolo/termange"
)

// The bytes is injected from main.go downward
var ThisGyByte []byte

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the version of gomore cli",
	Long:  "Displays the version of gomore cli",
	Run: func(cmd *cobra.Command, args []string) {
		doc, _ := godotyaml.Parse(bytes.NewReader(ThisGyByte))
		termange.PrintInfof("v%s\n", doc.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
