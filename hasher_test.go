package main

import (
	"testing"
)

//testing hashing and validation functions
func Test_Hasher(t *testing.T) {

	secret := "hard password!"

	hashing := &Hasher{}
	p, s, err := hashing.NewHash(secret)

	if err != nil {
		t.Errorf("failed generating hash")
	}

	compare_err := hashing.CompareHash(p, s, secret)

	if compare_err != nil {
		t.Errorf("failed comparing hashes")
	}
}
