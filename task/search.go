package task

import (
	"net/http"
	"strconv"
	"todo/model"
	"todo/share"

	"github.com/gin-gonic/gin"
)

func (t *TaskController) Search(title, status string, ctx *gin.Context) {
	var task model.Task

	isComplete, err := strconv.ParseBool(status)
	if err != nil {
		share.NewError(http.StatusBadRequest, err.Error(), ctx)
		return
	}
	tasks, err := task.Search(title, isComplete)

	if err != nil {
		share.NewError(http.StatusInternalServerError, err.Error(), ctx)
		return
	}

	if len(tasks) == 0 {
		share.NewError(http.StatusNotFound, "No tasks found", ctx)
		return
	}

	share.NewResponse(http.StatusOK, "Tasks retrieved successfully", tasks, ctx)
}