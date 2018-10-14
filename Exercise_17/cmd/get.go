package cmd

import (
	"fmt"

	vault "github.com/IMVgaur/Gophercises/Exercise_17/vault"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a secret from your secret Space",
	Run: func(cmd *cobra.Command, args []string) {
		v := vault.File(encodingKey, secretsPath())
		key := args[0]
		value, err := v.Get(key)
		if err != nil {
			fmt.Printf("Error occured in Func Get : %v\n", err)
		}
		fmt.Printf("Key : %s 	Value : %s \n", key, value)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
