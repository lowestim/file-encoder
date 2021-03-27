package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("file name expected")
		return
	}

	var filename string = os.Args[1]
	file, err := os.Open(filename)
	HandleError(err)

	defer file.Close()

	var reader = bufio.NewReader(file)
	content, err := ioutil.ReadAll(reader)
	HandleError(err)
	encoded := base64.StdEncoding.EncodeToString(content)
	fmt.Println(encoded)
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
