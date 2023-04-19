package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var showHiddenFiles = flag.Bool("a", false, "show hidden files")

func main() {
	flag.Parse()

	dir := flag.Arg(0)
	if dir == "" {
		dir = "."
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !*showHiddenFiles && file.Name()[0] == '.' {
			continue
		}
		fmt.Println(file.Name())
	}
}
