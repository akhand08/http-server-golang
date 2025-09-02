package headers

import (
	"bytes"
	"errors"
	"strings"
)

const CRLF = "\r\n"
const DoubleCRLF = "\r\n\r\n"

type Headers map[string]string

func NewHeaders() Headers {
	return make(Headers)
}

func (h Headers) Parse(data []byte) (n int, done bool, err error) {

	newLineIndex := 0

	if newLineIndex = bytes.Index(data, []byte(DoubleCRLF)); newLineIndex != -1 {

		err = h.mapper(data[:newLineIndex])

		if err != nil {
			return 0, false, err
		}

		return newLineIndex + 4, true, nil

	}

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

	key, err := keyValidators(key)
	if err != nil {
		return err
	}
	value = strings.TrimSpace(value)

	if _, exist := h[key]; exist == true {
		h[key] = h[key] + ", " + value
	} else {
		h[key] = value
	}

	// fmt.Println("key: ", key, "value: ", value)

	return nil

}

func keyValidators(key string) (string, error) {
	keyLen := len(key)
	if keyLen < 1 {
		return "", errors.New("Invalid header key length")
	}

	if key[keyLen-1] == ' ' {
		return "", errors.New("Invalid key")
	}

	key = strings.TrimSpace(key)

	for i := 0; i < keyLen; i++ {

		if int(key[i]) >= 48 && int(key[i]) <= 57 {
			continue
		}

		if int(key[i]) >= 65 && int(key[i]) <= 90 {
			continue
		}

		if int(key[i]) >= 97 && int(key[i]) <= 123 {
			continue
		}

		if key[i] == '!' || key[i] == '#' || key[i] == '$' || key[i] == '%' ||
			key[i] == '&' || key[i] == '\'' || key[i] == '*' || key[i] == '+' ||
			key[i] == '-' || key[i] == '.' || key[i] == '^' || key[i] == '_' ||
			key[i] == '`' || key[i] == '|' || key[i] == '~' {

			continue

		}

		return "", errors.New("Invalid character in header field name")

	}

	key = strings.ToLower(key)

	return key, nil
}
