package db

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
)

func TestInIt(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	dbc, err := Init(dbPath)
	if err != nil {
		fmt.Println("Error while testing...")
	}
	dbc.Close()

}

func TestAddTask(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	dbc, _ := Init(dbPath)
	_, err := AddTask("Testing add task")
	if err != nil {
		t.Error("Error occured while testing..")
	}
	dbc.Close()
}

func TestGetAllTasks(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	dbc, err := Init(dbPath)
	if err != nil {
		t.Error("Error occured while testing...")
	}
	_, err = GetAllTasks()
	if err != nil {
		t.Error("Error occured while testing")
	}
	dbc.Close()
}

func TestDeleteTask(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	dbc, _ := Init(dbPath)
	err := DeleteTask(79)
	if err != nil {
		t.Error("Error occured : ", err)
	}
	dbc.Close()
}
