package service

import (
	todo "MaksJash"
	"MaksJash/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(rep repository.TodoList) *TodoListService {
	return &TodoListService{
		repo: rep,
	}
}

// Передача данных на некст слой - создание списка
func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

// Передача данных на некст слой - выдача всех списков
func (s *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userId)
}

// Передача данных на некст слой - выдача списка по переданому ID
func (s *TodoListService) GetListById(user_id, list_id int) (todo.TodoList, error) {
	return s.repo.GetListById(user_id, list_id)
}

// Передача данных на некст слой - удаление списка по переданому ID
func (s *TodoListService) DeleteListById(user_id, list_id int) error {
	return s.repo.DeleteListById(user_id, list_id)
}

// Передача данных на некст слой - обновление списка по переданому ID
func (s *TodoListService) Update(user_id, list_id int, input todo.UpdateListInput) error {
	// Валидация
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(user_id, list_id, input)
}
