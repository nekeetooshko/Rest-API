package todo

// Описание сущности пользователя
type User struct {
	Id       int    `json:"-"` // Хз че так
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
