// Cryptopals shows basic crypto operations.
//
// Always operate on raw bytes, never on encoded strings. Only use hex and
// base64 for pretty-printing.
//
// Adapted from https://github.com/0xfe/cryptopals.
//
// Level: intermediate
// Topics: crypto, encoding, xor
package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
)

// HexToBase64 decodes hex string h into bytes and then bytes into base64
// string. (Set 1, challenge 1.)
func HexToBase64(h string) (string, error) {
	b, err := hex.DecodeString(h)
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(b), nil
}

// FixedXOR takes the bytes represented by hex1 and hex2 and then does XOR
// bitwise operation on these two byte slices, byte by byte. (Set 1, challenge
// 2.)
func FixedXOR(hex1, hex2 string) (string, error) {
	h1, err := hex.DecodeString(hex1)
	if err != nil {
		return "", err
	}
	h2, err := hex.DecodeString(hex2)
	if err != nil {
		return "", err
	}

	var out = make([]byte, len(h1))
	for i := range h1 {
		out[i] = h1[i] ^ h2[i]
	}
	return hex.EncodeToString(out), nil
}
