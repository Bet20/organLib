package paths

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
)

type RootDir struct {
	Root string
	Size int
	File []fs.FileInfo
}

// Criar um novo root
func New() *RootDir {
	root := RootDir{
		Root: root(),
		Size: rootSize("."),
	}
	var err error
	root.File, err = ioutil.ReadDir(".")
	if err != nil {
		fmt.Printf("Error while getting the root.")
	}
	return &root
}

func NewWithPath(path string) *RootDir {

	fixedPath := path

	if strings.Contains(path, "/") {
		fixedPath = strings.ReplaceAll(path, "/", "\\")
	}

	root := RootDir{
		Root: fixedPath,
		Size: rootSize(path),
	}

	var err error
	root.File, err = ioutil.ReadDir(path)

	if err != nil {
		fmt.Printf("Error while getting the root.")
	}

	return &root
}

func root() string {
	d, err := os.Getwd()
	if err != nil {
		println("Error while reading the root directory.")
	}
	return d + "\\"
}

func rootSize(path string) int {
	d, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	var size int64 = 0

	for _, x := range d {
		size += x.Size()
	}

	return int(size)
}
