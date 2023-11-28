//go:generate mockery --name PersonService --inpackage --filename=person_service_mock.go --structname=PersonServiceMock
package service

import (
	"github.com/google/uuid"
	"github.com/rpolnx/rinha-backend-go/internal/model"
	"github.com/rpolnx/rinha-backend-go/internal/repository"
	"github.com/samber/do"
)

type PersonService interface {
	CreatePerson(entity *model.PersonEntity) (uuid.UUID, error)
	GetPersonById(id uuid.UUID) (*model.PersonEntity, error)
	GetAllPeople(query string) ([]model.PersonEntity, error)
	CountAllPeople() (int64, error)
}

type personService struct {
	personRepo repository.PersonRepository
}

func (p personService) CreatePerson(entity *model.PersonEntity) (uuid.UUID, error) {
	entity.Id = uuid.New()

	return p.personRepo.InsertPerson(entity)
}

func (p personService) GetPersonById(id uuid.UUID) (*model.PersonEntity, error) {
	return p.personRepo.FindPersonById(id)
}

func (p personService) GetAllPeople(query string) ([]model.PersonEntity, error) {
	return p.personRepo.FindAllPeople(query)
}

func (p personService) CountAllPeople() (int64, error) {
	return p.personRepo.CountAllDbPeople()
}

func NewPersonService(injector *do.Injector) (PersonService, error) {
	personRepo := do.MustInvoke[repository.PersonRepository](injector)

	return &personService{personRepo}, nil
}
