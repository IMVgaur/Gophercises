package cmd

import (
	"fmt"
	"strings"

	"github.com/Gophercises/Exercise_7/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add your task to the List.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.AddTask(task)
		if err == nil {
			fmt.Printf("Task \"%s\" has been added to the list.\n", task)
		}
		fmt.Printf("Something went wrong : %v", err)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
