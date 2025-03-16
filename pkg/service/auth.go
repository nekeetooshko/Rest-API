package service

import (
	todo "MaksJash"
	"MaksJash/pkg/repository"
	"crypto/sha1"
	"fmt"
)

const saltForPass string = "alsdfljsld2084alkdj" // Соль для пароля

type AuthService struct {
	rep repository.Authorization // Доступ к интерфейсу Authorization
}

// Конструктор для AuthService
func newAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo, // Можно и так
	}
}

// Передает пользователя еще ниже - в репозиторий
func (s *AuthService) CreateUser(user todo.User) (int, error) {

	user.Password = generatePasswordHash(user.Password)

	return s.rep.CreateUser(user)
}

// Генерит и солит хеш пароля
func generatePasswordHash(password string) string {

	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(saltForPass))) // Sum вычисляет "финальный" хеш на основе того,
	// что ему было передано через write

}
