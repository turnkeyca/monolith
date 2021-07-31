package auth

type AuthDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

type UserId struct {
	Id string `json:"id"`
}
