package vigenerecipher_test

import (
	"testing"

	cipher "github.com/gabrielwere/encryption/vigenere-cipher"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		key     string
		output  string
		wantErr bool
	}{
		{
			name:    "lowercase",
			input:   "gabriel",
			key:     "lemon",
			output:  "RENFVPP",
			wantErr: false,
		},
		{
			name:    "uppercase",
			input:   "GABRIEL",
			key:     "lemon",
			output:  "RENFVPP",
			wantErr: false,
		},
		{
			name:    "mixedcase",
			input:   "GAbriEl",
			key:     "lemon",
			output:  "RENFVPP",
			wantErr: false,
		},
		{
			name:    "sentence_with_spaces",
			input:   "attack at dawn",
			key:     "lemon",
			output:  "LXFOPV EF RNHR",
			wantErr: false,
		},
		{
			name:    "using_were_as_key",
			input:   "gabriel",
			key:     "WERE",
			output:  "CESVEIC",
			wantErr: false,
		},
		{
			name:    "using_null_as_key",
			input:   "attack at dawn",
			key:     "null",
			output:  "NNELPE LE QUHY",
			wantErr: false,
		},
		{
			name:    "alphanumeric",
			input:   "0012dawn",
			key:     "lemon",
			output:  "",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ciphertext, err := cipher.Encrypt(tc.input, tc.key)

			if tc.wantErr && err == nil {
				t.Errorf("Expected an error but did not get one\n")
			}

			if !tc.wantErr && err != nil {
				t.Errorf("Did not expect error %v but got it\n", err)
			}

			if tc.output != ciphertext {
				t.Errorf("Output %v did not match expected out %v", ciphertext, tc.output)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		key     string
		output  string
		wantErr bool
	}{
		{
			name:    "lowercase",
			input:   "renfvpp",
			key:     "lemon",
			output:  "GABRIEL",
			wantErr: false,
		},
		{
			name:    "uppercase",
			input:   "RENFVPP",
			key:     "lemon",
			output:  "GABRIEL",
			wantErr: false,
		},
		{
			name:    "mixedcase",
			input:   "RenFvPp",
			key:     "lemon",
			output:  "GABRIEL",
			wantErr: false,
		},
		{
			name:    "sentence_with_spaces",
			input:   "lxfopv ef rnhr",
			key:     "lemon",
			output:  "ATTACK AT DAWN",
			wantErr: false,
		},
		{
			name:    "using_were_as_key",
			input:   "CESVEIC",
			key:     "were",
			output:  "GABRIEL",
			wantErr: false,
		},
		{
			name:    "using_null_as_key",
			input:   "NNELPE LE QUHY",
			key:     "null",
			output:  "ATTACK AT DAWN",
			wantErr: false,
		},
		{
			name:    "alphanumeric",
			input:   "0012dawn",
			key:     "lemon",
			output:  "",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			plaintext, err := cipher.Decrypt(tc.input, tc.key)

			if tc.wantErr && err == nil {
				t.Errorf("Expected an error but did not get one\n")
			}

			if !tc.wantErr && err != nil {
				t.Errorf("Did not expect error %v but got it\n", err)
			}

			if tc.output != plaintext {
				t.Errorf("Output %v did not match expected out %v", plaintext, tc.output)
			}
		})
	}
}
