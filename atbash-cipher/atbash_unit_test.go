package atbashcipher_test

import (
	"testing"

	cipher "github.com/gabrielwere/encryption/atbash-cipher"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		output  string
		wantErr bool
	}{
		{
			name:    "encrypt",
			input:   "gabriel",
			output:  "TZYIRVO",
			wantErr: false,
		},
		{
			name:    "encrypt_with_spaces",
			input:   "attack at dawn",
			output:  "ZGGZXP ZG WZDM",
			wantErr: false,
		},
		{
			name:    "decrypt",
			input:   "TZYIRVO",
			output:  "GABRIEL",
			wantErr: false,
		},
		{
			name:    "decrypt_with_spaces",
			input:   "ZGGZXP ZG WZDM",
			output:  "ATTACK AT DAWN",
			wantErr: false,
		},
		{
			name:    "encrypt_uppercase",
			input:   "GaBRIel",
			output:  "TZYIRVO",
			wantErr: false,
		},
		{
			name:    "invalid_input",
			input:   "12gabr3",
			output:  "",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			ciphertext, err := cipher.Encrypt(tc.input)

			if tc.wantErr && err == nil {
				t.Errorf("expected an error but did not get one\n")
			}

			if !tc.wantErr && err != nil {
				t.Errorf("got error %v but was not expecting it\n", err)
			}

			if tc.output != ciphertext {
				t.Errorf("output %v does not match expected output %v\n", tc.output, ciphertext)
			}
		})
	}
}
