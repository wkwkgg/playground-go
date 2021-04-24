package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func gls(path string) {
	// check file path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println(path, "does not exist.")
		return
	}

	// list directory contents
	fileList, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range fileList {
		fmt.Println("\t", file.Name())
	}
}

func main() {
	path := os.Args[1]

	fmt.Println("Input path: ", path)
	fmt.Println("outputs:")
	gls(path)
}
