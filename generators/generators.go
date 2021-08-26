package generators

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"organLib/paths"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var original = "\033[0m"
var blue = "\033[34m"
var yellow = "\033[33m"
var red = "\033[31m"

type File struct {
	Name  string
	Paths []string
	Size  []int64
}

func OutputTreeToScreen(root *paths.RootDir) error {
	fmt.Println(blue + "\nDirectory Tree")
	addedPath := root.Root
	var i int = 0
	var x int = 0
	str, ok := printOutput(root.File, i)
	var tempStr string
	for {
		fmt.Println(str)
		if !ok {
			i++
			tempStr = str + "/"
			addedPath += tempStr
			fs, err := ioutil.ReadDir(addedPath)
			if err != nil {
				panic(err)
			}
			str, ok = printOutput(fs, i)
		} else {
			fmt.Println(i)
			i--
			addedPath = ""
			str = strings.Trim(addedPath, tempStr)
			ok = false
		}
		if x >= 40 {
			break
		}
		x++
	}
	return nil
}

func printOutput(dir []fs.FileInfo, i int) (string, bool) {
	for _, f := range dir {
		fmt.Printf(original+"\n"+tabRepeater(i, false)+
			"| %s", f.Name())
		if f.IsDir() {
			return f.Name(), false
		}
	}
	return "", true
}

func tabRepeater(count int, can bool) string {
	var ret string
	for i := 0; i < count; i++ {

		if can && i == count-1 {
			ret += "    ==> "
			break
		}
		ret += "\t"
	}

	return ret
}

func BadOutputTreeToScreen(root *paths.RootDir) error {
	fmt.Println(blue + "\nDirectory Tree")
	o := strings.Count(root.Root, "\\")
	can := false
	err := filepath.Walk(root.Root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		x := strings.Count(path, "\\")
		col := original
		if info.IsDir() {
			col = yellow
			can = false
		}
		fmt.Println(col + tabRepeater(x-o, can) + "| " + info.Name())
		can = false
		if info.IsDir() {
			can = true
		}
		return nil
	})

	return err
}

func SearchFile(root *paths.RootDir, fileName string) (File, error) {
	var files File
	files.Name = fileName
	err := filepath.Walk(root.Root, func(path string, info fs.FileInfo, err error) error {
		if info.Name() == fileName {
			files.Paths = append(files.Paths, path)
			files.Size = append(files.Size, info.Size())
		}
		return nil
	},
	)

	if err != nil {
		return files, err
	}

	if len(files.Paths) > 0 {
		return files, nil
	}

	return files, nil
}

func (file *File) CreateLogFile(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(f)
	defer f.Close()

	writer.WriteString("Paths in Root for " + file.Name + "\n\n")
	for i, str := range file.Paths {
		writer.WriteString("Size :> " + strconv.Itoa(int(file.Size[i])) + " | Path :> " + str + "\n")
	}

	writer.Flush()
	return nil
}

func SearchFileInOutputTree(root *paths.RootDir, fileName string) error {
	fmt.Println(blue + "\nDirectory Tree")
	o := strings.Count(root.Root, "\\")
	can := false
	isFile := false
	err := filepath.Walk(root.Root, func(path string, info fs.FileInfo, err error) error {
		if info.Name() == fileName {
			isFile = true
		}
		if err != nil {
			panic(err)
		}
		x := strings.Count(path, "\\")
		col := original
		if info.IsDir() {
			col = yellow
			can = false
		}
		if isFile {
			fmt.Println(red + tabRepeater(x-o, can) + "| " + "**" + info.Name() + "**")
			isFile = false
		} else {
			fmt.Println(col + tabRepeater(x-o, can) + "| " + info.Name())
		}

		can = false
		if info.IsDir() {
			can = true
		}
		return nil
	})

	return err
}
