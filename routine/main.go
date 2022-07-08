// Requirements
// find all files in <input/path> exclude subdir
// input should be: toolName path/to/file
// exit if path not exist
// calc hash of each find in input path
// output the calculation to file with pattern:
// 	filename | file hash
// Ex: helloword.txt | hash number

package main

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	// ---- validation
	if len(os.Args) < 3 {
		fmt.Printf("usage: %s /path/to/directory /path/to/output\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	// second argument
	err := errors.New("")
	entry, err := os.Stat(os.Args[1])
	if err != nil || !entry.IsDir() {
		fmt.Println("ERROR: path is not a directory or exist")
		os.Exit(1)
	}
	// third argument
	ofp, err := os.OpenFile(os.Args[2], os.O_APPEND|os.O_CREATE, 777)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// ---- validation ----
	// wait for multiple goroutines to finish
	var wg sync.WaitGroup
	files, err := ioutil.ReadDir(os.Args[1])
	for _, file := range files {
		wg.Add(1)
		go func(f fs.FileInfo) {
			defer wg.Done()
			ofp.Write([]byte(calcHash(filepath.Join(os.Args[1], f.Name()))))
			ofp.Write([]byte("\n"))
		}(file)
	}
	wg.Wait()
}

func calcHash(fileName string) string {

	fmt.Printf(`[%s] start process on "%s"%s`, time.Now().Format(time.RFC3339Nano), fileName, "\n")
	time.Sleep(5 * time.Second)
	fmt.Println("[", time.Now().Format(time.RFC3339Nano), "]", "done process on", fileName)
	return time.Now().Format(time.RFC3339Nano)
}
