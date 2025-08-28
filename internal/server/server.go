package server

import (
	"log"
	"net"
	"strconv"
	"sync/atomic"
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

	response := []byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 12\r\n\r\nHello World!")

	_, err = connection.Write(response)

	if err != nil {
		log.Printf("Error at writing: %v", err)
	}

}
