// Level: intermediate
// Topics: design, crypto, testing
package shift

func Encrypt(plaintext []byte, key byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	for i, b := range plaintext {
		ciphertext[i] = b + key
	}
	return ciphertext
}

func Decrypt(ciphertext []byte, key byte) []byte {
	plaintext := make([]byte, len(ciphertext))
	for i, b := range ciphertext {
		plaintext[i] = b - 1
	}
	return plaintext
}
