package handlers

import (
	"io"

	"github.com/akhand08/http-server-golang/internal/server"
)

var coffeeHandler = func(w io.Writer) *server.HandlerError {

	responseBody := "tan tana tan....Helloooo My dear, Here is your cappachino special coffee\r\n"

	_, err := w.Write([]byte(responseBody))
	if err != nil {
		return &server.HandlerError{StatusCode: "500", Message: "Server Problem\n"}
	}

	return nil

}
