package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root, _ := filepath.Abs(".")
	fmt.Println(root)
	err := filepath.Walk(root, pr)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func pr(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if path != "." {
		if info.IsDir() {
			fmt.Println("d:", path)
		} else {
			fmt.Println("f:", path)
		}
	}
	return nil

}
