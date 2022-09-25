package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/IqbalLx/alterra-agmc/entities"
	"github.com/golang-jwt/jwt"
)

type IJWTUtil interface {
	Generate(userId uint) (string, error)
	Validate(authToken string) (entities.Token, error)
}

type gojwtJWTUtil struct {
	secretKey string
}

func NewGoJWTUtil(secretKey string) *gojwtJWTUtil {
	return &gojwtJWTUtil{secretKey}
}
func (gj *gojwtJWTUtil) Generate(userId uint) (string, error) {
	var mySigningKey = []byte(gj.secretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["UserId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix() // valid 7 days

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (gj *gojwtJWTUtil) Validate(authToken string) (entities.Token, error) {
	var mySigningKey = []byte(gj.secretKey)
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return mySigningKey, nil
	})

	if err != nil {
		return entities.Token{}, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return entities.Token{
			UserId: uint(claims["UserId"].(float64)),
		}, nil
	}

	return entities.Token{}, errors.New("invalid token")
}
