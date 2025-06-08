package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fileData, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	if err = Lint(bytes.NewReader(fileData)); err != nil {
		log.Fatal(err)
	}
}
