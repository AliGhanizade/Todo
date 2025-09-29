package task

import (
	"net/http"
	"todo/model"
	"todo/share"

	"github.com/gin-gonic/gin"
)
/*
Create a task :

Router /tasks [post]
{
	"title": "Sample Task",
	"description": "This is a sample task",
	"completed": false
}

*/
func (t *TaskController) Create(ctx *gin.Context) {
	var task model.Task
	err := ctx.ShouldBindBodyWithJSON(&task)
	if err != nil {
		share.NewError(http.StatusBadRequest, err.Error(), ctx)
		return
	}
	err = task.Create()
	if err != nil {
		share.NewError(http.StatusInternalServerError, err.Error(), ctx)
		return
	}
	
	share.NewResponse(http.StatusCreated, "Task created successfully", task, ctx)
}
