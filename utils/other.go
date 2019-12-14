package utils

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"piedpiper/utils/log"
	"regexp"
)

// Close error checking for defer close
func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// ValidateEmail .
func ValidateEmail(email string) bool {
	Re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?" +
		"(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return Re.MatchString(email)
}

// SHA512 hashes using SHA512 algorithm
func SHA512(text string) string {
	algorithm := sha512.New()
	_, err := algorithm.Write([]byte(text))
	if err != nil {
		return ""
	}
	return hex.EncodeToString(algorithm.Sum(nil))
}

// Base64Encode ...
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Base64Decode ...
func Base64Decode(str string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", fmt.Errorf("can't decode the string to base64")
	}
	return string(data), nil
}
