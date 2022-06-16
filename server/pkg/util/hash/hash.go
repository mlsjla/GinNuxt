package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	MinCost     int = 4  // the minimum allowable cost as passed in to GenerateFromPassword
	MaxCost     int = 31 // the maximum allowable cost as passed in to GenerateFromPassword
	DefaultCost int = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
)

func BcryptEncode(password string) (str string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	strPassword := string(hash)
	return strPassword, err
}

func BcryptCompare(encodePW string, passwordOK string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(passwordOK))
}

// MD5 MD5哈希值
func MD5(b []byte) string {
	h := md5.New()
	_, _ = h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// MD5String MD5哈希值
func MD5String(s string) string {
	return MD5([]byte(s))
}

// SHA1 SHA1哈希值
func SHA1(b []byte) string {
	h := sha1.New()
	_, _ = h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA1String SHA1哈希值
func SHA1String(s string) string {
	return SHA1([]byte(s))
}
