package base

import (
	"crypto/md5"
	"oneday-infrastructure/login/domain"
)

/**
to be changeable
*/
const commonSalt = "one-day"

func init() {
	md5Way = ChooseEncrypt(MD5)
	domain.EncryptMap[MD5] = md5Way
}

var md5Way domain.EncryptHelper

const MD5 domain.EncryptWay = "MD5"

func ChooseEncrypt(way domain.EncryptWay) domain.EncryptHelper {
	switch way {
	case MD5:
		return MD5Way{}
	default:
		panic("can not find helper")

	}
}

type MD5Way struct{}

func (md5 MD5Way) Match(source, encryptedString string) bool {
	return md5.Encrypt(source) == encryptedString
}

func (MD5Way) Encrypt(password string) string {
	data := []byte(password)
	md5Bytes := md5.Sum(data)
	return string(md5Bytes[:])
}

func (md5 MD5Way) Decrypt(password string) string {
	panic("not support")
}
