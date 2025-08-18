package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func getLinesChannel(file io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {

		defer file.Close()
		defer close(out)

		str := ""

		for {
			data := make([]byte, 8)

			n, err := file.Read(data)

			if err != nil {

				if len(str) > 0 {
					out <- str
					out <- "end"
				}

				if err != io.EOF {
					log.Fatal("error: ", err)
				}

				break
			}

			if newLineIndex := bytes.IndexByte(data, '\n'); newLineIndex != -1 {
				str += string(data[:newLineIndex])
				// fmt.Printf("read: %s\n", str)
				out <- str
				str = string(data[newLineIndex+1:])
			} else {
				str += string(data[:n])
			}

		}

	}()

	return out
}

func main() {

	file, err := os.Open("trial.txt")
	if err != nil {
		log.Fatal("error", err)
	}

	lines := getLinesChannel(file)

	for line := range lines {

		fmt.Printf("read: %s\n", line)

	}

}
