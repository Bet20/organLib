package main

import (
	"fmt"
	"organLib/files"
	"organLib/generators"
	"organLib/paths"
)

func main() {
	r := paths.NewWithPath("D:\\docs\\pasta")
	// generators.SearchFileInOutputTree(r, "o.txt")
	file, err := files.SearchFile(r, "o.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	file.CreateLogFile("testLocal.txt")
	//files.DeleteFileInRoot(&file, r)
	generators.CreateDirectoryWithTemplate(*r, `
	/pasta1
	/pasta1/pasta11.txt
	/pasta1/pasta12.txt
	/pasta1/pasta13.txt
	/pasta2
	/pasta2/pasta2.txt
	`)
	fmt.Printf("Thanks for using organLib")
}
