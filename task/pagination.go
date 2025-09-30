package task

import (
	"net/http"
	"todo/model"
	"todo/share"

	"github.com/gin-gonic/gin"
)

func (t *TaskController) Pagination(page, limit int, ctx *gin.Context) {
	var model model.Task

	tasks, err := model.Pagination(page, limit)
	if err != nil {
		share.NewError(http.StatusInternalServerError, err.Error(), ctx)
		return
	}

	if len(tasks) == 0 {
		share.NewNotFound(http.StatusNotFound, "Not Page found", ctx)
		return
	}
	share.NewResponse(http.StatusOK, "Tasks retrieved successfully", tasks, ctx)
}