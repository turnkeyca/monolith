package auth

import (
	"encoding/json"
	"io"
)

type RegisterTokenDto struct {
	LoginId   string `json:"id"`
	Secret    string `json:"secret"`
	IsNewUser bool   `json:"newUser"`
}

type TokenDto struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

func Read(r io.Reader) (*RegisterTokenDto, error) {
	d := RegisterTokenDto{}
	err := json.NewDecoder(r).Decode(&d)
	return &d, err
}
