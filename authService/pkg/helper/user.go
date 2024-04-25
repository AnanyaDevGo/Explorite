package helper

import (
	"authservice/pkg/utils/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type authCustomClaimsUser struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	jwt.StandardClaims
}

func GenerateTokenUser(user models.UserDetailResponse) (string, error) {
	claims := &authCustomClaimsUser{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("123456789"))
	if err != nil {
		fmt.Println("Error is", err)
		return "", err
	}

	return tokenString, nil
}

func ValidateTokenUser(tokenString string) (*authCustomClaimsUser, error) {
	token, err := jwt.ParseWithClaims(tokenString, &authCustomClaimsUser{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("123456789"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*authCustomClaimsUser); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
