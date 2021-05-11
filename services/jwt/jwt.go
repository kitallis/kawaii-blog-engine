package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"kawaii-blog-engine/config"
)

func Create(claims map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenClaims := token.Claims.(jwt.MapClaims)

	for key, value := range claims {
		tokenClaims[key] = value
	}

	signedToken, err := token.SignedString([]byte(config.Config("SECRET")))

	return signedToken, err
}

func Parse(tokenString string) (*jwt.Token, error) {
	if tokenString == "" {
		return nil, errors.New("missing token cookie")
	}

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config("SECRET")), nil
	})
}
