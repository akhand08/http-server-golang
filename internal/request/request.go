package request

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"unicode"
)

const bufferSize = 4096
const chunkSize = 8

type State int

const (
	Initialized State = iota
	Done
)

type Request struct {
	RequestLine RequestLine
	State       State
}

type RequestLine struct {
	HttpVersion   string
	Method        string
	RequestTarget string
}

func (r *Request) parse(buffer []byte) (n int, err error) {

	newLineIndex := 0

	if newLineIndex = bytes.Index(buffer, []byte("\r\n")); newLineIndex == -1 {
		return 0, nil
	}

	line := buffer[:newLineIndex]
	newRequestLine, err := parseRequestLine(string(line))

	if err != nil {
		return 0, err
	}

	r.RequestLine = *newRequestLine
	r.State = Done

	return len(line) + 2, nil

}

func RequestFromReader(reader io.Reader) (*Request, error) {

	buffer := make([]byte, bufferSize, bufferSize)
	readToIndex := 0
	newRequest := &Request{State: Initialized}

	for newRequest.State != Done {

		chunk := make([]byte, 8)

		n, err := reader.Read(chunk)

		if err != nil {
			return nil, err
		}

		copy(buffer[readToIndex:], chunk)
		readToIndex += n

		consumed, err := newRequest.parse(buffer)

		if err != nil {
			return nil, err
		}

		if consumed > 0 {

			readToIndex -= consumed
			buffer = buffer[consumed+1:]
		}

	}

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

	if httpVersion != "1.1" {
		return nil, errors.New("Invalid http version")
	}

	newRequestLine := &RequestLine{HttpVersion: httpVersion, Method: method, RequestTarget: requestTarget}

	return newRequestLine, nil

}
