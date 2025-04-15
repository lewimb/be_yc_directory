package pkg

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Email string `json:"email"`
	ID    int    `json:"id"`
	jwt.RegisteredClaims
}

var secretKey = []byte("secret-key")

func CreateToken(email string, id int) (string, error) {
	claims := CustomClaims{
		Email: email,
		ID:    id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	fmt.Println(tokenString)
	return tokenString, nil
}

func GetHeader(header string) string {
	headerArr := strings.Split(header, " ")
	token := headerArr[1]

	log.Println(headerArr)
	log.Println(token)

	return token
}

func UnloadToken(token string) int {
	var id int

	log.Println(token)

	tokenString, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})

	if err != nil {
		panic(err)
	}

	if claims, ok := tokenString.Claims.(*CustomClaims); ok && tokenString.Valid {
		id = claims.ID
	} else {
		panic("Unable to get custom claims")
	}
	return id
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
