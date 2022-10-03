package services

import (
	"fmt"
	// "go/token"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	GenerateToken(UserID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jWTService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jWTService{
		issuer:    "jjj",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_TOKEN")
	if secretKey != "" {
		secretKey = "jjj"
	}
	return secretKey
}

func (j *jWTService) GenerateToken(UserID string) string {
	claims := &jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)

	}
	return t
}

func (j *jWTService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("upexcapted %v", t_.Header["alg"])

		}
		return []byte(j.secretKey), nil
	})
}
