package auth

type RegisterTokenDto struct {
	LoginId   string `json:"id"`
	Secret    string `json:"secret"`
	IsNewUser bool   `json:"newUser"`
}

type Token struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}
