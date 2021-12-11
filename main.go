package main

import (
	"fmt"
	"organLib/generators"
	"organLib/paths"
)

func main() {
	r := paths.NewWithPath(".")
	// generators.SearchFileInOutputTree(r, "o.txt")
	//files.DeleteFileInRoot(&file, r)
	generators.SearchFileInOutputTree(r, "README.md")
	fmt.Printf("Thanks for using organLib")
}
