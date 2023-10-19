// Cryptopals shows basic crypto operations. Always operate on raw bytes, never
// on encoded strings. Only use hex and base64 for pretty-printing.
//
// Adapted from https://github.com/0xfe/cryptopals.
//
// Level: beginner
// Topics: encoding, xor
package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
)

// HexToBase64 decodes hex string into base64 string. (Set 1, challenge 1.)
func HexToBase64(h string) (string, error) {
	b, err := hex.DecodeString(h)
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(b), nil
}

// FixedXOR does bitwise XOR of two hex strings. The strings must be of equal
// length. (Set 1, challenge 2.)
func FixedXOR(h1, h2 string) (string, error) {
	b1, err := hex.DecodeString(h1)
	if err != nil {
		return "", err
	}
	b2, err := hex.DecodeString(h2)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(b1))
	for i := range b1 {
		out[i] = b1[i] ^ b2[i]
	}
	return hex.EncodeToString(out), nil
}
