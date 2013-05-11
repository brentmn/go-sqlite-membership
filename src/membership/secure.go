package main

import (
	"github.com/jameskeane/bcrypt"
)

func Encrypt(p string) string {
	// generate a random salt with 10 rounds of complexity
	//salt, _ := bcrypt.Salt()
	hash, _ := bcrypt.Hash(p)
	return hash
}

func PassMatch(p string, hash string) bool {
	if bcrypt.Match(p, hash) {
		return true
	}
	return false
}
