package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rpolnx/rinha-backend-go/internal/dto"
	"github.com/rpolnx/rinha-backend-go/internal/service"
)

type PersonController interface {
	CreatePerson(c *gin.Context)
	GetPersonById(c *gin.Context)
	GetAllPeople(c *gin.Context)
	CountAllPeople(c *gin.Context)
}

type personController struct {
	svc service.PersonService
}

func (p personController) CreatePerson(c *gin.Context) {
	fmt.Println("start CreatePerson")
	var req dto.PersonDTO

	if err := c.BindJSON(&req); err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	entity := req.ToEntity()
	id, err := p.svc.CreatePerson(entity)

	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.Header("Location", fmt.Sprintf("/pessoas/%s", id))

	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusUnprocessableEntity)
		return
	}
	

	c.Status(201)
}

func (p personController) GetPersonById(c *gin.Context) {
	inputId := c.Param("id")

	id, err := uuid.Parse(inputId)

	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	person, err := p.svc.GetPersonById(id)

	if err != nil {
		if err.Error() == "record not found" {
			c.Status(http.StatusNotFound)
			return
		}
		fmt.Println(err)
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	personDto := dto.PersonEntityToDTO(person)

	c.JSON(http.StatusOK, personDto)
}
func (p personController) GetAllPeople(c *gin.Context) {
	query := c.Query("t")

	people, err := p.svc.GetAllPeople(query)

	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	dtoList := make([]*dto.PersonDTO, 0)

	for _, entity := range people {
		dtoList = append(dtoList, dto.PersonEntityToDTO(&entity))
	}

	c.JSON(http.StatusOK, dtoList)
}
func (p personController) CountAllPeople(c *gin.Context) {
	value, err := p.svc.CountAllPeople()

	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.String(200, strconv.FormatInt(int64(value), 10))
}

func NewPersonController(svc service.PersonService) PersonController {
	return &personController{
		svc,
	}
}
