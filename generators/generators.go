package generators

import (
	"fmt"
	"io/fs"
	"organLib/paths"
	"os"
	"path/filepath"
	"strings"
)

var original = "\033[0m"
var blue = "\033[34m"
var yellow = "\033[33m"
var red = "\033[31m"

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

func CreateDirectoryWithTemplate(root paths.RootDir, template string) {
	var linesWOSep []string
	lines := strings.Split(template, "\n")
	for i, l := range lines {
		linesWOSep = append(linesWOSep, strings.TrimSpace(l))
		fmt.Println(linesWOSep[i])
	}
	for _, line := range linesWOSep {
		if !strings.Contains(line, ".") {
			os.Mkdir(root.Root+line, 0755)
		} else {
			os.Create(root.Root + line)
		}
	}
}
