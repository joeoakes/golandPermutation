package main

import (
	"fmt"
	"strings"
)

func main() {
	// Define a permutation key as a mapping of characters
	// I used an Enigma rotor
	key := map[rune]rune{
		'a': 'e',
		'b': 'k',
		'c': 'm',
		'd': 'f',
		'e': 'l',
		'f': 'g',
		'g': 'd',
		'h': 'q',
		'i': 'v',
		'j': 'z',
		'k': 'n',
		'l': 't',
		'm': 'o',
		'n': 'w',
		'o': 'y',
		'p': 'h',
		'q': 'x',
		'r': 'u',
		's': 's',
		't': 'p',
		'u': 'a',
		'v': 'i',
		'w': 'b',
		'x': 'r',
		'y': 'c',
		'z': 'j',
	}

	plaintext := "hello world"
	ciphertext := encrypt(plaintext, key)
	decryptedText := decrypt(ciphertext, key)

	fmt.Println("Plaintext: ", plaintext)
	fmt.Println("Ciphertext: ", ciphertext)
	fmt.Println("Decrypted text: ", decryptedText)
}

// Function to encrypt the text using the permutation key
func encrypt(text string, key map[rune]rune) string {
	var encrypted strings.Builder

	for _, char := range text {
		// If the character exists in the key, replace it; otherwise, keep it as is
		encryptedChar, exists := key[char]
		if exists {
			encrypted.WriteRune(encryptedChar)
		} else {
			encrypted.WriteRune(char)
		}
	}

	return encrypted.String()
}

// Function to decrypt the text using the reverse permutation key
func decrypt(text string, key map[rune]rune) string {
	// Create a reverse key by swapping the keys and values of the original key
	reverseKey := make(map[rune]rune)
	for k, v := range key {
		reverseKey[v] = k
	}

	var decrypted strings.Builder

	for _, char := range text {
		// If the character exists in the reverse key, replace it; otherwise, keep it as is
		decryptedChar, exists := reverseKey[char]
		if exists {
			decrypted.WriteRune(decryptedChar)
		} else {
			decrypted.WriteRune(char)
		}
	}

	return decrypted.String()
}
