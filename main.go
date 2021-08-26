package main

import (
	"fmt"
	"organLib/files"
	"organLib/generators"
	"organLib/paths"
)

func main() {
	r := paths.New()
	generators.SearchFileInOutputTree(r, "LICENSE")
	file, err := files.SearchFile(r, "LICENSE")
	if err != nil {
		fmt.Println(err.Error())
	}
	file.CreateLogFile("testLocal.txt")
	fmt.Printf("Thanks for using organLib")
}
