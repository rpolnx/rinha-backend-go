package service

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/rpolnx/rinha-backend-go/internal/model"
	"github.com/rpolnx/rinha-backend-go/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type PersonServiceTestSuite struct {
	suite.Suite
	repository *repository.PersonRepositoryMock
	service    PersonService
}

func (suite *PersonServiceTestSuite) SetupTest() {
	suite.repository = repository.NewPersonRepositoryMock(suite.T())
	suite.service = NewPersonService(suite.repository)
}

func (suite *PersonServiceTestSuite) Test_CreatePerson_ErrorCreatingPerson() {
	expectedErr := "error inserting person on DB"

	entity := model.PersonEntity{}

	suite.repository.On("InsertPerson", mock.Anything).Return(uuid.New(), errors.New(expectedErr))

	_, receivedError := suite.service.CreatePerson(&entity)

	assert.ErrorContains(suite.T(), receivedError, expectedErr)
}

func (suite *PersonServiceTestSuite) Test_CreatePerson_SuccessCreatingPerson() {
	entity := model.PersonEntity{Name: "xpto"}

	entity.Id = uuid.New()

	suite.repository.On("InsertPerson", mock.Anything).Return(entity.Id, nil)

	received, receivedError := suite.service.CreatePerson(&entity)

	assert.Nil(suite.T(), receivedError)

	entity.Id = received

	assert.Equal(suite.T(), entity.Id, received)
}

func Test_PersonServiceTestSuite(t *testing.T) {
	suite.Run(t, new(PersonServiceTestSuite))
}
