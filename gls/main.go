package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func gls(path string, option string) {
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
		if option != "all" {
			if strings.HasPrefix(file.Name(), ".") {
				continue
			}
		}
		fmt.Println("\t", file.Name())
	}
}

func check_options(option string) string {
	switch option {
	case "-a", "--all":
		return "all"
	default:
		return "nil"
	}
}

func main() {
	path := os.Args[1]
	option := check_options(os.Args[2])

	fmt.Println("Input path  :", path)
	fmt.Println("Input option:", option)
	fmt.Println("outputs:")
	gls(path, option)
}
