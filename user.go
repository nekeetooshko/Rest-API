package todo

// Описание сущности пользователя
type User struct {
	Id       int    `json:"-"`                       // Хз че так
	Name     string `json:"name" binding:"required"` // binding - встроенный в гин тег для валидации
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
