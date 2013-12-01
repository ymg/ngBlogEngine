package main

import "testing"

//testing hashing and validation functions
func TestHasher(t *testing.T) {

	secret := "hard password!"
	hashing := &Hasher{}
	p, s, err := hashing.NewHash(secret)

	if err != nil {
		t.Error("failed generating hash")
		t.Fail()
	}

	compare_err := hashing.CompareHash(p, s, secret)
	if compare_err != nil {
		t.Error("failed comparing hashes")
		t.Fail()
	}

	compare_err2 := hashing.CompareHash(p, s, "should fail!")
	if compare_err2 == nil {
		t.Error("authenticated wrong passwords")
		t.Fail()
	}
}
