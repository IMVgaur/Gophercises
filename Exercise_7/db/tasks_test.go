package db

import (
	"testing"

	"github.ibm.com/CloudBroker/dash_utils/dashtest"
)

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
func TestAddTask(t *testing.T) {
	db, _ := Init("/home/gslab/tasks.db")
	err := AddTask("testing123")
	if err != nil {
		t.Errorf("Expected result No error, But got Error %v", err)
	}
	db.Close()
}

func TestInit(t *testing.T) {
	db, err := Init("/home/gslab/tasks.db")
	if err != nil {
		t.Errorf("Expected result No error, But got Error %v", err)
	}
	db.Close()
}

func TestInitNegative(t *testing.T) {
	_, err := Init("/home/gslab/task.db")
	if err == nil {
		t.Errorf("Expected result error, But got NO Error")
	}
}

func TestListTasks(t *testing.T) {
	db, _ := Init("/home/gslab/tasks.db")
	_, err := GetAllTasks()
	if err != nil {
		t.Errorf("Expected result No error, But got Error %v", err)
	}
	db.Close()
}

func TestDeleteTask(t *testing.T) {
	db, _ := Init("/home/gslab/tasks.db")
	err := DeleteTask(1)
	if err != nil {
		t.Errorf("Expected result No error, But got Error %v", err)
	}
	db.Close()
}
