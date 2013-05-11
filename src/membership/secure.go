package main

import (
	"github.com/jameskeane/bcrypt"
)

func Encrypt(p string) string {
	hash, _ := bcrypt.Hash(p)
	return hash
}

func PassMatch(p string, hash string) bool {
	if bcrypt.Match(p, hash) {
		return true
	}
	return false
}
