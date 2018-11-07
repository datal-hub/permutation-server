package dbtest

import (
	"database/sql"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"permutation-server/models"
)

type DbTest struct {
	DB              *sql.DB
	Mock            sqlmock.Sqlmock
	TestPermutation models.Permutation
}

func (db *DbTest) SqlDB() *sql.DB {
	return db.DB
}

func (db *DbTest) Close() {
	db.DB.Close()
}
