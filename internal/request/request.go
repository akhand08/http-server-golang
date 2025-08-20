package request

import (
	"bytes"
	"errors"
	"io"
	"log"
	"strings"
	"unicode"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	Method        string
	RequestTarget string
}

func RequestFromReader(reader io.Reader) (*Request, error) {

	data, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal("cannot read error: ", err)
	}

	requestLine := ""

	if newLineIndex := bytes.Index(data, []byte("\r\n")); newLineIndex != -1 {

		requestLine = string(data[:newLineIndex])

	}
	newRequestLine, err := parseRequestLine(requestLine)

	if err != nil {
		return nil, err
	}

	newRequest := &Request{RequestLine: *newRequestLine}

	return newRequest, nil

}

func parseRequestLine(requestLine string) (*RequestLine, error) {
	parts := strings.SplitN(requestLine, " ", 3)

	if len(parts) != 3 {
		return nil, errors.New("Invalid Requstline")
	}

	method := parts[0]
	requestTarget := parts[1]
	httpVersion := parts[2][5:]

	for _, letr := range method {
		if !unicode.IsUpper(letr) {
			return nil, errors.New("Method is not in uppercase")
		}
	}

	newRequestLine := &RequestLine{HttpVersion: httpVersion, Method: method, RequestTarget: requestTarget}

	return newRequestLine, nil

}
