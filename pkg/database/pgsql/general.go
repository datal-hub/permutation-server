package pgsql

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"permutation-server/pkg/settings"
)

func (db *PgSQL) IsEmpty() bool {
	rows, err := db.DB.Query(`
		SELECT EXISTS (
			SELECT 1 FROM pg_tables
			WHERE  tablename IN ('users')
		);`)
	if err != nil {
		panic("IsEmpty: check exist tables error.")
	}
	defer rows.Close()

	if !rows.Next() {
		panic("IsEmpty: error fetching exists query result.")
	}

	var notEmpty bool
	if err := rows.Scan(&notEmpty); err != nil {
		panic("IsEmpty: scan not empty error.")
	}
	return !notEmpty
}

func (db *PgSQL) Clear() {
	if _, err := db.DB.Exec("DROP TABLE IF EXISTS users"); err != nil {
		panic("Clear: drop function error.")
	}
}

func (db *PgSQL) createTables() error {
	createSQL := `
		CREATE TABLE users (
			uuid uuid PRIMARY KEY,
			login VARCHAR(256),
			dttm timestamp
        );
        CREATE INDEX users_uuid_idx ON users(uuid);`

	if _, err := db.DB.Exec(createSQL); err != nil {
		return err
	}
	return nil
}

func (db *PgSQL) Init(force bool) error {
	if !db.IsEmpty() {
		if force {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print(
				"Database is not empty. All data will be lost.\n",
				"Please confirm database initialization, type database name: ")
			dbname, _ := reader.ReadString('\n')
			if dbname[:len(dbname)-1] != settings.DB.Database {
				return errors.New("database force initializing confirmation failed")
			}
		} else {
			return errors.New("initialization failed - the database is not empty")
		}
		db.Clear()
	}
	err := db.createTables()
	if err != nil {
		return err
	}
	return nil
}
