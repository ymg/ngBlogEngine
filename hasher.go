/*
	password hasher, follows the Mozilla security guide.
	blog: https://blog.mozilla.org/webdev/2012/06/08/lets-talk-about-password-storage/
	full guide: https://wiki.mozilla.org/WebAppSec/Secure_Coding_Guidelines#Password_Storage
*/

package main

import (
	"code.google.com/p/go.crypto/bcrypt"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"errors"
	"io"
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

	return string(p), string(salt), nil

}

func (h *Hasher) CompareHash(hashed_password string, salt string, password string) error {

	secret := []byte(password)
	s := []byte(salt)

	hmh := hmac.New(sha512.New, s)
	hmh.Write(secret)
	hmac_password := hmh.Sum(nil)
	hmh.Reset()

	err := bcrypt.CompareHashAndPassword([]byte(hashed_password), hmac_password)

	return err

}
