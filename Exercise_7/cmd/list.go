package cmd

import (
	"fmt"

	"github.com/Gophercises/Exercise_7/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show List of Tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, _ := db.GetAllTasks()
		if len(tasks) != 0 {
			fmt.Println("You have following tasks ...")
			for _, task := range tasks {
				fmt.Printf("%10d. %s\n", task.Key, task.Value)
			}
		}
		fmt.Println("No Task available")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
