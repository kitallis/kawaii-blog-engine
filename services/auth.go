package services

import (
	"kawaii-blog-engine/config"

	"github.com/dgrijalva/jwt-go"
)

func CreateSignedToken(claims map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenClaims := token.Claims.(jwt.MapClaims)
	for key, value := range claims {
		tokenClaims[key] = value
	}
	signedToken, err := token.SignedString([]byte(config.Config("SECRET")))

	return signedToken, err
}