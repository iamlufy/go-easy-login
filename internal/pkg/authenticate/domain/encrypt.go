package domain

import (
	"crypto/md5"
	"fmt"
)

type Matcher func(source, encryptedString string) bool
type Encrypter func(string) string

const MD5 string = "MD5"

func ChooseEncrypter(way string) Encrypter {
	switch way {
	case MD5:
		return func(source string) string {
			return md5Encrypt(source)
		}
	default:
		panic("can not find helper")
	}
}

func ChooseMatcher(way string) Matcher {
	switch way {
	case MD5:
		return func(source, encryptedString string) bool {
			return encryptedString == md5Encrypt(source)
		}
	default:
		return func(source, encryptedString string) bool {
			return source == encryptedString
		}
	}
}

func md5Encrypt(password string) string {
	data := []byte(password)
	md5Bytes := md5.Sum(data)
	return fmt.Sprintf("%x", md5Bytes)
}
