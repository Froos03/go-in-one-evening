package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fileNames := listFiles("testdata")
	fmt.Println(strings.Join(fileNames, " "))
}

func listFiles(dirname string) []string {
	var dirs []string

	files, err := os.ReadDir(dirname)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}
