package auth

import (
	. "server/config"
	"github.com/golang-jwt/jwt/v4"
	"server/store"
)

var signingMethod = jwt.SigningMethodHS256
var signingKey = []byte(GetEnv("JWT_SECRET", "secret"))

func Sign(user store.CleanUser) (string, error) {
	token := jwt.New(signingMethod)

	token.Claims.(jwt.MapClaims)["user"] = user

	return token.SignedString(signingKey)
}

func Validate(tokenString string) (store.CleanUser, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return store.CleanUser{}, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := claims["user"].(map[string]interface{})
		if err != nil {
			return store.CleanUser{}, err
		}

		return store.CleanUser{
			Name: user["Name"].(string),
			APIKey: user["APIKey"].(string),
		}, nil
	}
	return store.CleanUser{}, err
}