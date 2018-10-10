package cmd

import (
	"fmt"
	"strings"

	"github.com/Gophercises/Exercise_7/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		err := db.AddTask(task)
		msg := "Error occurred in add cmd"
		if err == nil {
			msg = fmt.Sprintf("Added \"%s\" to your task list.\n", task)
		}
		fmt.Printf(msg)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
