package res

type Token struct {
	Token string `json:"token"`
}

type User struct {
	Name          string `json:"userid"`
	DisplayName   string `json:"name"`
	Email         string `json:"email"`
	Avatar        string `json:"avatar"`
	EmailVerified bool   `json:"email_verified"`
	Token         string
}
