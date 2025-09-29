package task

import (
	"net/http"
	"todo/model"
	"todo/share"

	"github.com/gin-gonic/gin"
)

func (t *TaskController) GetAll(ctx *gin.Context) {
	var task model.Task
	tasks, err := task.GetAll()
	if err != nil {
		share.NewError(http.StatusInternalServerError, err.Error(), ctx)
		return
	}
	share.NewResponse(http.StatusOK, "Tasks retrieved successfully", tasks, ctx)
}

func (t *TaskController) GetByID(ctx *gin.Context) {
	var task model.Task
	var uri share.TaskURI

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		share.NewError(http.StatusBadRequest, err.Error(), ctx)
		return
	}

	err = task.GetByID(uri.ID)
	if err != nil {
		share.NewError(http.StatusInternalServerError, err.Error(), ctx)
		return
	}

	share.NewResponse(http.StatusOK, "Task retrieved successfully", task, ctx)
}
