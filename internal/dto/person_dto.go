package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/rpolnx/rinha-backend-go/internal/model"
)

type PersonDTO struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"apelido", binding:"required,max=32"`
	Name     string    `json:"nome", binding:"required,max=100"`
	Birthday Datetime  `json:"nascimento", binding:"date=2006-01-02"`
	Stack    []string  `json:"stack"`
}

func (p *PersonDTO) ToEntity() *model.PersonEntity {
	return &model.PersonEntity{
		Username: p.Username,
		Name:     p.Name,
		Birthday: (time.Time)(p.Birthday),
		Stack:    p.Stack,
	}
}

func PersonEntityToDTO(entity *model.PersonEntity) *PersonDTO {
	return &PersonDTO{
		Id:       entity.Id,
		Username: entity.Username,
		Name:     entity.Name,
		Birthday: (Datetime)(entity.Birthday),
		Stack:    entity.Stack,
	}
}
