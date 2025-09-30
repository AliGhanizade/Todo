package model

import (
	"gorm.io/gorm"
	"todo/config"
)

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}

func (m *Task) TableName() string {
	return "tasks"
}

func (t *Task) Create() error {
	if err := config.Db.Model(t).Create(t).Error; err != nil {
		return err
	}
	return nil
}

func (t *Task) Update(id uint) error {
	if err := config.Db.Model(t).Where("id = ?", id).Updates(t).Error; err != nil {
		return err
	}
	return nil
}

func (t *Task) GetAll() ([]Task, error) {
	var tasks []Task
	if err := config.Db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *Task) GetByID(id uint) error {
	if err := config.Db.Model(t).Where("id = ?", id).First(t).Error; err != nil {
		return err
	}
	return nil
}

func (t *Task) Delete(id uint) error {
	if err := config.Db.Model(t).Where("id = ?", id).Delete(t).Error; err != nil {
		return err
	}
	return nil
}

func (t *Task) Search(title string, isComplete bool) ([]Task, error) {
	var tasks []Task
	dbQuery := config.Db.Model(&Task{})

	if title != "" {
		dbQuery = dbQuery.Where("title LIKE ?", "%"+title+"%")
	}

	if isComplete {
		dbQuery = dbQuery.Where("is_completed = ?", true)
	} else {
		dbQuery = dbQuery.Where("is_completed = ?", false)
	}

	if err := dbQuery.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *Task) Pagination(page, limit int) ([]Task, error) {
	var tasks []Task
	if err := config.Db.Offset((page - 1) * limit).Limit(limit).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}