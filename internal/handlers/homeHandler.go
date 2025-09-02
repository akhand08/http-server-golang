package handlers

import (
	"io"

	"github.com/akhand08/http-server-golang/internal/server"
)

var homeHandler = func(w io.Writer) *server.HandlerError {

	responseBody := "Hurreh, a warm welcome to Home\r\n"

	_, err := w.Write([]byte(responseBody))
	if err != nil {
		return &server.HandlerError{StatusCode: "400", Message: "Not Found\n"}
	}

	return nil

}
