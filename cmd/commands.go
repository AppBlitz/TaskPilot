// Package cmd
package cmd

import (
	"fmt"
	"log"

	"github.com/AppBlitz/task_tracker/internal/handlers"
	"github.com/spf13/cobra"
)

var (
	commandAdd = &cobra.Command{
		Use:     "add",
		Short:   "Create new task",
		Example: "./Task-Tracker add [nameTaks]",
		Aliases: []string{"a"},
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				panic("Arguments no is valid")
			}
			if len(args) > 1 {
				panic("")
			}
			fmt.Printf("%v\n", handlers.CreateTask(args[0]))
		},
	}

	deleteTask = &cobra.Command{
		Use:     "delete",
		Short:   "Delete a task for id",
		Aliases: []string{"d"},
		Example: "./Task-tracker [id]",
		Run: func(cmd *cobra.Command, args []string) {
			number := VerificationNumber(args[0])
			if len(args) == 0 {
				panic("Pleaso enter index")
			}
			if len(args) > 1 {
				panic("count of parameter no valid")
			}
			if number < 0 {
				fmt.Printf("%s\n", " id of task no acepted, please verification")
			}
			handlers.DeleteTask(number)
		},
	}
	listAllTask = &cobra.Command{
		Use:     "list",
		Short:   "Show all tasks save",
		Aliases: []string{"l"},
		Run: func(cmd *cobra.Command, args []string) {
			CreateResponseListTasks()
		},
	}
	updateTask = &cobra.Command{
		Use: "update",
		Run: func(cmd *cobra.Command, args []string) {
			number := VerificationNumber(args[0])
			if len(args) < 2 {
				log.Fatal("arguments no ")
			}
			if len(args) > 2 {
				log.Fatal("")
			}
			if number > 0 {
				handlers.UpdateTaks(number, args[1])
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(commandAdd)
	rootCmd.AddCommand(deleteTask)
	rootCmd.AddCommand(listAllTask)
	rootCmd.AddCommand(updateTask)
}
