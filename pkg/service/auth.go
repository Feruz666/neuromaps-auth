package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/Feruz666/neuromaps-auth/models"
	"github.com/Feruz666/neuromaps-auth/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "ajsdfgakshkjerwuiycbakjb12312wqeq"
	signingKey = "asdklhjeklwjr12#52dksjf"
	tokenTTL   = 12 * time.Hour
)

// tokenClaims ...
type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

// AuthService ...
type AuthService struct {
	repo repository.Authorization
}

// NewAuthService ...
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// CreateUser ...
func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

// GenerateToken ...
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, nil
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil

}

// generatePasswordHash ...
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
