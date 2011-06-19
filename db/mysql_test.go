package db

import (
	"testing"
)

const (
	DB_SOCK   = "/tmp/mysql.sock"
	DB_USER   = "gorel_test"
	DB_PASSWD = "abc123"
	DB_NAME   = "gorel_test"
)


func TestMySQLGetTable(t *testing.T) {
	connection, err := MySQLNewConnection(DB_SOCK, DB_USER, DB_PASSWD, DB_NAME)
	if err != nil {
		t.Log(err.String())
	}
	table, err := connection.GetTable("Users")

	if len(table.ColumnMap) != 2 {
		t.Log(len(table.ColumnMap))
		t.Errorf("Failed to get Mysql Table Users")
	}
}
