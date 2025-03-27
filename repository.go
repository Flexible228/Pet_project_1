package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(task Task) (Task, error)

	GetAllTasks() ([]Task, error)

	UpdateTaskByID(id uint, task Task) (Task, error)

	DeleteTaskByID(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	var existingTask Task
	if err := r.db.First(&existingTask, id).Error; err != nil {
		return Task{}, err
	}

	existingTask.Task = task.Task
	existingTask.IsDone = task.IsDone

	if err := r.db.Save(&existingTask).Error; err != nil {
		return Task{}, err
	}

	return existingTask, nil

}

func (r *taskRepository) DeleteTaskByID(id uint) error {

	if err := r.db.Delete(&Task{}, id).Error; err != nil {
		return err
	}
	return nil
}
