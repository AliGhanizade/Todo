package task

import (
	"net/http"
	"strconv"
	"todo/model"
	"todo/share"

	"github.com/gin-gonic/gin"
)

/*
Show all Tasks

Router /tasks [get]
*/
func (t *TaskController) GetAll(ctx *gin.Context) {
	title  := ctx.Query("title")
	status := ctx.Query("is_completed")

	if title != "" || status != "" {
		t.Search(title, status, ctx)
		return
	}

	limit := 5
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		share.NewError(http.StatusBadRequest, err.Error(), ctx)
		return
	}

	t.Pagination(page, limit, ctx)
}

/*
Show a task by ID

Router /tasks/{id} [get]
*/
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
		if err.Error() == "record not found" {
			share.NewNotFound(http.StatusNotFound, "Task not found", ctx)
			return
		}
		share.NewError(http.StatusInternalServerError, err.Error(), ctx)
		return
	}

	if task.ID == 0 {
		share.NewNotFound(http.StatusNotFound, "Task not found", ctx)
		return
	}

	share.NewResponse(http.StatusOK, "Task retrieved successfully", task, ctx)
}
