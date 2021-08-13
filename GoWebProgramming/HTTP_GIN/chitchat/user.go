package main

import (
	"crypto/sha1"
	"fmt"
)

type User struct {
	Email    string
	Password string
}

func encrypt(pwd string) string {
	enc_str := fmt.Sprintf("%x", sha1.Sum([]byte(pwd)))
	return enc_str
}

func userbyEmail(email string) {
}
