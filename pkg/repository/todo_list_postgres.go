package repository

import (
	todo "MaksJash"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func newTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{
		db: db,
	}
}

// Создает новый список в БД todo_lists и связь userId и listId в таблице users_lists
func (t *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error) {

	// Начинаем транзакцию
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	// 1-ая операция транзакции - создание записи в таблице todo Lists
	var id int // Сюда придет id нового списка
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)

	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback() // Откат транзакции
		return 0, err
	}

	// 2-ая операция транзакции - вставка в Users List, что связывает id поль-ля и id нового списка
	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)

	_, err = tx.Exec(createUsersListQuery, userId, id) // Поначалу вставил вместо ", id" ", list.Id"
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit() // Не знал, что так можно
}

// Дергает все списки из бд по переданному id пользователя
func (t *TodoListPostgres) GetAll(userId int) ([]todo.TodoList, error) {

	var user_lists []todo.TodoList

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s 
	tl INNER JOIN %s ul on	tl.id = ul.list_id WHERE ul.user_id = $1`, todoListsTable, usersListsTable)
	// tl и ul - алиасы. Вот как они задаются: "Имя таблицы alias" - в данном случае вместо имени таблицы -
	// placeholders.

	err := t.db.Select(&user_lists, query, userId)
	if err != nil {
		return nil, err
	}

	return user_lists, err

}

func (t *TodoListPostgres) GetListById(user_id, list_id int) (todo.TodoList, error) {

	var user_list todo.TodoList

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl 
	INNER JOIN %s ul on	tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`, todoListsTable,
		usersListsTable)

	err := t.db.Get(&user_list, query, user_id, list_id)
	if err != nil {
		return todo.TodoList{}, err // Пустой. А как иначе? Я хз можно ли кастануть nil к toto.TodoList
	}

	return user_list, err

}
