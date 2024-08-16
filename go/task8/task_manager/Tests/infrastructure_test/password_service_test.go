package infrastructuretest_test

import (
	"task-manager/infrastructure"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PasswordServiceTestSuite struct {
	suite.Suite
	passwordSvc infrastructure.PasswordService
}

func (suite *PasswordServiceTestSuite) SetupTest() {
	suite.passwordSvc = infrastructure.NewPasswordService()
}

func (suite *PasswordServiceTestSuite) TestHashPassword() {
	hash, err := suite.passwordSvc.HashPassword("password123")
	suite.NoError(err)
	suite.NotEmpty(hash)
}

func (suite *PasswordServiceTestSuite) TestCheckPasswordHashSuccess() {
	hash, _ := suite.passwordSvc.HashPassword("password123")
	err := suite.passwordSvc.CheckPasswordHash("password123", hash)
	suite.NoError(err)
}

func (suite *PasswordServiceTestSuite) TestCheckPasswordHashFailure() {
	hash, _ := suite.passwordSvc.HashPassword("password123")
	err := suite.passwordSvc.CheckPasswordHash("wrongpassword", hash)
	suite.Error(err)
}

func TestPasswordServiceSuite(t *testing.T) {
	suite.Run(t, new(PasswordServiceTestSuite))
}
