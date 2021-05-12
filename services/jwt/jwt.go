package jwt

import (
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"kawaii-blog-engine/config"
	"time"
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

	parser := jwt.Parser{UseJSONNumber: true}
	return parser.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config("SECRET")), nil
	})
}

func Refresh(tokenString string) (*jwt.Token, error) {
	verifiedToken, err := Parse(tokenString)

	switch {
	case isExpired(err):
		claims := verifiedToken.Claims.(jwt.MapClaims)

		if isRefreshable(claims) {
			claims["exp"] = config.ExpirationTime(2).Unix()

			newTokenString, createErr := Create(claims)
			newToken, parseErr := Parse(newTokenString)
			if createErr != nil || parseErr != nil {
				return nil, errors.New("couldn't create or parse the jwt token")
			}

			return newToken, nil
		} else {
			return nil, errors.New("couldn't refresh jwt token")
		}

	case err != nil:
		return nil, err

	default:
		return verifiedToken, nil
	}
}

func isRefreshable(claims jwt.MapClaims) bool {
	refreshUntil, _ := claims["refresh_until"].(json.Number).Int64()
	return refreshUntil > time.Now().Unix()
}

func isExpired(err error) bool {
	ve, ok := err.(*jwt.ValidationError)
	return ok && ve.Errors&(jwt.ValidationErrorExpired) != 0
}

