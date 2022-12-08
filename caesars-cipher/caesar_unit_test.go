package caesarscipher_test

import (
	"testing"

	cipher "github.com/gabrielwere/encryption/caesars-cipher"
)

func TestEncrypt(t *testing.T) {

	tests := []struct {
		name    string
		input   string
		key     int
		output  string
		wantErr bool
	}{
		{
			name:    "uppercase",
			input:   "GABRIEL",
			key:     3,
			output:  "JDEULHO",
			wantErr: false,
		},
		{
			name:    "lowercase",
			input:   "gabriel",
			key:     3,
			output:  "JDEULHO",
			wantErr: false,
		},
		{
			name:    "mixedcase",
			input:   "gABriEl",
			key:     3,
			output:  "JDEULHO",
			wantErr: false,
		},
		{
			name:    "using_4_as_key",
			input:   "gabriel",
			key:     4,
			output:  "KEFVMIP",
			wantErr: false,
		},
		{
			name:    "using_5_as_key",
			input:   "gabriel",
			key:     5,
			output:  "LFGWNJQ",
			wantErr: false,
		},
		{
			name:    "sentence_with_spaces",
			input:   "attack at dawn",
			output:  "DWWDFN DW GDZQ",
			key:     3,
			wantErr: false,
		},
		{
			name:    "empty_string",
			input:   " ",
			output:  " ",
			key:     3,
			wantErr: false,
		},
		{
			name:    "alphanumeric",
			input:   "gab001",
			output:  "",
			key:     3,
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ciphertext, err := cipher.Encrypt(tc.input, tc.key)

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

func TestDecrypt(t *testing.T) {

	tests := []struct {
		name    string
		input   string
		key     int
		output  string
		wantErr bool
	}{
		{
			name:    "uppercase",
			input:   "JDEULHO",
			output:  "GABRIEL",
			key:     3,
			wantErr: false,
		},
		{
			name:    "lowercase",
			input:   "jdeulho",
			output:  "GABRIEL",
			key:     3,
			wantErr: false,
		},
		{
			name:    "mixedcase",
			input:   "JDEulhO",
			key:     3,
			output:  "GABRIEL",
			wantErr: false,
		},
		{
			name:    "using_4_as_key",
			input:   "KEFVMIP",
			key:     4,
			output:  "GABRIEL",
			wantErr: false,
		},
		{
			name:    "using_5_as_key",
			input:   "LFGWNJQ",
			key:     5,
			output:  "GABRIEL",
			wantErr: false,
		},
		{
			name:    "sentence_with_spaces",
			input:   "DWWDFN DW GDZQ",
			output:  "ATTACK AT DAWN",
			key:     3,
			wantErr: false,
		},
		{
			name:    "empty_string",
			input:   " ",
			output:  " ",
			key:     3,
			wantErr: false,
		},
		{
			name:    "alphanumeric",
			input:   "gab001",
			output:  "",
			key:     3,
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			plaintext, err := cipher.Decrypt(tc.input, tc.key)

			if tc.wantErr && err == nil {
				t.Errorf("expected an error but did not get one\n")
			}

			if !tc.wantErr && err != nil {
				t.Errorf("got error %v but was not expecting it\n", err)
			}

			if tc.output != plaintext {
				t.Errorf("output %v does not match expected output %v\n", tc.output, plaintext)
			}
		})
	}
}
