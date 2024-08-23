package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func checkFileType(path string) (int) {
	fileExt := filepath.Ext(path)
	if fileExt != ".txt" {
		return 0
	} else {
		return 1
	}
} 

func checkValidFilePath(path string) {
	if path == "" {
		log.Fatal("Error: Invalid file path")
	}
	if checkFileType(path) == 0 {
		log.Fatal("Error: Invalid file path")
	}
}

func checkValidKey(key string) {
	if key == "" {
		log.Fatal("Error: Must provide secret key")
	}
	if len(key) != 32 {
		log.Fatal("Error: Invalid key size, must be size 32")
	}
}

func clearFile(path string) {
	err := os.Truncate(path, 0)
	if err != nil {
		log.Fatal(err)
	}
}

func readFile(path string) ([]byte) {
	body, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func writeFile(path string, strToWrite string) {
	err := os.WriteFile(path, []byte(strToWrite), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func EncryptFile(path string, key string) {
	checkValidFilePath(path) 
	checkValidKey(key)

	body := readFile(path)

	// Encryption logic
	aes, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		log.Fatal(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		log.Fatal(err)
	}

	cipherText := gcm.Seal(nonce, nonce, []byte(body), nil)

	clearFile(path)
	writeFile(path, string(cipherText))

	fmt.Println("File has been encrypted with secret key!")
}

func DecryptFile(path string, key string) {
	checkValidFilePath(path) 
	checkValidKey(key)

	body := readFile(path)

	aes, err := aes.NewCipher([]byte(key))
    if err != nil {
        log.Fatal(err)
    }

    gcm, err := cipher.NewGCM(aes)
    if err != nil {
        log.Fatal(err)
    }

	nonceSize := gcm.NonceSize()
    nonce, ciphertext := body[:nonceSize], body[nonceSize:]

    originText, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
    if err != nil {
        log.Fatal(err)
    }

	clearFile(path)
	writeFile(path, string(originText))

	fmt.Println("File has been decrypted with secret key!")
}

