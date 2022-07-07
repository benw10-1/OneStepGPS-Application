package auth

import (
	. "server/config"
	"github.com/golang-jwt/jwt/v4"
	"server/store"
)

var signingMethod = jwt.SigningMethodHS256
var signingKey = []byte(GetEnv("JWT_SECRET", "secret"))
// create a new token for the given user
func Sign(user store.CleanUser) (string, error) {
	token := jwt.New(signingMethod)

	token.Claims.(jwt.MapClaims)["user"] = user

	return token.SignedString(signingKey)
}
// parse the token and return the user
func Validate(tokenString string) (store.CleanUser, error) {
	// parsing token. Pass function that returns the original signing key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return store.CleanUser{}, err
	}
	// if token is valid, return the user
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// cast the claims to a map
		user := claims["user"].(map[string]interface{})
		if err != nil {
			return store.CleanUser{}, err
		}
		// return the user
		return store.CleanUser{
			Name: user["Name"].(string),
			APIKey: user["APIKey"].(string),
		}, nil
	}
	return store.CleanUser{}, err
}