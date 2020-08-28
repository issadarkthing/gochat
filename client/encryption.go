// Copyright (C) 2020 Raziman Mahathir

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"

	"golang.org/x/crypto/nacl/secretbox"
)


func hashKey(key string) []byte {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hasher.Sum(nil)
}

func encrypt(message, key []byte) []byte {

	var secretKey [32]byte
	copy(secretKey[:], key)

	var nonce [24]byte
	if _, err := io.ReadAtLeast(rand.Reader, nonce[:], 24); err != nil {
		panic(err)
	}

	encrypted := secretbox.Seal(nonce[:], message, &nonce, &secretKey)
	encoded := make([]byte, hex.EncodedLen(len(encrypted)))
	hex.Encode(encoded, encrypted)
	return encoded
}

func decrypt(message, key []byte) ([]byte, error) {

	var secretKey [32]byte
	copy(secretKey[:], key)

	var decryptNonce [24]byte

	decoded := make([]byte, hex.DecodedLen(len(message)))
	_, err := hex.Decode(decoded, message)
	if err != nil {
		return []byte{}, err
	}

	copy(decryptNonce[:], decoded[:24])
	decrypted, ok := secretbox.Open(nil, decoded[24:], &decryptNonce, &secretKey)
	if !ok {
		return []byte{}, errors.New("Unable to decrypt")
	}

	return decrypted, nil
}
