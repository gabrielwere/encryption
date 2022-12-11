package asymmetricencryption

import (
	"math"
)

// just encrypt numbers as they are easier to calculate
// the same concept can be applied on strings given a selected alphabet scheme
func Encrypt(plaintext, prime1, prime2 int64) (int64, error) {

	_, e, n, err := GenerateKey(prime1, prime2)

	if err != nil {
		return 0, err
	}
	var ciphertext int64
	ciphertext = int64(math.Pow(float64(plaintext), float64(e))) % n

	return ciphertext, nil
}

func Decrypt(ciphertext, prime1, prime2 int64) (int64, error) {

	d, _, n, err := GenerateKey(prime1, prime2)

	if err != nil {
		return 0, err
	}
	var plaintext int64
	plaintext = int64(math.Pow(float64(ciphertext), float64(d))) % n

	return plaintext, nil
}
