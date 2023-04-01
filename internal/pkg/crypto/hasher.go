package crypto

import (
	"log"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

type Hasher interface {
	HashAndSalt(value []byte) string
	CompareHashes(hashed string, plain []byte) bool
}

type BCryptHasher struct{}

var bCryptHasher *BCryptHasher
var generateFromPassword func(password []byte, cost int) ([]byte, error)
var hasherOnce sync.Once

func GetHasher() Hasher {
	hasherOnce.Do(func() {
		bCryptHasher = &BCryptHasher{}
		generateFromPassword = bcrypt.GenerateFromPassword
	})
	return bCryptHasher
}

func (h *BCryptHasher) HashAndSalt(value []byte) string {
	hash, err := generateFromPassword(value, bcrypt.MinCost)
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
