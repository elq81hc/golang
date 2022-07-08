// read file path from arg
// read file
// print out file content to the terminal

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// print out usage
	if len(os.Args) < 2 {
		fmt.Println("usage:", filepath.Base(os.Args[0]), "/path/to/file")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, file)
}
