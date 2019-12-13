package api

import (
	"github.com/kywang1/go-web-app/api/todos"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	todosRoute := r.Groupt("/api")
	todos.ApplyRoutes(todosRoute)
}
