package utils

import (
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const chars = "_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = chars[seededRand.Intn(len(chars))]
	}
	return string(b)
}

func TrimFileExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

// Get a value from the environment or use default
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
