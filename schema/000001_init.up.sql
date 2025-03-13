CREATE TABLE users 
(
    id serial not null unique , -- вроде как лучше писать id serial primary key
    name varchar(255) not null, 
    username varchar(255) not null unique, 
    password_hash varchar(255) not null
);

CREATE TABLE todo_lists 
(
    id serial not null unique , -- serial - тот же int, но с автоинкрементом
    title varchar(255) not null, 
    description varchar(255) 
);

CREATE TABLE users_lists
(
    id serial not null unique , 
    user_id int references users (id) on delete cascade not null,  -- вроде лучше писать так: FOREIGN KEY (user_id) REFERENCES users ON DELETE CASCADE
    list_id int references todo_lists (id) on delete cascade not null
    /* 

    references todo_lists (id)
    Это говорит, что list_id является внешним ключом, что ссылается 
    на столбец id в todo_lists. Таким образом 2 таблицы связываются между собой
    Т.е. list_id должно соответствовать todo_lists.id

    on delete cascade
    Говорит, что если todo_lists.id будет удалена - все связанные записи из текущей таблицы тоже удаляться 
    */
);

CREATE TABLE todo_items 
(
    id serial not null unique , 
    title varchar(255) not null, 
    description varchar(255), 
    done boolean not null default false -- default - значение по умолчанию
);

CREATE TABLE lists_items
(
    id serial not null unique , 
    item_id int references todo_items (id) on delete cascade not null, 
    list_id int references todo_lists (id) on delete cascade not null
)