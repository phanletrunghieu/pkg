package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var ErrInvalidToken = errors.New("Invalid token")

type Claims struct {
	jwt.StandardClaims
	User interface{}
}

func NewSignedToken(userClaims interface{}, key []byte) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		User: userClaims,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	signedToken, _ := token.SignedString(key)

	return signedToken
}

func VerifyToken(signedToken string, key []byte) (*Claims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}
