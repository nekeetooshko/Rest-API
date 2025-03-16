package repository

// Здесь уже непосредственно работа с БД

import (
	todo "MaksJash"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Имплементит интерфейс репозитория и работает с постгре
type AuthPostgres struct {
	db *sqlx.DB
}

func newAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

// Добавляет пользователя в бд и возвращает его id
func (a *AuthPostgres) CreateUser(user todo.User) (int, error) {

	var id int // Заготовка под id

	// Подготовка
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	// Выполнение
	row := a.db.QueryRow(query, user.Name, user.Username, user.Password) // Вернет id, т.к. это мы указали в запросе

	if err := row.Scan(&id); err != nil { // Кладем id в заготовку
		return 0, err
	}

	return id, nil

}
