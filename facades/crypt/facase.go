package crypt

import (
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/crypt"
)

func getCrypt() crypt.ICrypt {
	return container.MustMake[crypt.ICrypt]("crypt")
}

func EncryptString(value string) (string, error) {
	return getCrypt().EncryptString(value)
}

// DecryptString decrypts the given string payload, returning the decrypted string and an error if any.
func DecryptString(payload string) (string, error) {
	return getCrypt().DecryptString(payload)
}
