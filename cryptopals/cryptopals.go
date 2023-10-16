// Cryptopals shows basic crypto operations.
//
// Always operate on raw bytes, never on encoded strings. Only use hex and
// base64 for pretty-printing.
//
// Adapted from https://github.com/0xfe/cryptopals.
//
// Level: intermediate
// Topics: encoding, xor
package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
)

// HexToBase64 decodes hex string h into bytes and then bytes into base64
// string. (Set 1, challenge 1.)
func HexToBase64(h string) (string, error) {
	data, err := hex.DecodeString(h)
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(data), nil
}

// FixedXOR gets the bytes represented by h1 and h2 hex strings. Then it does
// XOR bitwise operation on these two byte slices, byte by byte. H1 and h2 must
// be of equal length. (Set 1, challenge 2.)
func FixedXOR(h1, h2 string) (string, error) {
	data1, err := hex.DecodeString(h1)
	if err != nil {
		return "", err
	}
	data2, err := hex.DecodeString(h2)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(data1))
	for i := range data1 {
		out[i] = data1[i] ^ data2[i]
	}
	return hex.EncodeToString(out), nil
}
