package controllers

import (
	"downloader/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) Retrieve(c *gin.Context) {
	if c.Param("id") != "" {
		user, err := userModel.GetByID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User founded!", "user": user})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"meesage": "bad request"})
	c.Abort()
	return
}
