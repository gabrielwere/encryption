package vigenerecipher

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

func Encrypt(plaintext string, key string) (string, error) {

	if plaintext == "" {
		return "", nil
	}

	plaintext = strings.ToUpper(plaintext)
	key = strings.ToUpper(key)
	ciphertext := ""
	i := 0

	for _, letter := range plaintext {

		if unicode.IsSpace(letter) {
			ciphertext += " "
		} else if unicode.IsLetter(letter) {
			//column value in the tabula recta
			plain_letter_index := int(letter - 'A')

			//row value in the tabula recta
			key_letter_index := int(key[i%len(key)] - 'A')
			i++

			//use the instersection of row and column value as the index
			ciphertext += alphabet[(plain_letter_index+key_letter_index)%26]
		} else {
			return "", fmt.Errorf("Input is not part of the English Alphabet")
		}
	}

	return ciphertext, nil
}

func Decrypt(ciphertext string, key string) (string, error) {

	ciphertext = strings.ToUpper(ciphertext)
	key = strings.ToUpper(key)

	plaintext := ""
	i := 0

	for _, letter := range ciphertext {

		if unicode.IsSpace(letter) {
			plaintext += " "
		} else if unicode.IsLetter(letter) {
			//instersection value in the tabula recta
			cipher_letter_index := int(letter - 'A')

			//row value in the tabula recta
			key_letter_index := int(key[i%len(key)] - 'A')
			i++

			//use the column value as the index
			index := (cipher_letter_index - key_letter_index) % 26

			//no ternary operator for go??!
			if index < 0 {
				//its a negative number so subtracting it would make it positive
				index = 26 + index
			}
			plaintext += alphabet[index]
		} else {
			return "", fmt.Errorf("Input is not part of the English Alphabet")
		}
	}

	return plaintext, nil
}
