//go:generate mockery --name PersonRepository --inpackage --filename=person_repository_mock.go --structname=PersonRepositoryMock
package repository

import (
	"github.com/google/uuid"
	"github.com/rpolnx/rinha-backend-go/internal/model"
	"gorm.io/gorm"
)

type PersonRepository interface {
	InsertPerson(entity *model.PersonEntity) (uuid.UUID, error)
	FindPersonById(id uuid.UUID) (*model.PersonEntity, error)
	FindAllPeople(query string) ([]model.PersonEntity, error)
	CountAllDbPeople() (int64, error)
}

type personRepository struct {
	db *gorm.DB
}

func (repo personRepository) InsertPerson(entity *model.PersonEntity) (uuid.UUID, error) {
	result := repo.db.
		Create(entity)

	return entity.Id, result.Error
}

func (repo personRepository) FindPersonById(id uuid.UUID) (*model.PersonEntity, error) {
	person := &model.PersonEntity{}

	result := repo.db.First(person, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return person, result.Error
}

func (repo personRepository) FindAllPeople(query string) ([]model.PersonEntity, error) {
	people := make([]model.PersonEntity, 0)

	result := repo.db.
		Where("? <% (coalesce(name, '') || ' ' || coalesce(username, '') || ' ' || immutable_array_to_string(coalesce(stack, '{}'), ' '))",
			query).
		Find(&people)
	// 	SELECT ... FROM tab
	// WHERE 'postgres' <% concat(title, ' ', description);

	return people, result.Error
}

func (repo personRepository) CountAllDbPeople() (int64, error) {
	var count int64

	result := repo.db.Model(&model.PersonEntity{}).Count(&count)

	return count, result.Error
}

func NewPersonRepository(db *gorm.DB) PersonRepository {
	return &personRepository{db}
}
