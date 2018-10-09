package cmd

import (
	"fmt"

	"github.com/IMVgaur/Gophercises/Exercise_7/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.GetAllTasks()
		if err != nil {
			fmt.Printf("error occured in list cmd")
		}
		if len(tasks) == 0 {
			fmt.Println("you have no Task pending to do...")
			return
		}
		fmt.Println("You have the following tasks:")
		for i, t := range tasks {
			fmt.Printf("%d. %s\n", i+1, t.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
