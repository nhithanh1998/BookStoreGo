package resources

import (
	"bookstore/handles"

	"github.com/gin-gonic/gin"
)

// InitCategories is a ...
func InitCategories(router *gin.Engine) {
	categories := router.Group("/genres")
	{
		categories.GET("/", handles.GetGenres)
	}
}
