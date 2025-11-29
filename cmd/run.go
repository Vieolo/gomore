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
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:     "run <command>",
	Example: "gomore run build",
	Short:   "Runs a command from go.yaml",
	Long:    `Runs a pre-defined command from go.yaml.`,
	Run: func(cmd *cobra.Command, args []string) {
		gy, gyErr := goyaml.ReadGoYAML()
		if gyErr != nil {
			termange.PrintErrorln(gyErr.Error())
			os.Exit(1)
		}

		if len(args) == 0 {
			termange.PrintErrorln("No command was provided!")
			fmt.Printf("\nUsage:\n  gomore run <command>\n\n")
			gy.PrintCommandList("Here are the available commands")
			return
		}

		name := args[0]
		if name == "" {
			os.Exit(1)
		}

		c, ok := gy.Commands[name]
		if !ok {
			termange.PrintErrorln("The selected command is not listed in go.yaml")
			gy.PrintCommandList("Here are the available commands")
			return
		}

		termange.PrintInfof("Running %s...\n", name)
		stdo, stde, cerr := termange.RunRawCommand(c)
		if cerr != nil {
			termange.PrintErrorln(cerr.Error())
			return
		}
		stdeStr := stde.String()
		if stdeStr != "" {
			fmt.Println(stde.String())
		}
		stdoStr := stdo.String()
		if stdoStr != "" {
			fmt.Println(stdo.String())
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
