package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"
)

func main() {
	data := "private reord"
	printText := []byte(data)
	log.Println("print text", printText)

	key := make([]byte, 32)
	if _, err := rand.Reader.Read(key); err != nil {
		log.Println("err in generate encryption key", err.Error())
		return
	}
	//aes block cipher Creation
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err.Error())
		return
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println("error in create gcm mode", err.Error())
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Println("error ingenerating nonce", err.Error())
		return
	}
	ciphertext := gcm.Seal(nonce, nonce, printText, nil)
	enc := hex.EncodeToString(ciphertext)
	log.Println(enc)
}
func DecryptData() {

}
