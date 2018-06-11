package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//os.Args -- slice of type string
	//fmt.Println(os.Args)

	//to get filepath we want to open -- in this case 1st argument given after "go run main.go MyFile.txt"
	//MyFile.txt is 1st Arg
	//os.Args[1]

	//to test if opening the file passed on
	//fmt.Println(os.Args[1])

	//pointer to a file, this is a reference to a struct (rather than the address)
	f, err := os.Open(os.Args[1])

	//add error handling
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
