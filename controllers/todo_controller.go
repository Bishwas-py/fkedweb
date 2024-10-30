package controllers

import (
	"fkedweb/models"
	"fkedweb/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	service services.TodoService
}

func NewTodoController(service services.TodoService) *TodoController {
	return &TodoController{service}
}

func (c *TodoController) Create(ctx *gin.Context) {
	var todo models.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.Create(&todo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, todo)
}

func (c *TodoController) GetAll(ctx *gin.Context) {
	todos, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

func (c *TodoController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	todo, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	ctx.JSON(http.StatusOK, todo)
}

func (c *TodoController) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var todo models.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo.ID = uint(id)
	if err := c.service.Update(&todo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, todo)
}

func (c *TodoController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.service.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
