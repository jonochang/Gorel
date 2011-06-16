package gorel

import (
	"testing"
	"fmt"
	//"strconv"
)

const (
	DB_SOCK   = "/tmp/mysql.sock"
	DB_USER   = "gorel_test"
	DB_PASSWD = "abc123"
	DB_NAME   = "gorel_test"
)

func TestVisit(t *testing.T) {
  binary := new(Equality)
  l := new(Literal)
  l.value = "leftb"
  binary.left = l
  
  
  r := new(Literal)
  r.value = "rightb"
  binary.right = r
  
  p := new(Literal)
  p.value = "p"
  u := new(Unary)
  u.expression = p
  join := new(Join)
  join.left = binary
  join.right = u
  
  v := new(MySQL)
  //v.Visit(*join)
  s := join.Visit(v)
  fmt.Println(s)
}

//func TestTable(t *testing.T) {
//	conn := GetConnection()
//	table := conn.Table("users")
//	if table.name != "users" {
//		t.Errorf("Error wrong table name")
//	}
//
//	if len(table.columns) != 2 {
//		t.Errorf("invalid column number")
//	}
//
//	if table.columns["id"].name != "id" {
//		t.Errorf("invalid id column")
//	}
//
//	if table.columns["username"].name != "username" {
//		t.Errorf("invalid username column")
//	}
//
//	table = table.where_string("id = 1")
//
//	if table.wheres[0].ToSql() != "(id = 1)" {
//		t.Errorf("invalid where")
//	}
//
//	var sql = table.execute()
//
//	if sql != "SELECT * FROM users WHERE (id = 1)" {
//		t.Errorf("invalid sql")
//	}
//
//	table = table.where_string("username = 'testuser'")
//
//	if table.wheres[1].ToSql() != "(username = 'testuser')" {
//		t.Errorf("invalid where")
//	}
//
//	sql = table.execute()
//
//	if sql != "SELECT * FROM users WHERE (id = 1) AND (username = 'testuser')" {
//		t.Errorf("invalid sql")
//	}
//
//}
//
//func TestEquality(t *testing.T) {
//	conn := GetConnection()
//	table := conn.Table("users")
//	table = table.where(table.columns["id"].eq(9))
//	sql := table.execute()
//
//	if sql != "SELECT * FROM users WHERE (`users`.`id` = 9)" {
//		t.Errorf("invalid sql")
//	}
//}
//
//
//func TestGreaterThan(t *testing.T) {
//	conn := GetConnection()
//	table := conn.Table("users")
//	table = table.where(table.columns["id"].gt(9))
//	sql := table.execute()
//
//	if sql != "SELECT * FROM users WHERE (`users`.`id` > 9)" {
//		t.Errorf("invalid sql")
//	}
//}
//
//func TestLessThan(t *testing.T) {
//	conn := GetConnection()
//	table := conn.Table("users")
//	table = table.where(table.columns["id"].lt(9))
//	sql := table.execute()
//
//	if sql != "SELECT * FROM users WHERE (`users`.`id` < 9)" {
//		t.Errorf("invalid sql")
//	}
//}
