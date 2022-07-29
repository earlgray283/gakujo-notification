package lib

import (
	"testing"
)

func TestValidateString(t *testing.T) {
	assert(t, ValidateString("aiueo", MinLen(1), MaxLen(5)))
	// assert(t, !ValidateString("aiu", MinLen(4), MaxLen(5)))
	// assert(t, ValidateString("aいueo", MinLen(1), MaxLen(5), AllowMultibyte()))
	// assert(t, !ValidateString("aいueo", MinLen(1), MaxLen(5)))
	//assert(t, ValidateString("aiueo", MinLen(1), MaxLen(5), WhiteList([]rune{'a', 'i', 'u', 'e', 'o'})))
	//assert(t, ValidateString("aiueo", MinLen(1), MaxLen(5), WhiteList([]rune{'a', 'i', 'u', 'e', 'o', 'k'})))

}

func TestCrypto(t *testing.T) {
	key := "amazingmightyyyy"
	crypto, err := NewCrypto([]byte(key))
	if err != nil {
		t.Fatal(err)
	}
	plainText := "aiueo"
	encrypted, err := crypto.Encrypt([]byte(plainText))
	if err != nil {
		t.Fatal(err)
	}
	decrypted, err := crypto.Decrypt(encrypted)
	if err != nil {
		t.Fatal()
	}
	assert(t, plainText == decrypted)
}

func assert(t *testing.T, b bool) {
	if !b {
		t.Fatal("assertion failed")
	}
}
