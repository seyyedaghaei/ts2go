package main

import (
	"crypto/rand"
	"log"
	"os"
	"path/filepath"
)

var chars = "abcdefghijklmnopqrstuvwxyz1234567890"

func randomText(length int) string {
	ll := len(chars)
	b := make([]byte, length)
	_, _ = rand.Read(b) // generates len(b) random bytes
	for i := 0; i < length; i++ {
		b[i] = chars[int(b[i])%ll]
	}
	return string(b)
}

func newCacheDir() string {
	dir, err := os.UserHomeDir()
	handleError(err)
	path := filepath.Join(dir, ".ts2go", randomText(16))
	err = os.MkdirAll(path, os.ModePerm)
	handleError(err)
	return path
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}