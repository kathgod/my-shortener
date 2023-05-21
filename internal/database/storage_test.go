package storage_test

import (
	"testing"
	st "urlshortener/internal/database"
)

func TestOpenDB(t *testing.T) {
	dsn := "postgres://user:password@localhost:5432/dbname?sslmode=disable"
	db, err := st.OpenDB(dsn)
	if (db == nil) || (err != nil) {
		t.Error("expected non-nil db, got nil")
	}
	defer func() {
		err := db.Close()
		if err != nil {
			t.Error(err)
		}
	}()
}
