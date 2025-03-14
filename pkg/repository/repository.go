package repository

import "github.com/jmoiron/sqlx"

// Тут работа с БД

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

// Собирает интерфейсы в 1-м месте
type Repository struct {
	// Это композитная структура. Благодаря такому подходу можно управлять разными интерфейсами
	// (по факту - сервисами приложений) из 1-го места
	// (по факту - сервисами приложений) из 1-го места
	// Благодаря такому встраиванию, методы интерфейсов будут доступны прямо из структуры
	Authorization
	TodoItem
	TodoList
}

// Конструктор
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
