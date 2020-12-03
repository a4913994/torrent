package controllers

import (
	"downloader/models"
	"downloader/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) Retrieve(c *gin.Context) {
	id := utils.StrConvertUint(c.Param("id"))
	if id > 0 {
		user, err := userModel.GetByID(id)
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
