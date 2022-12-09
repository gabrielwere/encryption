package atbashcipher

import (
	"fmt"
	"strings"
	"unicode"
)

var alphabet []string = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z",
}

func Encrypt(plaintext string) (string, error) {

	ciphertext := ""
	if plaintext == "" {
		return "", nil
	}

	plaintext = strings.ToUpper(plaintext)
	for _, letter := range plaintext {
		if unicode.IsSpace(letter) {
			ciphertext += " "
		} else if unicode.IsLetter(letter) {
			ciphertext += alphabet['Z'-letter]
		} else {
			return "", fmt.Errorf("Input is not part of English alphabet")
		}
	}

	return ciphertext, nil
}
