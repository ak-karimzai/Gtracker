package auth_token

type Maker interface {
	CreateToken(userId int, username string) (string, error)
	VerifyToken(token string) (*Payload, error)
}
