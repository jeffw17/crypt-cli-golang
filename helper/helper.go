package helper

import (
	"encoding/base64"
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

func checkValidFilePath(path string) (bool) {
	if path == "" {
		return false
	}
	if checkFileType(path) == 0 {
		return false
	}
	return true
}

func EncryptFile(path string) {
	if !checkValidFilePath(path) {
		log.Fatal("Invalid file path")
	}

	body, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	encoding := base64.StdEncoding.EncodeToString(body)

	err = os.Truncate(path, 0)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(path, []byte(encoding), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func DecryptFile(path string) {
	if !checkValidFilePath(path) {
		log.Fatal("Invalid file path")
	}

	body, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	decoding, err := base64.StdEncoding.DecodeString(string(body))
	if err != nil {
		log.Fatal(err)
	}

	err = os.Truncate(path, 0)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(path, []byte(decoding), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

