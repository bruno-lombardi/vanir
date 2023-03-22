package crypto

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Hasher interface {
	HashAndSalt(value []byte) string
	CompareHashes(hashed string, plain []byte) bool
}

type BCryptHasher struct{}

var bCryptHasher *BCryptHasher

func GetHasher() Hasher {
	if bCryptHasher == nil {
		bCryptHasher = &BCryptHasher{}
	}
	return bCryptHasher
}

func (h *BCryptHasher) HashAndSalt(value []byte) string {
	hash, err := bcrypt.GenerateFromPassword(value, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func (h *BCryptHasher) CompareHashes(hashed string, plain []byte) bool {
	byteHash := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(byteHash, plain)
	return err == nil
}
