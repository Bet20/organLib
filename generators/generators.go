package generators

import (
	"fmt"
	"io/fs"
	"organLib/paths"
	"path/filepath"
	"strings"
)

var original = "\033[0m"
var blue = "\033[34m"
var yellow = "\033[33m"
var red = "\033[31m"

type node struct {
	name   string
	parent string
	indent int
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

// still undicided on the syntax of the templates
func CreateDirectoryWithTemplate(root paths.RootDir, template string) error {
	var linesWOSep []string
	var nodes []node
	//changedDir := false
	lines := strings.Split(template, "\n")
	for i, l := range lines {
		linesWOSep = append(linesWOSep, strings.TrimSpace(l))
		fmt.Println(linesWOSep[i])
	}
	for i, line := range linesWOSep {
		x := strings.Split(line, ":")
		ii := 0
		p := ""
		fmt.Printf("* number -> %d", strings.Count(line, "*"))
		switch x[0] {
		case "d":
			ii++
			numofast := strings.Count(line, "*")
			if numofast > 0 {
				if numofast == nodes[ii-1].indent {
					p = nodes[ii-1].parent
				} else if numofast < nodes[ii-1].indent {

				} else {
					p = nodes[ii-1].name
				}
			} else {
				ii = 0
			}
			if ii > 0 {
				//println("numofast : " + strconv.Itoa(numofast) + " | ii-1 numofast : " + strconv.Itoa(nodes[ii-1].indent))
			}
			nodes = append(nodes, node{name: line, parent: p, indent: numofast})
		case "f":
			nodes = append(nodes, node{name: line, parent: linesWOSep[i-1], indent: strings.Count(line, "*")})
		}

		/*

			if strings.Contains(line, "d:") {

				changedDir = true
			} else if changedDir {
				nodes = append(nodes, node{name: strings.Trim(line, "*"), parent: linesWOSep[i-1]})
				changedDir = false
			} else {
				nodes = append(nodes, node{name: strings.Trim(line, "*"), parent: nodes[i-1].parent})
			}


				if !strings.Contains(line, ".") {
					os.Mkdir(root.Root+line, 0755)
				} else {
					os.Create(root.Root + line)
				}
		*/
	}

	for _, x := range nodes {
		fmt.Printf("Name : %s | Parent : %s | Ident : %d \n", x.name, x.parent, x.indent)
	}

	return nil
}
