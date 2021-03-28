package cmd

import (
	"fmt"

	"github.com/BA1RY/prack/handlers"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates prack.yaml file in project directory",
	Long:  `Creates prack.yaml file in project directory. This file can be used to add project specific commands which are to be executed.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := handlers.HandleInit()
		if err != nil {
			fmt.Println(err.Error())
		}
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
}
