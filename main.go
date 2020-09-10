package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	for i, _ := range os.Args[1:] {

		path := os.Args[i+1]

		readPrint(path)

	}

}

func readPrint(path string) {

	print(read(path))
}

func read(path string) (s string) {

	buf := bytes.NewBuffer(nil)
	data, err := os.Open(path)

	if err != nil {
		log.Println("File reading error", err)
		os.Exit(-1)
		return
	} else {

		io.Copy(buf, data)
		data.Close()

		s = string(buf.Bytes())
		
		return
	}
}

func print(s string) {
	fmt.Println(s)
}
