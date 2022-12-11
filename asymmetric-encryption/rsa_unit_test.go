package asymmetricencryption_test

import (
	"testing"

	rsa "github.com/gabrielwere/encryption/asymmetric-encryption"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name    string
		input   int64
		prime1  int64
		prime2  int64
		output  int64
		wantErr bool
	}{
		{
			name:    "case_1",
			input:   2,
			prime1:  3,
			prime2:  11,
			output:  29,
			wantErr: false,
		},
		{
			name:    "case_2",
			input:   2,
			prime1:  2,
			prime2:  7,
			output:  4,
			wantErr: false,
		},
		{
			name:    "case_3",
			input:   3,
			prime1:  2,
			prime2:  11,
			output:  5,
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			ciphertext, err := rsa.Encrypt(tc.input, tc.prime1, tc.prime2)

			if !tc.wantErr && err != nil {
				t.Errorf("Did not expect to get error %v", err)
			}

			if ciphertext != tc.output {
				t.Errorf("Output %v and Expected output %v do not match", ciphertext, tc.output)
			}
		})
	}
}
func TestDecrypt(t *testing.T) {
	tests := []struct {
		name    string
		input   int64
		prime1  int64
		prime2  int64
		output  int64
		wantErr bool
	}{
		{
			name:    "case_1",
			input:   29,
			prime1:  3,
			prime2:  11,
			output:  2,
			wantErr: false,
		},
		{
			name:    "case_2",
			input:   4,
			prime1:  2,
			prime2:  7,
			output:  2,
			wantErr: false,
		},
		{
			name:    "case_3",
			input:   5,
			prime1:  2,
			prime2:  11,
			output:  3,
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			plaintext, err := rsa.Decrypt(tc.input, tc.prime1, tc.prime2)

			if !tc.wantErr && err != nil {
				t.Errorf("Did not expect to get error %v", err)
			}

			if plaintext != tc.output {
				t.Errorf("Output %v and Expected output %v do not match", plaintext, tc.output)
			}
		})
	}
}
