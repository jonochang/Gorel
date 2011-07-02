package gorel

import (
	"testing"
	"db"
	"ast"
)

const (
	DB_SOCK   = "/tmp/mysql.sock"
	DB_USER   = "gorel_test"
	DB_PASSWD = "abc123"
	DB_NAME   = "gorel_test"
)


func TestPredicationsEq(t *testing.T) {
	connection, err := db.MySQLNewConnection(DB_SOCK, DB_USER, DB_PASSWD, DB_NAME)

	if err != nil {
		t.Log(err.String())
	}

	table, err := GetTable("Users", connection)

	n := table.Field("id").Eq(&ast.Literal{2})

	v := ast.ToSql{connection}
	s := v.GetEquality(n.(ast.Equality))
	if s != "`Users`.`id` = 2" {
		t.Log(s)
		t.Errorf("Failed to get predication Equality")
	}
}
