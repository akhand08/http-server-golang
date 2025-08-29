package server

import (
	"log"
	"net"
	"strconv"
	"sync/atomic"

	"github.com/akhand08/http-server-golang/internal/response"
)

type Server struct {
	listener  net.Listener
	isRunning atomic.Bool
}

func Serve(port int) (*Server, error) {

	listenAt := ":" + strconv.Itoa(port)

	listener, err := net.Listen("tcp", listenAt)

	if err != nil {
		return nil, err
	}

	server := &Server{listener: listener}
	server.isRunning.Store(true)
	go server.listen()

	return server, nil

}

func (s *Server) Close() error {
	s.isRunning.Store(false)
	return s.listener.Close()

}

func (s *Server) listen() {

	for s.isRunning.Load() {
		connection, err := s.listener.Accept()

		if err != nil {

			if err == net.ErrClosed {
				break
			} else {
				log.Printf("Listening Accept Error: %v", err)
			}
		}

		go s.handleConn(connection)

	}

}

func (s *Server) handleConn(connection net.Conn) {

	defer connection.Close()

	// Create a buffer to read the request
	buffer := make([]byte, 1024)
	_, err := connection.Read(buffer)
	if err != nil {
		log.Printf("Error reading request: %v", err)
		return
	}

	responseBody := []byte("Hello World\r\n")
	bodyLen := len(responseBody)

	// response := []byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nHello World!")

	responseHeader := response.GetDefaultHeaders(bodyLen)

	// sending the status line

	err = response.WriteStatusLine(connection, "200")

	if err != nil {
		log.Printf("Error at sending status line: %v", err)
		return

	}

	err = response.WriteHeaders(connection, responseHeader)

	if err != nil {
		log.Printf("Error at writing header: %v", err)
	}

	_, err = connection.Write(responseBody)
	if err != nil {
		log.Printf("Error at writing body: %v", err)
	}

}
