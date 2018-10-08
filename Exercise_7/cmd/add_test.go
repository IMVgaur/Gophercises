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

func TestAddCmd(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	dbc, _ := db.Init(dbPath)
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0600)
	OldStdOut := os.Stdout
	os.Stdout = file
	a := []string{"complete exercise"}
	addCmd.Run(addCmd, a)
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("Error occured while testing ", err)
	}
	output := string(content)
	val := strings.Contains(output, "added")
	assert.Equalf(t, true, val, "They should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = OldStdOut
	file.Close()
	dbc.Close()
}
