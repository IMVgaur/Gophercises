package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Gophercises/Exercise_7/db"
	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

func TestListCmd(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	dbc, _ := db.Init(dbPath)
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	OldStdout := os.Stdout
	os.Stdout = file
	a := []string{""}
	listCmd.Run(listCmd, a)
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("Error occured while testing..")
	}
	output := string(content)
	val := strings.Contains(output, "You have following tasks")
	assert.Equalf(t, true, val, "They should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = OldStdout
	file.Close()
	dbc.Close()
}
