package task

import (
	"net/http"
	"todo/share"
	"todo/model"


	"github.com/gin-gonic/gin"
)
/*
Delete a task by ID:

Router /tasks/{id} [delete]
*/
func (t *TaskController) Delete(ctx *gin.Context) {
	var task model.Task
	var uri share.TaskURI

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		share.NewError(http.StatusBadRequest, err.Error(), ctx)
		return
	}

	err = task.Delete(uri.ID)
	if err != nil {
		share.NewError(http.StatusInternalServerError, err.Error(), ctx)
		return
	}

	share.NewResponse(http.StatusOK, "Task deleted successfully", nil, ctx)
}
