package handlers

import (
	"My_pet_project/internal/tasksService"
	"My_pet_project/internal/web/tasks"
	"golang.org/x/net/context"
)

type Handler struct {
	Service *tasksService.TaskService
}

func NewHandler(service *tasksService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех задач из сервиса
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.Id,
			Task:   &tsk.Text,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := tasksService.Task{
		Text:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.Id,
		Task:   &createdTask.Text,
		IsDone: &createdTask.IsDone,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	// Удаляем задачу по ID
	err := h.Service.DeleteTaskByID(uint(request.Id))
	if err != nil {
		// Обработайте ошибку - например, если задача не найдена
		return nil, err
	}

	// Возвращаем успешный ответ (можно вернуть пустой ответ)
	return tasks.DeleteTasksId204Response{}, nil
}

func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	// Получение данных для обновления
	updatedTask := tasksService.Task{}
	if request.Body.Task != nil {
		updatedTask.Text = *request.Body.Task
	}
	if request.Body.IsDone != nil {
		updatedTask.IsDone = *request.Body.IsDone
	}

	// Обновляем задачу
	taskId := request.Id
	updatedTask, err := h.Service.UpdateTaskByID(uint(taskId), updatedTask)
	if err != nil {
		// Обработайте ошибку - например, если задача не найдена
		return nil, err
	}

	// Успешно обновленная задача
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.Id,
		Task:   &updatedTask.Text,
		IsDone: &updatedTask.IsDone,
	}
	return response, nil
}
