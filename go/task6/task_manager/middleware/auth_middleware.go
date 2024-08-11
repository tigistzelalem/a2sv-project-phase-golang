package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"task_manager/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJwt(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user._id": user.ID,
		"username": user.Username,
		"role": user.Role,
		"exp": time.Now().Add(time.Hour * 24 *3 ).Unix(),
		
	})
	newToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	} 

	return newToken, nil
}


func ValidateToken(tokenValue string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwtSecret, nil
		

	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Store user information from token into context
			ctx.Set("userID", claims["id"])
			ctx.Set("username", claims["username"])
			ctx.Set("role", claims["role"])
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}


func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		userClaims := claims.(jwt.MapClaims)
		role, ok := userClaims["role"].(string)
		if !ok || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}