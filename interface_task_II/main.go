package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fn := os.Args[1]
	fmt.Println(fn)

	file, err := os.Open(fn)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println(file)
	io.Copy(os.Stdout, file)
}
