package controllers_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func beforeEach(t *testing.T) (sqlmock.Sqlmock, *sql.DB) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error creating mock database connection:", err)
	}

	return mock, db
}
