package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	for _, file := range listFiles("testdata") {
		fmt.Println(file)
	}
}

func listFiles(dirname string) []string {
	var dirs []string

	files, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.HasPrefix(f.Name(), ".") {
			continue
		}

		dirs = append(dirs, f.Name())
	}

	return dirs
}
