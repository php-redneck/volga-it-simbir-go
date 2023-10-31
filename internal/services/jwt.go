package services

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/google/uuid"
	"os"
	"time"
)

type JWTService struct{}

type TokenPayload struct {
	jwt.Payload
	Id       int    `json:"id"`
	IsAdmin  bool   `json:"isAdmin"`
	Username string `json:"username,omitempty"`
}

func (s JWTService) Make(id int, username string, isAdmin bool) string {
	now := time.Now()

	token, _ := jwt.Sign(TokenPayload{
		Payload: jwt.Payload{
			Issuer:         "Simbir.Go",
			Subject:        "auth",
			ExpirationTime: jwt.NumericDate(now.Add(24 * 30 * 12 * time.Hour)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          uuid.New().String(),
		},
		Id:       id,
		Username: username,
		IsAdmin:  isAdmin,
	}, jwt.NewHS256([]byte(os.Getenv("JWT_SECRET"))))

	return string(token)
}

func (s JWTService) IsValid(token string) bool {
	var (
		now             = time.Now()
		pl              TokenPayload
		iatValidator    = jwt.IssuedAtValidator(now)
		expValidator    = jwt.ExpirationTimeValidator(now)
		validatePayload = jwt.ValidatePayload(&pl.Payload, iatValidator, expValidator)
	)

	_, err := jwt.Verify([]byte(token), jwt.NewHS256([]byte(os.Getenv("JWT_SECRET"))), &pl, validatePayload)

	return err == nil
}

func (s JWTService) Decode(token string) TokenPayload {
	var tokenPayload TokenPayload

	_, _ = jwt.Verify([]byte(token), jwt.NewHS256([]byte(os.Getenv("JWT_SECRET"))), &tokenPayload)

	return tokenPayload
}
