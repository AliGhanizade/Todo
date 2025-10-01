package task

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/config"
	"todo/model"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)
func SeedTestData() {
    config.Db.Create(&model.Task{Title: "Task 1", Description: "First task"})
    config.Db.Create(&model.Task{Title: "Task 2", Description: "Second task", IsCompleted: true})
    config.Db.Create(&model.Task{Title: "Task 3", Description: "Third task"})
}

func setupTestDB(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	if err := db.AutoMigrate(&model.Task{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	config.Db = db
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	taskC := TaskController{}
	r.POST("/tasks/", taskC.Create)
	r.GET("/tasks/", taskC.GetAll)
	r.GET("/tasks/:id/", taskC.GetByID)
	r.PUT("/tasks/:id/", taskC.Update)
	r.DELETE("/tasks/:id/", taskC.Delete)

	return r
}
func TestCreateTask(t *testing.T) {
	r := setupRouter()
	setupTestDB(t)

	body := `{"title": "test task", "description": "test desc", "is_completed": false}`
	req := httptest.NewRequest("POST", "/tasks/", bytes.NewBufferString(body))

	w := httptest.NewRecorder()
	r.ServeHTTP(w,req)
	if w.Code != http.StatusCreated {
		t.Errorf("error in creating",)
	}
}

func TestUpdateTask(t *testing.T) {
	r := setupRouter()
	setupTestDB(t)

	body := `{"title": "test task", "description": "test desc", "is_completed": false}`
	req := httptest.NewRequest("PUT", "/tasks/1/", bytes.NewBufferString(body))

	w := httptest.NewRecorder()

	r.ServeHTTP(w,req)

	if w.Code != http.StatusOK {
		t.Errorf("error in updating",)
	}
}

func TestGetAllTask(t *testing.T) {
	r := setupRouter()
	setupTestDB(t)
	SeedTestData()

	req := httptest.NewRequest("GET", "/tasks/", nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w,req)

	if w.Code != http.StatusOK {
		t.Errorf("error in Get All",)
	}
}