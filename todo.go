package todo

import "errors"

// Описание всех необходимых сущностей (в соответствии с БД)
type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}
type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}
type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	// Если в запросе не будет какого-то из полей - nil
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

// Валидация. Если все значения пустые - не обновляемся.
func (u *UpdateListInput) Validate() error {

	if u.Description == nil && u.Title == nil {
		return errors.New("Updated structure has no values")
	}
	return nil
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (u *UpdateItemInput) Validate() error {

	if u.Description == nil && u.Title == nil && u.Done == nil {
		return errors.New("Updated structure has no values")
	}
	return nil
}
