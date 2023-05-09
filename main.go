package main

import (
	"fmt"
	"gocomplete/cmd"
	"log"
	"os"
)

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
	fmt.Println(files)
	cmd.RootCmd.Execute()
}
