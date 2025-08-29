package response

import (
	"io"
	"strconv"

	"github.com/akhand08/http-server-golang/internal/headers"
)

type statusCode string

const (
	Ok          statusCode = "200"
	BadReq      statusCode = "400"
	ServerError statusCode = "500"
)

var reasonPhrase = map[statusCode][]byte{
	"200": []byte("HTTP/1.1 200 OK\r\n"),
	"400": []byte("HTTP/1.1 400 Bad Request\r\n"),
	"500": []byte("HTTP/1.1 500 Internal Server Error\r\n"),
}

func WriteStatusLine(w io.Writer, statusCode statusCode) error {

	_, err := w.Write([]byte(reasonPhrase[statusCode]))
	return err

}

func GetDefaultHeaders(contentLen int) headers.Headers {
	responseHeader := headers.NewHeaders()

	responseHeader["Content=Lenth"] = strconv.Itoa(contentLen)
	responseHeader["Connection"] = "close"
	responseHeader["Content-Type"] = "text/plain"

	return responseHeader
}

func WriteHeaders(w io.Writer, headers headers.Headers) error {

	for key, val := range headers {
		byteFieldLine := []byte((key + ":" + val + "\r\n"))
		_, err := w.Write(byteFieldLine)
		if err != nil {
			return err
		}
	}
	_, err := w.Write([]byte("\r\n"))
	return err

}
