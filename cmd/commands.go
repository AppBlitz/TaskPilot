package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:     "add",
		Short:   "Print the version number of Hugo",
		Long:    `All software has versions. This is Hugo's`,
		Example: "./Task-Tracker add",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
		},
	}

	deleteTask = &cobra.Command{
		Use:     "delete",
		Aliases: []string{"d"},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s", "Delete task succesfully")
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(deleteTask)
}
