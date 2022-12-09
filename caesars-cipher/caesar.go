package caesarscipher

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

func Encrypt(plaintext string, key int) (string, error) {

	ciphertext := ""
	plaintext = strings.ToUpper(plaintext)
	for _, letter := range plaintext {
		if unicode.IsSpace(letter) {
			ciphertext += " "
		} else if num := ((int(letter) - 'A') + key) % 26; unicode.IsLetter(letter) && num >= 0 && num <= 25 {
			ciphertext += alphabet[num]
		} else {
			return "", fmt.Errorf("Input is not in the alphabet")
		}
	}

	return ciphertext, nil
}

func Decrypt(ciphertext string, key int) (string, error) {

	plaintext := ""
	ciphertext = strings.ToUpper(ciphertext)
	for _, letter := range ciphertext {
		if unicode.IsSpace(letter) {
			plaintext += " "
		} else if num := ((int(letter) - 'A') - key) % 26; unicode.IsLetter(letter) && num >= 0 && num <= 25 {
			plaintext += alphabet[num]
		} else {
			return "", fmt.Errorf("Input is not in the alphabet")
		}
	}

	return plaintext, nil
}
