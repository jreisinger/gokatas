package cookies

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
)

// WriteEncrypted encrypts (confidentiality) and authenticates (integrity) cookies.
func WriteEncrypted(w http.ResponseWriter, cookie http.Cookie, secretKey []byte) error {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return err
	}

	plaintext := fmt.Sprintf("%s:%s", cookie.Name, cookie.Value)
	ecnryptedValue := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)
	cookie.Value = string(ecnryptedValue)
	return Write(w, cookie)
}
