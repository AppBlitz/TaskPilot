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
		Example: "./task-cli add [description task]",
		Aliases: []string{"a"},
		Run: func(cmd *cobra.Command, args []string) {
			log := log.Fatal
			if len(args) == 0 {
				log("Arguments no is valid")
			}
			if len(args) > 1 {
				log("Amount arguments no valid")
			}
			CommandCreateTasks(args[0])
		},
	}

	deleteTask = &cobra.Command{
		Use:     "delete",
		Short:   "Delete a task for id",
		Aliases: []string{"d"},
		Example: "./task-cli [id]",
		Run: func(cmd *cobra.Command, args []string) {
			number := VerificationNumber(args[0])
			log := log.Fatal
			if len(args) == 0 {
				log("Pleaso enter index")
			}
			if len(args) > 1 {
				log("count of parameter no valid")
			}
			if number < 0 {
				fmt.Printf("%s\n", " id of task no acepted, please verification")
			}
			message, erro := handlers.DeleteTask(number)
			if erro != nil {
				log(erro)
			}
			fmt.Printf("%s", message)
		},
	}
	listAllTask = &cobra.Command{
		Use:     "list",
		Short:   "Show all tasks save",
		Aliases: []string{"l"},
		Example: "./task-cli list",
		Run: func(cmd *cobra.Command, args []string) {
			CreateResponseListTasks(args)
		},
	}
	updateTask = &cobra.Command{
		Use: "update",
		Run: func(cmd *cobra.Command, args []string) {
			number := VerificationNumber(args[0])
			log := log.Fatal
			if len(args) < 2 {
				log("arguments no ")
			}
			if len(args) > 2 {
				log("")
			}
			if number > 0 {
				message, erro := handlers.UpdateTaks(number, args[1])
				if erro != nil {
					log(erro)
				} else {
					fmt.Printf("%s", message)
				}
			}
		},
	}
	markDone = &cobra.Command{
		Use:     "mark-done",
		Example: "/task-cli mark-done [id task]",
		Run: func(cmd *cobra.Command, args []string) {
			number := VerificationNumber(args[0])
			log := log.Fatal
			length := len(args)
			if length > 2 || length < 0 {
				log("count options no valid")
			}
			if number < 0 {
				log("ID no valid")
			}
			handlers.MarkDone(number)
		},
	}
	markProgress = &cobra.Command{
		Use:     "mark-in-progress",
		Example: "./task-cli mark-in-progress [id task]",
		Run: func(cmd *cobra.Command, args []string) {
			number := VerificationNumber(args[0])
			length := len(args)
			log := log.Fatal
			if length < 0 || length > 2 {
				log("amount arguments no valid")
			}
			if number < 0 {
				log("ID negative, have positive")
			}
			handlers.MarkProgress(number)
		},
	}
)

func init() {
	rootCmd.AddCommand(commandAdd)
	rootCmd.AddCommand(deleteTask)
	rootCmd.AddCommand(listAllTask)
	rootCmd.AddCommand(updateTask)
	rootCmd.AddCommand(markDone)
	rootCmd.AddCommand(markProgress)
}
