package handlers

import (
	"io"

	"github.com/akhand08/http-server-golang/internal/request"
	"github.com/akhand08/http-server-golang/internal/server"
)

var RootHandler server.Handler = func(w io.Writer, req *request.Request) *server.HandlerError {

	switch req.RequestLine.RequestTarget {
	case "/home":
		return homeHandler(w)

	case "coffee":
		return coffeeHandler(w)

	default:
		return &server.HandlerError{StatusCode: "400", Message: "Not Found\r\n"}
	}

}
