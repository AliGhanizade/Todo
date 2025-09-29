package task

import (
	
	"net/http"
	"todo/model"
	"todo/share"

	"github.com/gin-gonic/gin"
)

/*
Update a task by ID:

Router /task/{id} [put]
{
	"title": "Updated Task",
	"description": "This is an updated task",
	"completed": true
}
*/



func (t *TaskController) Update(ctx *gin.Context) {
	var task model.Task
	var uri share.TaskURI

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		share.NewError(http.StatusBadRequest, err.Error(), ctx)
		return
	}

	err = ctx.ShouldBindBodyWithJSON(&task)
	if err != nil {
		share.NewError(http.StatusBadRequest, err.Error(), ctx)
		return
	}
	err = task.Update(uri.ID)
	if err != nil {
		share.NewError(http.StatusInternalServerError, err.Error(), ctx)
		return
	}

	share.NewResponse(http.StatusOK, "Task updated successfully", task, ctx)
}
