package user

import (
	"api-go/internal/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
)

type (
	Repository interface {
		Create(ctx context.Context, user *domain.User) error
		GetAll(ctx context.Context) ([]domain.User, error)
		GetUser(ctx context.Context, id uint64) (*domain.User, error)
		UpdateUser(ctx context.Context, id uint64, firstName, lastName, email *string) error
	}

	repo struct {
		db  *sql.DB
		log *log.Logger
	}
)

func NewRepo(db *sql.DB, l *log.Logger) Repository {
	return &repo{
		db:  db,
		log: l,
	}
}

func (r *repo) Create(ctx context.Context, user *domain.User) error {

	sqlQ := "INSERT INTO users(first_name,last_name, email) VALUES(?,?,?)"
	res, err := r.db.Exec(sqlQ, user.FirstName, user.LastName, user.Email)
	if err != nil {
		r.log.Println(err)
		return err
	}

	id, err := res.LastInsertId()

	if err != nil {
		r.log.Println(err)
		return err
	}

	user.ID = uint64(id)

	log.Println("user created with id: ", id)

	return nil
}

func (r *repo) GetAll(ctx context.Context) ([]domain.User, error) {

	var users []domain.User
	sqlQ := "SELECT * FROM users"

	rows, err := r.db.Query(sqlQ)

	if err != nil {
		r.log.Println(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email); err != nil {
			r.log.Println(err)
			return nil, err
		}
		users = append(users, u)
	}

	r.log.Println("user get all: ", len(users))

	return users, nil
}

func (r *repo) GetUser(ctx context.Context, id uint64) (*domain.User, error) {
	sqlQ := "SELECT * FROM users WHERE id=?"
	var u domain.User
	err := r.db.QueryRow(sqlQ, id).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *repo) UpdateUser(ctx context.Context, id uint64, firstName, lastName, email *string) error {

	var fields []string
	var values []interface{}

	if firstName != nil {
		fields = append(fields, "first_name = ?")
		values = append(values, *firstName)
	}

	if lastName != nil {
		fields = append(fields, "last_name = ?")
		values = append(values, *lastName)
	}

	if email != nil {
		fields = append(fields, "email = ?")
		values = append(values, *email)
	}

	values = append(values, id)

	sqlQ := fmt.Sprintf("UPDATE users SET %s WHERE id=?", strings.Join(fields, ","))

	res, err := r.db.Exec(sqlQ, values...)

	if err != nil {
		return err
	}

	row, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if row == 0 {
		return errors.New("not found")
	}

	return nil
}
