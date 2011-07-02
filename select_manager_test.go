package gorel

import (
	"testing"
	"db"
	"ast"
)

func TestSelectManagerAs(t *testing.T) {
	connection, err := db.MySQLNewConnection(DB_SOCK, DB_USER, DB_PASSWD, DB_NAME)

	if err != nil {
		t.Log(err.String())
	}

	table, err := GetTable("Users", connection)
	m1 := NewSelectManagerFromTable(-1, connection, table.Table)
	n := m1.Project(ast.SqlLiteral{"*"}).As("t")

	m2 := NewSelectManager(-1, connection)
	m2.From(n).Project("*")
	s := m2.ToSql()
	if s != "SELECT * FROM (SELECT * FROM `Users` ) `t` " {
		t.Log(s)
		t.Errorf("Failed to get As for select manager")
	}
}

func TestSelectManagerJoin(t *testing.T) {
	connection, err := db.MySQLNewConnection(DB_SOCK, DB_USER, DB_PASSWD, DB_NAME)

	if err != nil {
		t.Log(err.String())
	}

	table, err := GetTable("Users", connection)
	m1 := NewSelectManagerFromTable(-1, connection, table.Table)
	t2 := table.Alias()
	m1.Join(t2).Project("*")

	s := m1.ToSql()
	if s != "SELECT * FROM `Users` INNER JOIN `Users` `Users_1` " {
		t.Log(s)
		t.Errorf("Failed to get Join for select manager")
	}
}

func TestSelectManagerJoinOn(t *testing.T) {
	connection, err := db.MySQLNewConnection(DB_SOCK, DB_USER, DB_PASSWD, DB_NAME)

	if err != nil {
		t.Log(err.String())
	}

	table, err := GetTable("Users", connection)
	m1 := NewSelectManagerFromTable(-1, connection, table.Table)
	t2 := table.Alias()

	predicate := table.Field("id").Eq(t2.Field("id"))
	m1.Join(t2).On(predicate).Project("*")
	s := m1.ToSql()
	if s != "SELECT * FROM `Users` INNER JOIN `Users` `Users_1` ON `Users`.`id` = `Users_1`.`id`" {
		t.Log(s)
		t.Errorf("Failed to get Join On for select manager")
	}
}

func TestSelectManagerOrder(t *testing.T) {
	connection, err := db.MySQLNewConnection(DB_SOCK, DB_USER, DB_PASSWD, DB_NAME)

	if err != nil {
		t.Log(err.String())
	}

	table, err := GetTable("Users", connection)
	m1 := NewSelectManagerFromTable(-1, connection, table.Table)
	m2 := m1
	m1.Order(table.Field("id"))
	s := m1.ToSql()
	if s != "SELECT FROM `Users`  ORDER BY `Users`.`id`" {
		t.Log(s)
		t.Errorf("Failed to get Order On for select manager")
	}

	m2.Order([]ast.Node{table.Field("id"), ast.SqlLiteral{"login"}}...)
	s = m2.ToSql()
	if s != "SELECT FROM `Users`  ORDER BY `Users`.`id`, login" {
		t.Log(s)
		t.Errorf("Failed to get Order For for select manager")
	}
}

func TestSelectManagerOffset(t *testing.T) {
	connection, err := db.MySQLNewConnection(DB_SOCK, DB_USER, DB_PASSWD, DB_NAME)

	if err != nil {
		t.Log(err.String())
	}

	table, err := GetTable("Users", connection)
	m1 := NewSelectManagerFromTable(-1, connection, table.Table)
	m1.Offset(12)
	s := m1.ToSql()
	if s != "SELECT FROM `Users`  OFFSET 12" {
		t.Log(s)
		t.Errorf("Failed to get Offset for select manager")
	}

	m1.RemoveOffset()
	s = m1.ToSql()
	if s != "SELECT FROM `Users` " {
		t.Log(s)
		t.Errorf("Failed to get Offset for select manager")
	}
}

func TestSelectManagerLimit(t *testing.T) {
	connection, err := db.MySQLNewConnection(DB_SOCK, DB_USER, DB_PASSWD, DB_NAME)

	if err != nil {
		t.Log(err.String())
	}

	table, err := GetTable("Users", connection)
	m1 := NewSelectManagerFromTable(-1, connection, table.Table)
	m1.Limit(7)
	s := m1.ToSql()
	if s != "SELECT FROM `Users`  LIMIT 7" {
		t.Log(s)
		t.Errorf("Failed to get Offset for select manager")
	}

	m1.RemoveLimit()
	s = m1.ToSql()
	if s != "SELECT FROM `Users` " {
		t.Log(s)
		t.Errorf("Failed to get Offset for select manager")
	}
}
