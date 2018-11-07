package pgsql

import "database/sql"

type PgSQL struct {
	DB *sql.DB
}

func (db *PgSQL) SqlDB() *sql.DB {
	return db.DB
}

func (db *PgSQL) Close() {
	db.DB.Close()
}
