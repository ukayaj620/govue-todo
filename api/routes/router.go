package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ukayaj620/govue-todo/api/controllers"
)

func InitRouter() {
	router := gin.Default()

	todoRouter := router.Group("/todo")
	{
		todoRouter.GET("/all", controllers.FetchAllTodo)
		todoRouter.POST("/create", controllers.CreateTodo)
		todoRouter.PUT("/update/:id", controllers.UpdateTodo)
		todoRouter.DELETE("/delete/:id", controllers.DeleteTodo)
	}

	router.Run()
}
