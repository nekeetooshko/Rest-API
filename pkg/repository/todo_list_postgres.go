package repository

import (
	todo "MaksJash"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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

// Дергает конкретный список пользователя, от которого прошла рега
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

// Удаляет список по ID
func (t *TodoListPostgres) DeleteListById(user_id, list_id int) error {

	query := fmt.Sprintf(`DELETE from %s tl using %s ul where tl.id = ul.list_id and ul.user_id = $1 
	and ul.list_id = $2`, todoListsTable, usersListsTable)
	// USING - использование другой таблицы в запросе

	_, err := t.db.Exec(query, user_id, list_id)
	return err
}

// Обновляет список по ID
func (t *TodoListPostgres) Update(user_id, list_id int, input todo.UpdateListInput) error {

	// Если запрос изначально пустой - сформируем его самостоятельно

	setValues := make([]string, 0) // Для будущего sql запроса. Может быть таким: title=$1, description=$2
	args := make([]any, 0)         // Переданные значения полей, что будут подставлены вместо $1, $2, ...
	argId := 1                     // Счетчик placeholder'ов

	// Если поля input'a пустые - формируем их сами
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId)) // Тут placeholder, формируемый
		// через автоинкремент и вставку данных через %d
		args = append(args, *input.Title) // Добавляем переданное поле input.Title
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description) // Добавляем переданное поле input.Description
		argId++
	}

	setQuery := strings.Join(setValues, ", ") // Подготовленный нами по необходимости запрос
	// Его варианты:
	// title=$1 илия
	// description=$1 или
	// title=$1, description=$2

	query := fmt.Sprintf(`UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id = $%d 
	AND ul.user_id = $%d`, todoListsTable, setQuery, usersListsTable, argId, argId+1)

	args = append(args, list_id, user_id) // Формируются для корректной подстановки в Exec

	// Кароче епрст. Если отработают оба if'a, то argId == 3, при этом $1 будет занят title'ом, а $2 -
	// дескрипшном (да, я из англии), а при args = append(...) list_id займет $4, а user_id - $5
	// Understandable?

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %v", args)

	_, err := t.db.Exec(query, args...)
	return err
}
