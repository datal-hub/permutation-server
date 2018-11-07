package database

import (
	"database/sql"
	"fmt"

	"gopkg.in/reform.v1"

	"permutation-server/models"
	"permutation-server/pkg/database/pgsql"
	"permutation-server/pkg/settings"
)

// DB define interface for work with database
type DB interface {
	IsEmpty() bool
	Clear()
	Init(force bool) error
	Close()
	SqlDB() *sql.DB

	Save(record reform.Record) error
	Update(record models.Permutation, rvsIdx int) error
	Find(uid string) (*models.Permutation, error)
}

func NewDB() (DB, error) {
	if Testing == true {
		return testDb()
	}
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d password=%s user=%s dbname=%s sslmode=disable",
		settings.DB.Host, settings.DB.Port, settings.DB.Password, settings.DB.User, settings.DB.Database))
	return &pgsql.PgSQL{DB: db}, err
}
