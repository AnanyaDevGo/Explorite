package helper

import (
	"ExploriteGateway/pkg/utils/models"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type authCustomClaimsAdmin struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	jwt.StandardClaims
}

func PasswordHash(password string) (string, error) {
	log.Println("Hashing password...")
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Println("Error hashing password:", err)
		return "", errors.New("internal server error")
	}
	hash := string(hashPassword)
	log.Println("Password hashed successfully.")
	return hash, nil
}

func GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error) {
	log.Println("Generating token for admin:", admin.Email)
	claims := &authCustomClaimsAdmin{
		Firstname: admin.Firstname,
		Lastname:  admin.Lastname,
		Email:     admin.Email,
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
	log.Println("Token generated successfully for admin:", admin.Email)
	return tokenString, nil
}

func ValidateToken(tokenString string) (*authCustomClaimsAdmin, error) {
	log.Println("Validating token...")
	token, err := jwt.ParseWithClaims(tokenString, &authCustomClaimsAdmin{}, func(token *jwt.Token) (interface{}, error) {
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

	if claims, ok := token.Claims.(*authCustomClaimsAdmin); ok && token.Valid {
		log.Println("Token validated successfully.")
		return claims, nil
	}
	log.Println("Invalid token.")
	return nil, fmt.Errorf("invalid token")
}
