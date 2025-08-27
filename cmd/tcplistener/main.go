package main

import (
	"fmt"
	"log"
	"net"

	"github.com/akhand08/http-server-golang/internal/request"
)

// func getLinesChannel(conn io.ReadCloser) <-chan string {
// 	out := make(chan string, 1)

// 	go func() {

// 		defer conn.Close()
// 		defer close(out)

// 		str := ""

// 		for {
// 			data := make([]byte, 8)

// 			n, err := conn.Read(data)

// 			if err != nil {

// 				if err != io.EOF {
// 					log.Fatal("error: ", err)
// 				}

// 				break
// 			}

// 			data = data[:n]
// 			if newLineIndex := bytes.IndexByte(data, '\n'); newLineIndex != -1 {
// 				str += string(data[:newLineIndex])
// 				// fmt.Printf("read: %s\n", str)
// 				data = data[newLineIndex+1:]
// 				out <- str
// 				str = ""

// 			}

// 			str += string(data)

// 		}

// 		if len(str) != 0 {
// 			out <- str
// 		}

// 	}()

// 	return out
// }

func main() {

	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("error", err)
	}

	for {
		connection, err := listener.Accept()

		if err != nil {
			log.Fatal("error in accepting the connection: ", err)
		}

		httpRequest, err := request.RequestFromReader(connection)

		if err != nil {
			log.Fatal("Error in receiving the http request: ", err)
		}

		// Printing the Request Line
		fmt.Println("Request Line - ")
		fmt.Println("Method: ", httpRequest.RequestLine.Method)
		fmt.Println("Target: ", httpRequest.RequestLine.RequestTarget)
		fmt.Println("Method: ", httpRequest.RequestLine.HttpVersion)
		fmt.Println("Headers: ")
		for key, value := range httpRequest.RequestHeader {
			fmt.Println(key, "--> ", value)
		}
	}

	// lines := getLinesChannel(file)

	// for line := range lines {

	// 	fmt.Printf("read: %s\n", line)

	// }

}
