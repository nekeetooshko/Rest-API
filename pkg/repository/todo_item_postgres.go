package repository

import (
	todo "MaksJash"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func newTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{
		db: db,
	}
}

func (r *TodoItemPostgres) Create(list_id int, item todo.TodoItem) (int, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var item_id int
	createItemQuery := fmt.Sprintf(
		`INSERT INTO %s (title, description) values ($1, $2) RETURNING id`, todoItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&item_id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListQuery := fmt.Sprintf(`
	insert into %s (list_id, item_id) values ($1, $2)`, listsItemsTable)
	_, err = tx.Exec(createListQuery, list_id, item_id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return item_id, tx.Commit()

}

func (r *TodoItemPostgres) GetAll(user_id, list_id int) ([]todo.TodoItem, error) {

	var items []todo.TodoItem

	query := fmt.Sprintf(`
	select ti.id, ti.title, ti.description, ti.done 
	from %s ti 
	inner join %s li on li.item_id = ti.id 
	inner join %s ul on ul.list_id = li.list_id
	where li.list_id = $1 and ul.user_id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)

	err := r.db.Select(&items, query, list_id, user_id)
	if err != nil {
		return nil, err
	}

	return items, nil

}

func (r *TodoItemPostgres) GetItemById(user_id, item_id int) (todo.TodoItem, error) {

	var item todo.TodoItem

	query := fmt.Sprintf(`
	select ti.id, ti.title, ti.description, ti.done 
	from %s ti 
	inner join %s li on li.item_id = ti.id 
	inner join %s ul on ul.list_id = li.list_id
	where ti.id = $1 and ul.user_id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)

	err := r.db.Get(&item, query, item_id, user_id)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (r *TodoItemPostgres) Delete(user_id, item_id int) error {

	query := fmt.Sprintf(`
	DELETE from %s ti 
	using %s li, %s ul 
	where ti.id = li.item_id
	and li.list_id = ul.list_id
	and ul.user_id = $1
	and ti.id = $2
	`,
		todoItemsTable, listsItemsTable, usersListsTable)

	_, err := r.db.Exec(query, user_id, item_id)
	return err
}

func (r *TodoItemPostgres) Update(user_id, item_id int, input todo.UpdateItemInput) error {

	setValues := make([]string, 0)
	args := make([]any, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))

		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`
	UPDATE %s ti 
	SET %s FROM %s li, %s ul
	WHERE ti.id = li.item_id  
	and li.list_id = ul.list_id
	AND ul.user_id = $%d 
	AND ti.id = $%d`,
		todoItemsTable, setQuery, listsItemsTable, usersListsTable, argId, argId+1)

	args = append(args, user_id, item_id)

	_, err := r.db.Exec(query, args...)
	return err
}
