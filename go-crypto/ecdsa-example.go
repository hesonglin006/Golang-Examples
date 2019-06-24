package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	if privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader); err != nil {
		panic(err)
	}else {
		fmt.Println("Generated private key:", privateKey)

		// The message to be signed
		message := "hello"
		hash := sha256.Sum256([]byte(message))
		//Sign the message
		r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
		if err != nil{
			panic(err)
		}
		fmt.Printf("Signature: (0x%x, 0x%x)\n", r, s)
		//Verify the message
		valid := ecdsa.Verify(&privateKey.PublicKey, hash[:], r, s)
		fmt.Println("Signature verification result: ", valid)
		// The failed case of verification
		validFalse := ecdsa.Verify(&privateKey.PublicKey, hash[1:], r, s)
		fmt.Println("Signature verification result: ", validFalse)
	}
}
