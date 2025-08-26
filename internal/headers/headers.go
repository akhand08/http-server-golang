package headers

import (
	"bytes"
	"errors"
	"strings"
)

const CRLF = "\r\n"

type Headers map[string]string

func NewHeaders() Headers {
	return make(Headers)
}

func (h Headers) Parse(data []byte) (n int, done bool, err error) {

	newLineIndex := 0

	if newLineIndex = bytes.Index(data, []byte(CRLF)); newLineIndex == -1 {
		return 0, false, nil
	}

	if newLineIndex == 0 {
		return 2, true, nil
	}

	err = h.mapper(data[:newLineIndex])

	if err != nil {
		return 0, false, err
	}

	return newLineIndex + 2, false, nil
}

func (h Headers) mapper(data []byte) error {

	strData := string(data)

	parts := strings.SplitN(strData, ":", 2)

	if len(parts) != 2 {
		return errors.New("Invalid key-value pair")
	}

	key := parts[0]
	value := parts[1]

	keyLen := len(key)

	if key[keyLen-1] == ' ' {
		return errors.New("Invalid key")
	}

	key = strings.TrimSpace(key)
	value = strings.TrimSpace(value)

	h[key] = value

	return nil

}
