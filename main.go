// I found this sample of ECC using elktree's "GenerateKeys", "Encrypt" and "Decrypt" functions, I feel as if this page had the most simple example code-wise for me to understand ECC, I further edited the code to allow for user input
// https://pkg.go.dev/gitlab.com/elktree/ecc
// Only problem is it only encrypts first word
package main

import (
	"crypto/elliptic"
	"fmt"
	"gitlab.com/elktree/ecc"
)

func main() {
	// Create keys using his GenerateKeys function within ECC call and the highest level P521
	pub, priv, _ := ecc.GenerateKeys(elliptic.P521())

	// Asks for user input for plaintext
	fmt.Println("Enter Your Plaintext: ")
	var plaintext string
	fmt.Scanln(&plaintext)

	// Uses a public key to complete encrypt function
	encrypted, _ := pub.Encrypt([]byte(plaintext))
	fmt.Println(len(encrypted))

	// Uses private key to complete decrypt function
	decrypted, _ := priv.Decrypt(encrypted)
	fmt.Println(string(decrypted))

}
