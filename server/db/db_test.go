package db

import "testing"

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:     "localhost",
    Port:       "5432",
		User:       "postgres",
		Password:   "postgres",
		Host:       "localhost",
		DisableSSL: true,
	}
	if conn.ConnectionURL() != "postgres://postgres:password@localhost/restaurant?sslmode=disable" {
		t.Error("Unexpected connection string")
	}
}
