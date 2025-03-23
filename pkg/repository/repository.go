package repository

import (
	todo "MaksJash"

	"github.com/jmoiron/sqlx"
)

// Тут работа с БД

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetListById(user_id, list_id int) (todo.TodoList, error)
}

type TodoItem interface{}

// Собирает интерфейсы в 1-м месте
type Repository struct {
	// Это композитная структура. Благодаря такому подходу можно управлять разными интерфейсами
	// (по факту - сервисами приложений) из 1-го места
	// Благодаря такому встраиванию, методы интерфейсов будут доступны прямо из структуры
	Authorization
	TodoItem
	TodoList
}

// Конструктор
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: newAuthPostgres(db),
		TodoList:      newTodoListPostgres(db),
	}
}
