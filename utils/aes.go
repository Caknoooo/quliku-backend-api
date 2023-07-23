package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

const (
	KEY = "8e71bbce7451ba2835de5aea73e4f3f96821455240823d2fd8174975b8321bfc!"
)

func AESEncrypt(stringToEncrypt string) (encryptedString string, err error) {
	key, _ := hex.DecodeString(KEY)
	plainText := []byte(stringToEncrypt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nonce, nonce, plainText, nil)
	return fmt.Sprintf("%x", cipherText), nil
}

func AESDecrypt(encryptedString string) (decryptedString string, err error) {
	defer func() {
		if r := recover(); r != nil {
			decryptedString = ""
			err = errors.New("error in decrypting")
		}
	}()

	key, _ := hex.DecodeString(KEY)
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", nil
	}

	return fmt.Sprintf("%x", plaintext), nil
}