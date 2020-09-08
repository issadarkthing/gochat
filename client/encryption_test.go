package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestEncrypt(t *testing.T) {

	samples := []string{
		"jiman",
		"ahmad",
		"terra sasmple",
	}

	key := []byte("test")

	for _, in := range samples {

		out := encrypt([]byte(key), key)
		assert.NotEqual(t, out, string(in))
	}
}

func TestDecrypt(t *testing.T) {

	key := []byte("key")

	samples := []string{
		"eagle",
		"golang",
		"plan 9",
	}

	results := make(map[string]string, len(samples))

	for _, sample := range samples {
		results[sample] = string(encrypt([]byte(sample), key))
	}

	for expected, v := range results {

		decrypted, err := decrypt([]byte(v), key)
		if err != nil {
			panic(err)
		}

		got := string(decrypted)

		assert.Equal(t, expected, got)
	}
}
