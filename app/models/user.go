package models

import (
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
)

type User struct {
	ID int
	UUID string
	Name string
	Email string
	Password string
	CreatedAt time.Time
}

//新規作成用
func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values ($1, $2, $3, $4, $5)`

	_, err = Db.Exec(cmd, createUUID(), u.Name , u.Email, Encrypt(u.Password),time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err 
}

//ユーザー名で該当するデータを取ってくる用
func GetUser(username string) ([]User, error)  {
	var users []User
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	sql, args, err := psql.Select("*").From("users").Where(sq.Eq{"name":username}).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := Db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

//データをすべて取得するよう
func GetAllUser() ([]User, error)  {
	var users []User
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	sql, args, err := psql.Select("*").From("users").ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := Db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

//該当するIdのデータを更新するよう
func (u *User) UpdateUser(id int, updates map[string]interface{}) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	
	update := psql.Update("users")
	for column, value := range updates {
		update = update.Set(column, value)
	}
	
	sql, args, err := update.Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		log.Fatalln(err)
	}

	_, err = Db.Exec(sql, args...)
	if err != nil {
		log.Fatalln(err)
	}
}

//該当するIdのデータを削除するよう
func (u *User) DeleteUser(id int) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	
	sql, args, err := psql.Delete("users").Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		log.Fatalln(err)
	}

	_, err = Db.Exec(sql, args...)
	if err != nil {
		log.Fatalln(err)
	}
}