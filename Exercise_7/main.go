package main

import (
	"log"
	"path/filepath"

	cmd "github.com/IMVgaur/Gophercises/Exercise_7/cmd"
	db "github.com/IMVgaur/Gophercises/Exercise_7/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	Initializer()
}

//Initializer ...
//Intialize bolt db
func Initializer() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	_, err := db.Init(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	cmd.RootCmd.Execute()
}
