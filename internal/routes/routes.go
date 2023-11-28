package route

import (
	"github.com/gin-gonic/gin"
	"github.com/rpolnx/rinha-backend-go/internal/controller"
	"github.com/samber/do"
)

func RegisterRoutes(injector *do.Injector) {
	server := do.MustInvoke[*gin.Engine](injector)

	personController := do.MustInvoke[controller.PersonController](injector)

	pessoas := server.Group("/pessoas")
	{
		pessoas.POST("", personController.CreatePerson)
		pessoas.GET(":id", personController.GetPersonById)
		pessoas.GET("", personController.GetAllPeople)
	}

	server.GET("/contagem-pessoas", personController.CountAllPeople)
}
