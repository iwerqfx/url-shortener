package model

import "errors"

var (
	ErrURLNotFound = errors.New("url with this alias was not found")
)

type URL struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Alias string `json:"alias"`
	Views int    `json:"views"`
}
