package model

import (
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	if err := db.AutoMigrate(&Task{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	return db
}
func TestTaskCRUD(t *testing.T) {
	db := setupTestDB(t)

	task := Task{Title: "test", Description: "desc", IsCompleted: false}

	// Create
	if err := db.Create(&task).Error; err != nil {
		t.Fatalf("failed to create: %v", err)
	}

	var got Task

	// Read
	if err := db.First(&got, task.ID).Error; err != nil {
		t.Fatalf("failed to get: %v", err)
	}
	if got.Title != "test" {
		t.Errorf("expected title test, got %s", got.Title)
	}

	// Update
	got.Title = "updated"
	if err := db.Save(&got).Error; err != nil {
		t.Fatalf("failed to update: %v", err)
	}

	if err := db.First(&got, task.ID).Error; err != nil {
		t.Fatalf("failed to get: %v", err)
	}

	if got.Title != "updated" {
		t.Errorf("expected updated title, got %s", got.Title)
	}

	// Delete
	if err := db.Delete(&got).Error; err != nil {
		t.Fatalf("failed to delete: %v", err)
	}

	if got.DeletedAt.Time.IsZero() {
		t.Errorf("expected deleted at to be set, got %v", got.DeletedAt)
	}
}

func TestTasksCreate(t *testing.T) {
	db := setupTestDB(t)

	tests := []struct {
		name string
		task Task
	}{
		{"valid task 1 ", Task{Title: "task 1", Description: "Desc 1", IsCompleted: false}},
		{"valid task 2 ", Task{Title: "task 2", Description: "Desc 2"}},
		{"valid task 3 ", Task{Title: "task 3"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := db.Model(&tt.task).Create(&tt.task).Error
			if err != nil {
				t.Errorf("failed to create: %v", err)
			} else {
				t.Logf("created task with data: %s", tt.name)
			}
		})
	}

}
func TestTasksUpdate(t *testing.T) {
	db := setupTestDB(t)

	tests := []struct {
		name string
		task Task
	}{
		{"valid task 1 ", Task{Title: "up task 1", Description: "up Desc 1", IsCompleted: true}},
		{"valid task 2 ", Task{Title: "up task 2", Description: "up Desc 2"}},
		{"valid task 3 ", Task{Title: "up task 3"}},
		{"valid task 4 ", Task{Description: "up Task 4 and doesn't have a title"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := db.Model(&tt.task).Where("id = ?", tt.task.ID).Updates(&tt.task).Error

			if err != nil {
				t.Errorf("failed to update: %v", err)
			} else {

				t.Logf("updating task with data: %s", tt.name)
			}

		})
	}

}
