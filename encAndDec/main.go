package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
)

// Initialize victor, which is the random bytes
var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// Keep this secret key with you.
const secretKey string = "abc&1*~#^2^#s0^=)^^7%b34"

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(text, secretKey string) (string, error) {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return encodeBase64(cipherText), nil
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text, secretKey string) (string, error) {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}
	cipherText := decodeBase64(text)
	cfb := cipher.NewCFBDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

func main() {

	// phrase := "Hello World!"

	person := Person{
		Name:   "Ansun",
		Age:    25,
		Active: true,
	}

	// Marshalling Go struct to JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	// To encrypt the phrase
	encText, err := Encrypt(string(jsonData), secretKey)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("encrypted text: ", encText)

	// To decrypt the original phrase
	decText, err := Decrypt("U2FsdGVkX1/4Gxkihez7P83XelbZBFwcCofKGHV8EnUJ56187l9zumWEJixzGrBBd1Rv2Hwf9r+M1FGjeL0Fk7eRPtjfBmeWioYjdLVgIsl9w5rLl+4v72/38zLkT01lqpcUCWvk0blEdBwKKFl/vmQjWWPSfSO4E6xqGjnV6p8hxwYHZXtp/GIMhAm/96DThIHTwAh5g1GDuTbru3cuaaKSGA33o0obPbPM1vA15RRg149QRX/R0mUAtmmBaL/b3MhfQl0NXLHXsjSKyXLURL4dRgj5he4FBabpOwmR3S0=", "DC8B71BC3AB56DF8C90C8678B08A2BE8")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("decrypted text: ", decText)
}
