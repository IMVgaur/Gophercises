package primitive

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
)

type Mode int

const (
	ModeCombo Mode = iota
	ModeTriangle
	ModeRect
	ModeEllipse
	ModeCircle
	ModeRotatedRect
	ModeBeziers
	ModeRotatedEllipse
	ModePolygon
)

//Transform ...
func Transform(r io.Reader, ext string, numShapes int, opts []string) (io.Reader, error) {
	inputFile, err := tempFile("in_", ext)
	if err != nil {
		return nil, errors.New("New file could not be created")
	}
	defer os.Remove(inputFile.Name())
	outputFile, _ := tempFile("in_", ext)
	defer os.Remove(outputFile.Name())
	io.Copy(inputFile, r)
	_, err = primitive(inputFile.Name(), outputFile.Name(), numShapes, opts...)
	if err != nil {
		return nil, errors.New("Image transformation failed")
	}
	return dataBuffer(outputFile)
}

func dataBuffer(outputFile *os.File) (*bytes.Buffer, error) {
	buf := bytes.NewBuffer(nil)
	_, err := io.Copy(buf, outputFile)
	if err != nil {
		return nil, errors.New("Error occured while reading data")
	}
	return buf, err
}

func primitive(inFile, outFile string, numShape int, args ...string) (string, error) {
	arg := (fmt.Sprintf("-i %s -o %s -n %d", inFile, outFile, numShape))
	cmd := exec.Command("primitive", (append(strings.Fields(arg), args...))...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), err
}

//tempFile function creates a new file
func tempFile(name, ext string) (*os.File, error) {
	home, _ := homedir.Dir()
	imgPath := filepath.Join(home, "img")
	file, err := ioutil.TempFile(imgPath+"/", name)
	if err != nil {
		return nil, errors.New("failed to create new temp file")
	}
	defer os.Remove(file.Name())
	f, err := os.Create(fmt.Sprintf("%s.%s", file.Name(), ext))
	if err != nil {
		return nil, err
	}
	return f, nil
}

/*
func WithMode works as selector of one the available mode
based on the input : mode int
*/
func WithMode(mode Mode) []string {
	return []string{"-m", fmt.Sprintf("%d", mode)}
}
