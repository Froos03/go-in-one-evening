package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fileNames := listFiles("testdata")
	fmt.Println(strings.Join(fileNames, " "))
}

func listFiles(dirname string) []string {
	var dirs []string

	files, _ := os.ReadDir(dirname)

	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}
