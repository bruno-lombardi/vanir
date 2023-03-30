package crypto_tests

import (
	"testing"
	"vanir/internal/pkg/crypto"

	"github.com/stretchr/testify/suite"
)

type BCryptHasherSuite struct {
	suite.Suite
	hasher *crypto.BCryptHasher
}

func (sut *BCryptHasherSuite) BeforeTest(_, _ string) {
	sut.hasher = &crypto.BCryptHasher{}
}

func (sut *BCryptHasherSuite) TestSmokeGetHasher() {
	hasher, ok := crypto.GetHasher().(*crypto.BCryptHasher)
	sut.NotNil(hasher)
	sut.True(ok)
}

func (sut *BCryptHasherSuite) TestShouldReturnAHashedString() {
	plain := "a_string_to_hash"
	hashed := sut.hasher.HashAndSalt([]byte(plain))
	sut.NotEmpty(hashed)
	sut.NotEqual(plain, hashed)
}

func (sut *BCryptHasherSuite) TestShouldReturnTrueWhenComparingMatchingHashes() {
	plain := "a_string_to_hash"
	hashed := sut.hasher.HashAndSalt([]byte(plain))
	isCompareSuccessful := sut.hasher.CompareHashes(hashed, []byte(plain))
	sut.True(isCompareSuccessful)
}

func (sut *BCryptHasherSuite) TestShouldReturnFalseWhenComparingUnmatchingHashes() {
	plain := "a_string_to_hash"
	hashed := sut.hasher.HashAndSalt([]byte(plain))
	isCompareSuccessful := sut.hasher.CompareHashes(hashed, []byte("a_different_string_to_hash"))
	sut.False(isCompareSuccessful)
}

func TestBCryptHasherSuite(t *testing.T) {
	suite.Run(t, &BCryptHasherSuite{})
}
