package common

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserPayload struct {
	ID    int
	Email string
}

var (
	jwtKey = os.Getenv("JWT_KEY")
	kid    = os.Getenv("JWK_KID")
)

type claims struct {
	*UserPayload
	jwt.StandardClaims
}

func CreateToken(userPayload *UserPayload) (string, error) {
	return CreateTokenWithExpire(userPayload, time.Now().Add(336*time.Hour).Unix())
}

func CreateTokenWithExpire(userPayload *UserPayload, exp int64) (string, error) {
	sub := strconv.Itoa(userPayload.ID)
	claims := &claims{
		UserPayload: userPayload,
		StandardClaims: jwt.StandardClaims{
			Subject:   sub,
			ExpiresAt: exp,
			IssuedAt:  time.Now().Unix(),
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	at.Header["typ"] = "JWT"
	at.Header["kid"] = kid
	token, err := at.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ParseToken(token string) (*jwt.Token, *UserPayload, error) {
	claims := &claims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	return parsedToken, claims.UserPayload, err
}
