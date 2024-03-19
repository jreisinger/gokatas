package shift_test

import (
	"bytes"
	"testing"

	"github.com/jreisinger/gokatas/shift"
)

var testcases = []struct {
	plaintext  []byte
	ciphertext []byte
	key        byte
}{
	{[]byte("HAL"), []byte("IBM"), 1},
	{[]byte(""), []byte(""), 1},
	{[]byte{}, []byte{}, 1},
	{nil, nil, 1},
}

func TestEncrypt(t *testing.T) {
	for _, tc := range testcases {
		got := shift.Encrypt(tc.plaintext, tc.key)
		if !bytes.Equal(tc.ciphertext, got) {
			t.Errorf("want %v, got %v", tc.ciphertext, got)
		}
	}
}

func TestDecrypt(t *testing.T) {
	for _, tc := range testcases {
		got := shift.Decrypt(tc.ciphertext, tc.key)
		if !bytes.Equal(tc.plaintext, got) {
			t.Errorf("want %v, got %v", tc.plaintext, got)
		}
	}
}
