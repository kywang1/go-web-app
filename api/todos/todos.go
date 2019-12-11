package todos

import "github.com/gin-gonic/gin"

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/todos")
	{
		api.POST("/", create)

		api.GET("/", list)
		api.GET("/:id", get)

		api.DELETE("/:id", remove)

		api.PATCH("/:id", update)
	}
}
