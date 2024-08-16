package infrastructure

import (
	"os"
	"task-manager/domain"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(user *domain.User) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type JWTCustomClaims struct {
	UserID   string `json:"userId"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	secretKey := os.Getenv("JWT_SECERET")
	if secretKey == "" {
		secretKey = "defaultsecretkey"
	}

	return &jwtService{
		secretKey: secretKey,
		issuer:    "taskManagerApi",
	}

}

func (j *jwtService) GenerateToken(user *domain.User) (string, error) {
	claims := &JWTCustomClaims{
		UserID:   user.ID.Hex(),
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))

}

func (j *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
}
