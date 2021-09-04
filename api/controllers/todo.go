package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ukayaj620/govue-todo/api/db"
	"github.com/ukayaj620/govue-todo/api/models"
)

func FetchAllTodo(c *gin.Context) {
	var todo []models.Todo

	result := db.DB.Find(&todo)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": result.Error.Error(),
		})
	}

	c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)

	result := db.DB.Create(&todo)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": result.Error.Error(),
		})
	}

	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")

	result := db.DB.Where("id = ?", id).Updates(c.BindJSON(&todo))

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": result.Error.Error(),
		})
	}

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")

	result := db.DB.Where("id = ?", id).Delete(&todo)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": result.Error.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data deleted successfully",
	})
}
