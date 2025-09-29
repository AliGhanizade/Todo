package router

import (
	"todo/task"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	taskC := task.TaskController{}
	task := r.Group("/task")
	{
		task.POST("/", taskC.Create)
		task.GET("/", taskC.GetAll)
		task.GET("/:id/", taskC.GetByID)
		task.PUT("/:id/", taskC.Update)
		task.DELETE("/:id/", taskC.Delete)
	}

	return r

}
