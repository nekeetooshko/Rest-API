package todo

// Описание всех необходимых сущностей (в соответствии с БД)
type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"itlte"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
