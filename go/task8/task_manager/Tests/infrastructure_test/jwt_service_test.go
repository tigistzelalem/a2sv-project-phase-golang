package infrastructuretest_test

import (
	"task-manager/domain"
	"task-manager/infrastructure"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTServiceTestSuite struct {
	suite.Suite
	jwtSvc infrastructure.JWTService
	user   *domain.User
}

func (suite *JWTServiceTestSuite) SetupTest() {
	suite.jwtSvc = infrastructure.NewJWTService()
	suite.user = &domain.User{
		ID:       primitive.NewObjectID(),
		Username: "testuser",
		Role:     "user",
	}
}

func (suite *JWTServiceTestSuite) TestGenerateToken() {
	token, err := suite.jwtSvc.GenerateToken(suite.user)
	suite.NoError(err)
	suite.NotEmpty(token)
}

func (suite *JWTServiceTestSuite) TestValidateToken() {
	token, _ := suite.jwtSvc.GenerateToken(suite.user)
	validatedToken, err := suite.jwtSvc.ValidateToken(token)

	suite.NoError(err)
	suite.NotNil(validatedToken)
	suite.True(validatedToken.Valid)

	claims, ok := validatedToken.Claims.(*infrastructure.JWTCustomClaims)
	suite.True(ok)
	suite.Equal(suite.user.ID.Hex(), claims.UserID)
}

func TestJWTServiceSuite(t *testing.T) {
	suite.Run(t, new(JWTServiceTestSuite))
}
