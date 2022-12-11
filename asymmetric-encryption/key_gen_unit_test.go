package asymmetricencryption_test

import (
	"testing"

	rsa "github.com/gabrielwere/encryption/asymmetric-encryption"
)

func TestGenerateKey(t *testing.T) {
	tests := []struct {
		name    string
		prime1  int64
		prime2  int64
		d       int64
		e       int64
		n       int64
		wantErr bool
	}{
		//all the test cases are calculated by hand
		//added/wrote as many as I could find/calculate
		{
			name:    "case_1",
			prime1:  3,
			prime2:  11,
			n:       33,
			e:       7,
			d:       3,
			wantErr: false,
		},
		{
			name:    "case_2",
			prime1:  2,
			prime2:  11,
			n:       22,
			e:       3,
			d:       7,
			wantErr: false,
		},
		{
			name:    "case_3",
			prime1:  2,
			prime2:  7,
			n:       14,
			e:       5,
			d:       11,
			wantErr: false,
		},
		{
			name:    "case_4",
			prime1:  17,
			prime2:  7,
			n:       119,
			e:       5,
			d:       77,
			wantErr: false,
		},
		{
			name:    "case_5",
			prime1:  53,
			prime2:  59,
			n:       3127,
			e:       3,
			d:       2011,
			wantErr: false,
		},
		{
			name:    "case_6",
			prime1:  7,
			prime2:  13,
			n:       91,
			e:       5,
			d:       29,
			wantErr: false,
		},
		{
			name:    "case_7",
			prime1:  17,
			prime2:  13,
			n:       221,
			e:       5,
			d:       77,
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotD, gotE, gotN, err := rsa.GenerateKey(tc.prime1, tc.prime2)

			if !tc.wantErr && err != nil {
				t.Errorf("Did not expect to get error %v\n", err)
			}
			if gotD != tc.d || gotE != tc.e || gotN != tc.n {
				t.Errorf("Got values of d:%v e:%v n:%v did not match expected values %v %v %v\n", gotD, gotE, gotN, tc.d, tc.e, tc.n)
			}
		})
	}
}
func TestIsPrime(t *testing.T) {
	tests := []struct {
		name   string
		number int64
		output bool
	}{
		{
			name:   "prime_1",
			number: 17,
			output: true,
		},
		{
			name:   "prime_2",
			number: 967,
			output: true,
		},
		{
			name:   "composite_1",
			number: 888,
			output: false,
		},
		{
			name:   "composite_2",
			number: 1100,
			output: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			isprime := rsa.IsPrime(tc.number)

			if isprime != tc.output {
				t.Errorf("Output %v does not match expected output %v", tc.output, isprime)
			}
		})
	}
}
func TestGCD(t *testing.T) {

	tests := []struct {
		name string
		num1 int64
		num2 int64
		gcd  int64
	}{
		{
			name: "composite_numbers_1",
			num1: 270,
			num2: 192,
			gcd:  6,
		},
		{
			name: "prime_numbers_1",
			num1: 83,
			num2: 71,
			gcd:  1,
		},
		{
			name: "composite_numbers_2",
			num1: 6,
			num2: 9,
			gcd:  3,
		},
		{
			name: "prime_numbers_2",
			num1: 3,
			num2: 17,
			gcd:  1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			got := rsa.GCD(tc.num1, tc.num2)

			if got != tc.gcd {
				t.Errorf("output %v does not match expected output %v", got, tc.gcd)
			}
		})
	}
}
