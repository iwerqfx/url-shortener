package util

import gonanoid "github.com/matoous/go-nanoid/v2"

const (
	aliasLength  = 6
	aliasCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func GenerateAlias() (string, error) {
	return gonanoid.Generate(aliasCharset, aliasLength)
}
