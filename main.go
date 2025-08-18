package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("trial.txt")
	if err != nil {
		log.Fatal("error", err)
	}

	str := ""

	for {
		data := make([]byte, 8)

		n, err := file.Read(data)

		if err != nil {
			fmt.Printf("read: %s\n", str)
			break
		}

		if newLineIndex := bytes.IndexByte(data, '\n'); newLineIndex != -1 {
			str += string(data[:newLineIndex])
			fmt.Printf("read: %s\n", str)
			str = string(data[newLineIndex+1:])
		} else {
			str += string(data[:n])
		}

		// data = data[:n]

		// if i := bytes.IndexByte(data, '\n'); i != 0 {
		// 	str += string(data[:i])
		// 	data = data[i+1:]
		// 	fmt.Printf("read: %s\n", str)
		// 	str = ""
		// }

		// fmt.Printf("read: %s\n", string(data[:n]))
	}

}
