package routes

import (
	"fkedweb/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(todoController *controllers.TodoController) *gin.Engine {
	r := gin.Default()
	todoRoutes := r.Group("/todos")
	{
		todoRoutes.POST("/", todoController.Create)
		todoRoutes.GET("/", todoController.GetAll)
		todoRoutes.GET("/:id", todoController.GetByID)
		todoRoutes.PUT("/:id", todoController.Update)
		todoRoutes.DELETE("/:id", todoController.Delete)
	}
	return r
}
