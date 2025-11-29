/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vieolo/gomore/goyaml"
	"github.com/vieolo/termange"
	"github.com/vieolo/termange/tui"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		gt, gtErr := goyaml.ReadGoYAML()
		if gtErr != nil {
			termange.PrintErrorln(gtErr.Error())
			os.Exit(1)
		}
		fmt.Println(gt)
		var name string
		if len(args) == 0 {
			name = tui.TextInput(tui.TextInputOptions{
				Prompt: "Enter the command name you wish to run",
			})
		} else {
			name = args[0]
		}

		if name == "" {
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
