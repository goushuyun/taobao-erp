// Copyright 2015-2016, Wothing Co., Ltd.
// All rights reserved.

package token

import (
	"crypto/rsa"
	"io/ioutil"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var defaultPublicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0efAxgbs8Rggp15yt8FS
SnsXWOzKY2KxEqf0u2eVobUajU7+FoS2Xy6sPqVi+WPAwvn4fUhxidq5CuO4+bQj
hIsCJD6VJdUdE9UxJ1chGhJ7OxU7lNIS1P25s5aBaw2O8QZNpY7CVt8SZ1o0KaAi
YEyE6MfVTo/9wjDdXoTFYUc0rNkBFCxgPsWYPII8nSTM43o9432x5ywj9qC1hzdf
xjQztkomK9oiaZiCsGnamA7XT7+uUYEO7YYVLyG8Ap2VRvGoa8MKbLGl9UaG0eBX
pZv5yMDNGxW7+ByrHbQVbu+CIF2ZTvmkOrUULOUO5HsmaZFMHpsln+b88NiwlmJ6
BQIDAQAB
-----END PUBLIC KEY-----`)

var defaultPrivateKey = []byte(`-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDR58DGBuzxGCCn
XnK3wVJKexdY7MpjYrESp/S7Z5WhtRqNTv4WhLZfLqw+pWL5Y8DC+fh9SHGJ2rkK
47j5tCOEiwIkPpUl1R0T1TEnVyEaEns7FTuU0hLU/bmzloFrDY7xBk2ljsJW3xJn
WjQpoCJgTITox9VOj/3CMN1ehMVhRzSs2QEULGA+xZg8gjydJMzjej3jfbHnLCP2
oLWHN1/GNDO2SiYr2iJpmIKwadqYDtdPv65RgQ7thhUvIbwCnZVG8ahrwwpssaX1
RobR4Felm/nIwM0bFbv4HKsdtBVu74IgXZlO+aQ6tRQs5Q7keyZpkUwemyWf5vzw
2LCWYnoFAgMBAAECggEAIOf1/pVjW8Bujg5uaYQzBF4boOMuLzpvi/8sjJyGhp0/
lluF1b3kYTON6RxAUdxjga1yWSGcOwJA9AYTH4Iv9z1bjpcJBq9MKanIVOSB2fZ/
vxlrB7+PGDjWfeLgUwoDGKHmVkf1C21ZEz6+4q+p8/LK+zsoo3JLU8AZVBH4Z3Fb
VKIlhjalWZGpOllvoAS4cnRc/i20RuUwoApgk8Eqnk7HcwgWFrSPRVkPkLMbxdi0
yVGsbaUm2Hd01z7gmbi1cxkF4pMtuuRWsdJwEmItor7OiPH3QNKA8y7/hOHopOHM
NUuNbo/dSr78y4jnHryz3i5zkoxXh9YFyOShG0/MgQKBgQD7IZy7QUGKaulHoAfC
r5OGHJLM1l4VzOMit8IESemm/THhZVDm7q3PcoDRIX7m1fZYBJhHO7iYX+uo1PHu
YsAPXGhj0YXhOJjzUIr6IEJlnmkpFikI3he+h4l5tDsXp1MqBPVfY4too55m/Nvg
uaOhDwxshzOfJBPj+DP/I6pH9QKBgQDV+YhEOGBOXzj8vU2GswDYKjRAGaoWrkHX
NFRvmqZrAa+aLrLz0dA6/0HUspUgFk3asIGCgtTcE1oBJAYLtBDGDtqSzYq+55lH
/ltSXfBFhxzbK0eT7c+mSXTGxOIna/Q5oYSQYTUGwBhuod2PDf36ZJPhNgekR8I2
jNdGdHDv0QKBgQCrifQjPJnmUMz0Le3fIEtmylHENZGi1oc4CckvYMWHWWAfFDPE
6rgzAYXYVEb4qqJQ0SKrVbHr82lns71mFnIWjAqKVG5cv2pKmXO1EyAHhcNTW13A
PuR9MtvHFENhDtyR0T1CYR6y5UCoHISc5nFM7JiR8XBjfiNQDxSFbPk9mQKBgQCC
c24bQU+dmDozA+pG68zg/OH8Dml/CGAFptavb9ZzuIRpeH0LXugXf9WPqgx2koKG
iEN84OyK+5VMyryQ2Ae96AJjq8Ih3yq5FJ5yWekJSnVSPVGXI04McA4svI2wQQWV
OR6Ls2fTpLuAf1iHRZ2I0VbC3+DKzCDghv8b6hOMcQKBgD0xNBDpX1IZjTpbFEU6
jK7qPJgTifsFTCH6fs/Bx6M1pIScdWu0K2wGK+fTUcSbaVPyJuiUvW/5S6xfsJoq
xFUye6ILP6XrwYd0wrZi5r5fD03xDcmz48bjVCLUjAwumtZNqhshlUFSlpdYP+Ju
j04rmwRunyzAOg0T2neEiY0h
-----END PRIVATE KEY-----`)

var rsaPublicKey *rsa.PublicKey
var rsaPrivateKey *rsa.PrivateKey

var publicKey = readWithDefault("/17mei.crt", defaultPublicKey)
var privateKey = readWithDefault("/17mei.key", defaultPrivateKey)

func readWithDefault(f string, defaultKey []byte) []byte {
	key, err := ioutil.ReadFile(f)
	if err != nil {
		if os.IsNotExist(err) {
			return defaultKey
		}
		panic(err)
	}
	return key
}

func GetPublicKeyFromPEM() *rsa.PublicKey {
	if rsaPublicKey == nil {
		var err error
		rsaPublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKey)
		if err != nil {
			panic(err)
		}
	}
	return rsaPublicKey
}

func getPrivateKeyFromPEM() *rsa.PrivateKey {
	if rsaPrivateKey == nil {
		var err error
		rsaPrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKey)
		if err != nil {
			panic(err)
		}
	}
	return rsaPrivateKey
}
