package jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/patrickmn/go-cache"
	jwtCache "kawaii-blog-engine/cache/jwt"
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
	switch {
	case tokenString == "":
		return nil, errors.New("missing token cookie")
	case isRevoked(tokenString):
		fmt.Println("revoked token is...")
		fmt.Println(tokenString)
		fmt.Println(jwtCache.DenyList.Items())
		return nil, errors.New("token has been revoked")
	default:
		parser := jwt.Parser{UseJSONNumber: true}
		return parser.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Config("SECRET")), nil
		})
	}
}

func ParseOrRefresh(tokenString string) (*jwt.Token, error) {
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

			jwtCache.DenyList.Set(verifiedToken.Raw, true, cache.DefaultExpiration)

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

func isRevoked(tokenString string) bool {
	_, found := jwtCache.DenyList.Get(tokenString)
	return found
}
