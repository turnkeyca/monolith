package auth

type RegisterTokenDto struct {
	TokenString string `json:"token"`
	Secret      string `json:"secret"`
	IsNewUser   bool   `json:"newUser"`
}

type UserId struct {
	Id string `json:"id"`
}
