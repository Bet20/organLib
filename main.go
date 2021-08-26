package main

import (
	"fmt"
	"organLib/generators"
	"organLib/paths"
)

func main() {
	r := paths.NewWithPath("C:/Users/diogo/fm_factor_front/")
	path, err := generators.SearchFile(r, "LICENSE")
	if err != nil {
		panic(err)
	}
	path.CreateLogFile("License_Paths.txt")
	fmt.Printf("Thanks for using organLib")
}
