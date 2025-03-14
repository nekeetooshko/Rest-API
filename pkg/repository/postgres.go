package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Тут логика подключения БД, и имена таблиц в константах (нахуя то)

// БДшный конфиг
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDb(config Config) (*sqlx.DB, error) {

	// Открываем коннект с БД
	db, err := sqlx.Open("postgres", fmt.Sprintf( // Я реально дурею с этой передачей параметров конфига
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.DBName, config.Password, config.SSLMode,
	))

	if err != nil {
		return nil, err
	}

	// Чек
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
