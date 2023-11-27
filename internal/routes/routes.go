package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rpolnx/rinha-backend-go/internal/controller"
)

type RouterBuilder struct {
	server           *gin.Engine
	personController controller.PersonController
}

func (r RouterBuilder) AppendRoutes() {
	pessoas := r.server.Group("/pessoas")
	{
		pessoas.POST("", r.personController.CreatePerson)
		pessoas.GET(":id", r.personController.GetPersonById)
		pessoas.GET("", r.personController.GetAllPeople)
	}

	r.server.GET("/contagem-pessoas", r.personController.CountAllPeople)
}

func NewRouterBuilder(server *gin.Engine, personController controller.PersonController) *RouterBuilder {
	return &RouterBuilder{
		server,
		personController,
	}
}
