package main

import (
	"fmt"

	"github.com/miki799/vernam-cipher/utils"
	"github.com/miki799/vernam-cipher/vernam"
)

func main() {
	message, err := utils.ReadTextFromFile("message.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("message to encrypt: %v\n", message)

	bits := utils.GetMessageBitsLength(message)

	fmt.Printf("bit len of message: %v\n", bits)

	fmt.Println("Key generation...")
	key, err := vernam.CreateKey(bits)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Encryption...")
	cryptogram, err := vernam.Encrypt(message, key)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Decryption...")
	encryptedMessage, err := vernam.Decrypt(cryptogram, key)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Verification...")
	result := vernam.Verify(message, encryptedMessage)
	if result {
		fmt.Println("Vernam alghoritm works! Messages are the same!")
		fmt.Printf("encrypted message: %v\n", encryptedMessage)
	} else {
		fmt.Println("Vernam alghoritm doesn't works! Messages are not the same!")
	}
}
