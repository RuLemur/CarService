package repo

import (
	"car-service/internal/app/datastruct"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
}



func AddNewUser(db *sqlx.DB, user *datastruct.User) error {
	err := db.QueryRowx(`INSERT INTO users (username) VALUES ($1) RETURNING id`, user.Username).Scan(&user.ID)
	return err
}
