package auth

import (
	. "server/config"
	"server/store"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
			Name:   user["Name"].(string),
			APIKey: user["APIKey"].(string),
		}, nil
	}
	return store.CleanUser{}, err
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the token from the header
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" || len(tokenString) < 8 {
			c.Next()
			return
		}
		// parse the token
		user, err := Validate(tokenString[7:])
		if err != nil {
			c.Next()
			return
		}
		// set the user to the context
		c.Set("user", user)
		c.Next()
	}
}
