package helper

import (
	"ExploriteGateway/pkg/utils/models"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type authCustomClaimsUser struct {
	Firstname string `json:"firstname"`
	Email     string `json:"email"`
	ID        int    `json:"id"`
	jwt.StandardClaims
}

func GenerateTokenUser(user models.UserDetailResponse) (string, error) {
	log.Println("Generating token for user:", user.Email)
	claims := &authCustomClaimsUser{
		ID:        user.ID,
		Firstname: user.Firstname,
		Email:     user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("123456789"))
	if err != nil {
		log.Println("Error signing token:", err)
		return "", err
	}

	log.Println("Token generated successfully for user:", user.Email)
	return tokenString, nil
}

func ValidateTokenUser(tokenString string) (*authCustomClaimsUser, error) {
	log.Println("Validating token...")
	token, err := jwt.ParseWithClaims(tokenString, &authCustomClaimsUser{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("Unexpected signing method:", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("123456789"), nil
	})

	if err != nil {
		log.Println("Error parsing token:", err)
		return nil, err
	}

	if claims, ok := token.Claims.(*authCustomClaimsUser); ok && token.Valid {
		log.Println("Token validated successfully.")
		return claims, nil
	}

	log.Println("Invalid token.")
	return nil, fmt.Errorf("invalid token")
}
