package cmd

import (
	"fmt"
	"os"

	"github.com/AppBlitz/task_tracker/internal/handlers"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "add",
	Short:   "Command to add a new task",
	Long:    `Command to add a new task; does not create if the task already exists.`,
	Aliases: []string{"ad"},
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		if len(args) > 2 {
			fmt.Printf("%s\n", "No more than two arguments are accepted per command.")
		} else {
			fmt.Println(handlers.NewTasks(12, args[1]))
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
