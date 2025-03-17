package service

import (
	todo "MaksJash"
	"MaksJash/pkg/repository"
)

// Тут всякая бизнес - логика

// Отвечает за авторизацию
type Authorization interface {
	CreateUser(user todo.User) (int, error)                  // Создает пользователя, возвращает его id
	GenerateToken(username, password string) (string, error) // Возвращает сгенерированный токен
	ParseToken(token string) (int, error)                    // Возвращает id пользователя по токену
}

// Отвечает за списки
type TodoList interface{}

// Отвечает за элементы списков
type TodoItem interface{}

// Собирает интерфейсы в 1-м месте
type Service struct {
	// Это композитная структура. Благодаря такому подходу можно управлять разными интерфейсами
	// (по факту - сервисами приложений) из 1-го места
	// Благодаря такому встраиванию, методы интерфейсов будут доступны прямо из структуры
	Authorization
	TodoItem
	TodoList
}

// Конструктор
func NewService(rep *repository.Repository) *Service {
	// Т.к. сервисы обращаются к БД - нужно передать репозиторий, ведь работа с БД ложится именно на его плечи
	return &Service{
		Authorization: newAuthService(rep.Authorization),
	}
}
