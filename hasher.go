/*
	password hasher, follows the Mozilla security guide.
	blog: https://blog.mozilla.org/webdev/2012/06/08/lets-talk-about-password-storage/
	full guide: https://wiki.mozilla.org/WebAppSec/Secure_Coding_Guidelines#Password_Storage
*/

package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"io"

	"golang.org/x/crypto/bcrypt"
)

const (
	KEYLENGTH = 32
)

var ErrGeneratingSalt = errors.New("error generating salt")
var ErrGeneratingBcrypt = errors.New("error generating bcrypt hash")

type Hasher struct{}

func (h *Hasher) NewHash(password string) (string, string, error) {

	b := make([]byte, KEYLENGTH)
	_, err := io.ReadFull(rand.Reader, b)

	if err != nil {
		return "", "", ErrGeneratingSalt
	}

	secret := []byte(password)
	salt := b

	hmh := hmac.New(sha512.New, salt)
	hmh.Write(secret)
	hmac_hash := hmh.Sum(nil)
	hmh.Reset()

	p, err := bcrypt.GenerateFromPassword(hmac_hash, 16)

	if err != nil {
		return "", "", ErrGeneratingBcrypt
	}

	return string(p), hex.EncodeToString(salt), nil

}

func (h *Hasher) CompareHash(hashed_password string, salt string, password string) error {

	secret := []byte(password)
	s, _ := hex.DecodeString(salt)

	hmh := hmac.New(sha512.New, s)
	hmh.Write(secret)
	hmac_password := hmh.Sum(nil)
	hmh.Reset()

	return bcrypt.CompareHashAndPassword([]byte(hashed_password), hmac_password)

}
