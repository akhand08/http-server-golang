package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func getLinesChannel(conn io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {

		defer conn.Close()
		defer close(out)

		str := ""

		for {
			data := make([]byte, 8)

			n, err := conn.Read(data)

			if err != nil {

				if err != io.EOF {
					log.Fatal("error: ", err)
				}

				break
			}

			data = data[:n]
			if newLineIndex := bytes.IndexByte(data, '\n'); newLineIndex != -1 {
				str += string(data[:newLineIndex])
				// fmt.Printf("read: %s\n", str)
				data = data[newLineIndex+1:]
				out <- str
				str = ""

			}

			str += string(data)

		}

		if len(str) != 0 {
			out <- str
		}

	}()

	return out
}

func main() {

	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("error", err)
	}

	trialStr := "GET /cookie http/1.1"

	parts := strings.SplitN(trialStr, " ", 3)

	fmt.Println(parts)

	for {
		connection, err := listener.Accept()

		if err != nil {
			log.Fatal("error", err)
		}

		for line := range getLinesChannel(connection) {

			fmt.Printf("%s\n", line)
			// fmt.Printf("%T\n", line)
			// fmt.Println(string(line[0]), " --> ", line[1], " ---> ", line[2])

		}
	}

	// lines := getLinesChannel(file)

	// for line := range lines {

	// 	fmt.Printf("read: %s\n", line)

	// }

}
