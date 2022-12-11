package asymmetricencryption

import (
	"fmt"
)

// demonstrates the RSA key generation algorithm
func GenerateKey(prime1, prime2 int64) (d, e, n int64, err error) {

	if !IsPrime(prime1) || !IsPrime(prime2) {
		return 0, 0, 0, fmt.Errorf("numbers must be prime numbers")
	}
	n = prime1 * prime2

	phiOfN := (prime1 - 1) * (prime2 - 1)

	//find the smallest number,e that is coprime with phiOfN
	//and coprime with n
	var i int64
	for i = 2; i < phiOfN; i++ {
		if GCD(phiOfN, i) == 1 && GCD(n, i) == 1 {
			e = i
			break
		}
	}

	//find the smallest number,d such that (e *d) mod phiOfN = 1
	for i = 2; ; i++ {
		if (i*e)%phiOfN == 1 && i != e {
			d = i
			break
		}
	}
	return d, e, n, nil

}

func IsPrime(num int64) bool {

	var i int64
	for i = 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// GCD function of two numbers
// uses the euclidean algorithm implemented iteratively
// TODO : implement recursively
func GCD(num1, num2 int64) (gcd int64) {

	for num2 > 0 {
		gcd = num2
		num2 = (num1 % num2)
		num1 = gcd
	}
	return gcd
}
