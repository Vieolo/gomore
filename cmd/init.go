package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/vieolo/filange"
	"github.com/vieolo/godotyaml"
	"github.com/vieolo/termange"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generates the initial go.yaml file for the project",
	Long:  `Generates the initial go.yaml file for the project`,
	Run: func(cmd *cobra.Command, args []string) {
		// Checking the `go.yaml` file alreay exists
		if filange.FileExists("go.yaml") {
			termange.PrintInfoln("There is already a go.yaml at the current directory!")
			return
		}

		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			name = "myproject"
		}

		gy := godotyaml.New(godotyaml.Metadata{Name: name}, false)
		gy.SetExternalConfig("gomore", map[any]any{"commands": map[string]string{"build": "go build main.go", "test": "go test"}})
		wErr := gy.SaveNew("go.yaml")

		if wErr != nil {
			termange.PrintErrorln(wErr.Error())
			os.Exit(1)
		}

		termange.PrintSuccessln("go.yaml is generated for your project!")

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	initCmd.Flags().String("name", "myproject", "Name of the project")
}
