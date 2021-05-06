package rotroh

import (
	"encoding/base64"
	"strings"
)

const (
	rot13Src = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	rot13Des = "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"
)

// Base64String does a Base64 encoding or a Base64 decoding
func Base64String(input string) (result string, err error) {
	if IsBase64String(input) {
		// log.Println("Input is Base64 encoded")
		b, err := base64.StdEncoding.DecodeString(input)
		if err == nil {
			result = string(b)
		}
		if err != nil {
			return "", err
		}
	} else {
		// log.Println("Input is not Base64 encoded")
		result = base64.StdEncoding.EncodeToString([]byte(input))
	}

	return result, err
}

// Rot13String does a ROT-13 transform on a string
func Rot13String(input string) string {
	result := ""
	for _, r := range input {
		s := string(r)
		i := strings.Index(rot13Src, s)
		// log.Printf("Rot13String() | s: %q | i: %d\n", s, i)
		if i > -1 {
			result += string(rot13Des[i])
		} else {
			result += s
		}
	}
	return result
}

// Rot47String does a ROT-47 transform on a string
func Rot47String(input string) string {
	result := ""
	for _, r := range input {
		i := int(r)
		x := i + 47
		if x > 126 {
			x = i - 47
		}
		s := string(x)
		// log.Printf("Rot47String() | r: %q | i: %3d | x: %3d | s: %q\n", r, i, x, s)
		result += s
	}
	return result
}

// IsBase64String returns true if the input was a already base64 encoded
func IsBase64String(input string) bool {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		// log.Printf("IsBase64String(%s) | Error: %s\n", input, err.Error())
		return false
	}
	encoded := base64.StdEncoding.EncodeToString([]byte(decoded))
	if encoded == input {
		return true
	}
	return false
}

// RotRoh47String does a ROT-47 transform then Base64 encoding or a Base64 decoding then a ROT-47 transform
func RotRoh47String(input string) (result string, err error) {
	if IsBase64String(input) {
		// log.Println("Input is Base64 encoded")
		decoded, err := base64.StdEncoding.DecodeString(input)
		if err != nil {
			return result, err
		}
		result = Rot47String(string(decoded))
	} else {
		// log.Println("Input is not Base64 encoded")
		msg := Rot47String(input)
		result = base64.StdEncoding.EncodeToString([]byte(msg))
	}

	return result, err
}
