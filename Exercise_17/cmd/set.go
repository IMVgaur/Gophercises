package cmd

import (
	"fmt"

	vault "github.com/IMVgaur/Gophercises/Exercise_17/vault"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the secret key in your storage space",
	Run: func(cmd *cobra.Command, args []string) {
		v := vault.File(encodingKey, secretsPath())
		key, value := args[0], args[1]
		err := v.Set(key, value)
		if err != nil {
			fmt.Printf("Error occured in Set : %v\n", err)
		}
		fmt.Println("Key and values has been Added successfully...!")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
