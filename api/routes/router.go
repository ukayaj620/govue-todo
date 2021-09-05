package routes

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ukayaj620/govue-todo/api/controllers"
)

func InitRouter() {
	result := godotenv.Load()

	if result != nil {
		panic(result.Error())
	}

	router := gin.Default()

	todoRouter := router.Group("/todo")
	{
		todoRouter.GET("/all", controllers.FetchAllTodo)
		todoRouter.POST("/create", controllers.CreateTodo)
		todoRouter.PUT("/update/:id", controllers.UpdateTodo)
		todoRouter.DELETE("/delete/:id", controllers.DeleteTodo)
	}

	router.Use(cors.Default())

	router.Run(os.Getenv("SERVER_PORT"))
}
