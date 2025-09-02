package response

import (
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/akhand08/http-server-golang/internal/headers"
)

type statusCode string

const (
	Ok          statusCode = "200"
	BadReq      statusCode = "400"
	ServerError statusCode = "500"
)

var reasonPhrase = map[statusCode][]byte{
	Ok:          []byte("HTTP/1.1 200 OK\r\n"),
	BadReq:      []byte("HTTP/1.1 400 Bad Request\r\n"),
	ServerError: []byte("HTTP/1.1 500 Internal Server Error\r\n"),
}

func WriteStatusLine(w io.Writer, statusCode statusCode) error {

	_, err := w.Write([]byte(reasonPhrase[statusCode]))
	return err

}

func GetDefaultHeaders(contentLen int) headers.Headers {
	responseHeader := headers.NewHeaders()

	responseHeader["Content-Length"] = strconv.Itoa(contentLen)
	responseHeader["Connection"] = "close"
	responseHeader["Content-Type"] = "text/plain"
	responseHeader["Date"] = strings.Replace(time.Now().UTC().Format(time.RFC1123), "UTC", "GMT", 1)
	return responseHeader
}

func WriteHeaders(w io.Writer, headers headers.Headers) error {

	for key, val := range headers {
		byteFieldLine := []byte((key + ": " + val + "\r\n"))
		_, err := w.Write(byteFieldLine)
		if err != nil {
			return err
		}
	}

	// _, err := w.Write([]byte("\r\n"))
	return nil

}
