package infrastructuretest_test

import (
	"net/http"
	"net/http/httptest"
	"task-manager/domain"
	"task-manager/infrastructure"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	// "github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthMiddlewareTestSuite struct {
	suite.Suite
	jwtService infrastructure.JWTService
}

type mockJWTService struct {
	token string
	valid bool
	err   error
}

func (m *mockJWTService) GenerateToken(user *domain.User) (string, error) {
	return m.token, nil

}

func (m *mockJWTService) ValidateToken(tokenString string) (*jwt.Token, error) {
	if m.err != nil {
		return nil, m.err
	}

	return &jwt.Token{Valid: m.valid, Claims: &infrastructure.JWTCustomClaims{

		UserID:   primitive.NewObjectID().Hex(),
		Username: "testuser",
		Role:     "user",
	}}, nil

}

func (suite *AuthMiddlewareTestSuite) SetupTest() {
	suite.jwtService = &mockJWTService{}
	gin.SetMode(gin.TestMode)

}

func (suite *AuthMiddlewareTestSuite) TestNoAuthorizationHeader() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/", nil)

	middleware := infrastructure.AuthMiddleware(suite.jwtService)
	middleware(c)

	suite.Equal(http.StatusUnauthorized, w.Code)
	suite.JSONEq(`{"error": "Authorization header required"}`, w.Body.String())

}

func (suite *AuthMiddlewareTestSuite) TestInvalidToken() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
	c.Request.Header.Set("Authorization", "Bearer invalidtoken")

	suite.jwtService.(*mockJWTService).valid = false
	middleware := infrastructure.AuthMiddleware(suite.jwtService)
	middleware(c)

	suite.Equal(http.StatusUnauthorized, w.Code)
	suite.JSONEq(`{"error": "Invalid or expired token"}`, w.Body.String())

}

func (suite *AuthMiddlewareTestSuite) TestValidToken() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
	c.Request.Header.Set("Authorization", "Bearer validtoken")

	suite.jwtService.(*mockJWTService).valid = true

	middleware := infrastructure.AuthMiddleware(suite.jwtService)
	middleware(c)

	suite.Equal(http.StatusOK, w.Code)
	suite.Equal("testuser", c.GetString("username"))
}

func TestAuthMiddlewareSuite(t *testing.T) {
	suite.Run(t, new(AuthMiddlewareTestSuite))
}
