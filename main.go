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
	d:pasta1
	d:*pasta11
	f:**pasta12.txt
	d:**pp
	d:**ii
	d:***ty
	f:****u.txt
	f:****y.txt
	d:pasta13.txt
	d:pasta2
	f:*pasta2.txt
	`)
	fmt.Printf("Thanks for using organLib")
}
