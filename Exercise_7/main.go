package main

import (
	"log"
	"path/filepath"

	"github.com/Gophercises/Exercise_7/cmd"
	"github.com/Gophercises/Exercise_7/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	Initializer()
}

func Initializer() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	_, err := db.Init(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	cmd.RootCmd.Execute()
}
