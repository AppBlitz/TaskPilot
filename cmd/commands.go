// Package cmd
package cmd

import (
	"fmt"
	"strconv"

	"github.com/AppBlitz/task_tracker/internal/handlers"
	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:     "add",
		Short:   "Create new task",
		Example: "./Task-Tracker add [nameTaks]",
		Aliases: []string{"a"},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%v\n", args[0])
		},
	}

	deleteTask = &cobra.Command{
		Use:     "delete",
		Short:   "Delete a task for id",
		Aliases: []string{"d"},
		Example: "./Task-tracker [id]",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Printf("%s\n", "Pleaso enter index")
			}
			if len(args) > 1 {
				fmt.Printf("%s\n", "count of parameter no valid")
			}
			if VerificationNumber(args[0]) {
				fmt.Printf("%s\n", "Index no is number, please verification")
			}
			value, _ := strconv.Atoi(args[0])
			handlers.DeleteTask(value)
		},
	}
	listAllTask = &cobra.Command{
		Use:     "list",
		Aliases: []string{"l"},
		Run: func(cmd *cobra.Command, args []string) {
			CreateResponseListTasks()
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(deleteTask)
	rootCmd.AddCommand(listAllTask)
}
