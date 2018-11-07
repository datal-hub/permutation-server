package database

import (
	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	testData "permutation-server/models/testing"
	"permutation-server/pkg/database/dbtest"
)

// Testing define work mode
// True - for execute test
// False - for real work
var Testing = false

func testDb() (DB, error) {
	db, mock, err := sqlmock.New()

	return &dbtest.DbTest{
		DB:              db,
		Mock:            mock,
		TestPermutation: testData.TestPermutation,
	}, err
}
