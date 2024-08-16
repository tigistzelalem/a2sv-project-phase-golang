package tests

import (
	"task-manager/domain"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DomainTestSuite struct {
	suite.Suite
}

func (suite *DomainTestSuite) SetUpSuite() {

}

func (suite *DomainTestSuite) TestUserEntity() {
	user := domain.User{
		Username: "testusername",
		Password: "testpassword",
	}

	suite.Equal("testusername", user.Username)
	suite.Equal("testpassword", user.Password)
}

func (suite *DomainTestSuite) TestTaskEntity() {
	task := domain.Task{
		Title:       "title1",
		Description: "this is the first task",
	}

	suite.Equal("title1", task.Title)
	suite.Equal("this is the first task", task.Description)

}

func TestDomainSuite(t *testing.T) {
	suite.Run(t, new(DomainTestSuite))
}
