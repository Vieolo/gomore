/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/vieolo/gomore/goyaml"
	"github.com/vieolo/termange"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:     "run <command>",
	Example: "gomore run build",
	Short:   "Runs a command from go.yaml",
	Long:    `Runs a pre-defined command from go.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		gy, gyErr := goyaml.ReadGoYAML()
		if gyErr != nil {
			termange.PrintErrorln(gyErr.Error())
			os.Exit(1)
		}

		listFlag, _ := cmd.Flags().GetBool("list")
		if listFlag {
			gy.PrintCommandList("Available commands in go.yaml")
			return
		}

		if len(args) == 0 || args[0] == "" {
			termange.PrintErrorln("No command was provided!")
			fmt.Printf("\nUsage:\n  gomore run <command>\n\n")
			gy.PrintCommandList("Here are the available commands")
			return
		}

		name := args[0]
		c, ok := gy.Commands[name]
		if !ok {
			termange.PrintErrorln("The selected command is not listed in go.yaml")
			gy.PrintCommandList("Here are the available commands")
			return
		}

		if c == "" {
			termange.PrintErrorln("The selected command is empty!")
			return
		}

		// Preparing for the SIGTERM
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

		// Starting the process of the command
		termange.PrintInfof("Running %s...\n", name)
		commandRunner := exec.CommandContext(ctx, "sh", "-c", c)
		commandRunner.Stdout = os.Stdout
		commandRunner.Stderr = os.Stderr
		commandRunner.Env = os.Environ()

		cerr := commandRunner.Start()
		if cerr != nil {
			termange.PrintErrorln(cerr.Error())
			return
		}

		go func() {
			<-sigChan
			// If after 10 seconds, the process is still
			// not completed, we force run the `cancel`
			go func() {
				time.Sleep(10 * time.Second)
				cancel()
			}()
		}()

		commandRunner.Wait()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().BoolP("list", "l", false, "List the available commands defined in go.yaml")
}
