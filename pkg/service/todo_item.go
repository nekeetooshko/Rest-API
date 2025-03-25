package service

import (
	todo "MaksJash"
	"MaksJash/pkg/repository"
)

type TodoItemService struct { // 2 поля, т.к. нужен доступ и к спискам, и к их элементам
	item_repo repository.TodoItem
	list_repo repository.TodoList
}

func NewTodoItemService(item_repo repository.TodoItem, list_repo repository.TodoList) *TodoItemService {
	return &TodoItemService{
		item_repo: item_repo,
		list_repo: list_repo,
	}
}

func (s *TodoItemService) Create(user_id, list_id int, item todo.TodoItem) (int, error) {

	_, err := s.list_repo.GetListById(user_id, list_id)
	if err != nil { // Список не сущ, или не принадлежит юзеру
		return 0, err
	}

	return s.item_repo.Create(list_id, item)

}

func (s *TodoItemService) GetAll(user_id, list_id int) ([]todo.TodoItem, error) {
	return s.item_repo.GetAll(user_id, list_id)
}

func (s *TodoItemService) GetItemById(user_id, item_id int) (todo.TodoItem, error) {
	return s.item_repo.GetItemById(user_id, item_id)
}

func (s *TodoItemService) Delete(user_id, item_id int) error {
	return s.item_repo.Delete(user_id, item_id)
}

func (s *TodoItemService) Update(user_id, item_id int, input todo.UpdateItemInput) error {
	return s.item_repo.Update(user_id, item_id, input)
}
