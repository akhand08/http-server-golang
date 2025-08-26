package request

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"unicode"

	"github.com/akhand08/http-server-golang/internal/headers"
)

const bufferSize = 4096
const chunkSize = 8
const CRLF = "\r\n"

type State int

const (
	ParsingRequestLine State = iota
	ParsingRequestLineDone
	ParsingHeader
	ParsingHeaderDone
	ParsingBody
)

type Request struct {
	RequestLine   RequestLine
	RequestHeader headers.Headers
	State         State
}

type RequestLine struct {
	HttpVersion   string
	Method        string
	RequestTarget string
}

func (r *Request) parse(buffer []byte) (n int, err error) {

	switch r.State {
	case ParsingRequestLine:
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

		return len(line) + 2, nil
	case ParsingRequestLineDone:
		r.State += 1
		return 0, nil
	case ParsingHeaderDone:
		r.State -= 1
		return 0, nil
	case ParsingHeader:

		byteConsumed, isEnd, err := r.RequestHeader.Parse(buffer)

		if err != nil {
			return 0, err
		}

		if isEnd == true {
			r.State += 1
			return 2, nil
		}

		return byteConsumed, nil

	}

	return 0, errors.New("Failed to Parsing")

}

func RequestFromReader(reader io.Reader) (*Request, error) {

	buffer := make([]byte, bufferSize, bufferSize)
	readToIndex := 0
	newRequest := &Request{State: ParsingRequestLine}
	newHeader := headers.NewHeaders()
	newRequest.RequestHeader = newHeader

	for newRequest.State != ParsingBody {

		chunk := make([]byte, 8)

		n, err := reader.Read(chunk)

		if err != nil {

			if err == io.EOF {
				return newRequest, nil
			}
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
			buffer = buffer[consumed:]

			newRequest.State = newRequest.State + 1 // moving to the next stage
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
