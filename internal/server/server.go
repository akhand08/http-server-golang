package server

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"sync/atomic"

	"github.com/akhand08/http-server-golang/internal/request"
	"github.com/akhand08/http-server-golang/internal/response"
)

type Handler func(w io.Writer, req *request.Request) *HandlerError

type HandlerError struct {
	StatusCode string
	Message    string
}

type Server struct {
	listener  net.Listener
	isRunning atomic.Bool
	handler   Handler
}

func Serve(port int, handler Handler) (*Server, error) {

	listenAt := ":" + strconv.Itoa(port)

	listener, err := net.Listen("tcp", listenAt)

	if err != nil {
		return nil, err
	}

	server := &Server{listener: listener, handler: handler}
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

	httpRequest, err := request.RequestFromReader(connection)
	if err != nil {
		log.Printf("Error reading request: %v", err)
		return
	}

	var responseBody bytes.Buffer

	handlerError := s.handler(&responseBody, httpRequest)
	fmt.Println("Response Body: ", responseBody)
	fmt.Println(handlerError)

	if handlerError != nil {

		s.writeError(connection, handlerError)

	} else {
		s.ResponseWriter(connection, &responseBody)
	}

}

func (s *Server) ResponseWriter(w io.Writer, responseBody *bytes.Buffer) {

	bodyLen := len(responseBody.Bytes())
	fmt.Println(bodyLen)

	responseHeaders := response.GetDefaultHeaders(bodyLen)
	err := response.WriteStatusLine(w, response.Ok)

	if err != nil {
		log.Printf("Error at sending status line: %v", err)
		return

	}

	err = response.WriteHeaders(w, responseHeaders)

	if err != nil {
		log.Printf("Error at writing header: %v", err)
		return
	}

	_, err = w.Write([]byte("\r\n"))
	if err != nil {
		log.Printf("Error at writing empty line: %v", err)
		return
	}

	_, err = w.Write(responseBody.Bytes())
	if err != nil {
		log.Printf("Error at writing body: %v", err)
		return
	}

}

func (s *Server) writeError(conn net.Conn, error *HandlerError) {

	statusCode := error.StatusCode
	reasonPhrase := error.Message

	response := fmt.Sprintf(
		"HTTP/1.1 %s %s\r\nContent-Length: 0\r\nConnection: close\r\n\r\n",
		statusCode, reasonPhrase,
	)
	_, err := conn.Write([]byte(response))
	if err != nil {
		log.Printf("Failed writing error response: %v", err)
	}
}
