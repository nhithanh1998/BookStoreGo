package handles

import (
	"bookstore/database"
	"bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetGenres is function
func GetGenres(c *gin.Context) {
	var genres []models.Genre
	database.BookStoreDB.Find(&genres)
	c.JSON(http.StatusOK, genres)
}
