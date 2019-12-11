package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kywang1/go-web-app/api/todos"
)

func ApplyRoutes(r *gin.Engine) {
	todosRoute := r.Groupt("/api")
	todos.ApplyRoutes(todosRoute)
}
