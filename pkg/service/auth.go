package service

import (
	todo "MaksJash"
	"MaksJash/pkg/repository"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	saltForPass string        = "alsdfljsld2084alkdj"           // Соль для пароля
	signingKey  string        = "jlh3w8oisalkfas0aq13q2as;fj02" //
	tokenTTL    time.Duration = 12 * time.Hour                  // Время жизни жвт токена

)

// Кастомные клэймсы для токена
type customTokenClaims struct {
	jwt.StandardClaims     // Встраиваем
	UserId             int `json:"user_id"`
}

type AuthService struct {
	rep repository.Authorization // Доступ к интерфейсу Authorization на репе
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

func (s *AuthService) GenerateToken(username, password string) (string, error) {

	// Вытаскиваем пользователя
	user, err := s.rep.GetUser(username, generatePasswordHash(password)) // Пароль передается захешированным
	if err != nil {
		return "", err
	}

	// Генерация токена
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &customTokenClaims{ // Подпись - SHA256;
		// Также, здесь используются кастомные claims'ы

		jwt.StandardClaims{ // Встроенные стандартные claims'ы с 2 определенными полями
			ExpiresAt: time.Now().Add(tokenTTL).Unix(), // Время истечения токена
			IssuedAt:  time.Now().Unix(),               // Время генерации токена
		},

		user.Id, // id пользователя
	})

	return token.SignedString([]byte(signingKey)) // Подписываем токен заготовленным ключом

}

// Генерит и солит хеш пароля
func generatePasswordHash(password string) string {

	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(saltForPass))) // Sum вычисляет "финальный" хеш на основе того,
	// что ему было передано через write

}
