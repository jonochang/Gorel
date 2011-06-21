package gorel

import (
	"ast"
	"db"
	"strconv"
)

type SelectManager struct {
	Adapter db.Adapter
	Visitor ast.Visitor
	Ast     ast.SelectStatement
}

func VisitorsFor(adapter db.Adapter, connection db.Connection) (v ast.Visitor) {
	switch adapter {
	case db.Adapter_MySQL:
		v = ast.MySQL{ast.ToSql{connection}}
	case db.Adapter_PostgreSQL:
	default:
		v = ast.ToSql{connection}
	}
	return
}

func NewSelectManager(adapter db.Adapter, connection db.Connection) (m SelectManager) {
	v := VisitorsFor(adapter, connection)
	m.Adapter = adapter
	m.Visitor = v
	m.Ast = ast.NewSelectStatement()
	return
}

func NewSelectManagerFromTable(adapter db.Adapter, connection db.Connection, table ast.Table) (m SelectManager) {
	m = NewSelectManager(adapter, connection)
	(&m).From(table)
	return
}

func (m SelectManager) As(alias string) (n ast.Node) {
	grouping := ast.Grouping{ast.Unary{m.Ast}}
	n = ast.TableAlias{ast.Binary{grouping, ast.SqlLiteral{alias}}}
	return
}

func (m *SelectManager) From(table ast.Node) *SelectManager {
	//TODO get pointer to SelectCore
	c := m.Ast.Cores[len(m.Ast.Cores)-1].(ast.SelectCore)
	js := c.Source.(ast.JoinSource)
	switch val := table.(type) {
	case ast.InnerJoin, ast.OuterJoin, ast.StringJoin:
		r := js.JoinOn
		r = append(r, val)
		js.JoinOn = r
	default:
		js.Source = val
	}
	c.Source = js
	m.Ast.Cores[len(m.Ast.Cores)-1] = c
	return m
}

func (m *SelectManager) Project(any interface{}) *SelectManager {
	switch val := any.(type) {
	case ast.Node:
		m.project(val)
	case string:
		m.project(ast.SqlLiteral{any.(string)})
	case bool:
		m.project(ast.SqlLiteral{strconv.Btoa(any.(bool))})
	case int:
		m.project(ast.SqlLiteral{strconv.Itoa(any.(int))})
	case int64:
		m.project(ast.SqlLiteral{strconv.Itoa64(any.(int64))})
	case float32:
		m.project(ast.SqlLiteral{strconv.Ftoa32(any.(float32), 'f', 0)})
	case float64:
		m.project(ast.SqlLiteral{strconv.Ftoa64(any.(float64), 'f', 0)})
	}
	return m
}

func (m *SelectManager) project(n ast.Node) *SelectManager {
	c := m.Ast.Cores[len(m.Ast.Cores)-1].(ast.SelectCore)
	c.Projections = append(c.Projections, n)
	m.Ast.Cores[len(m.Ast.Cores)-1] = c
	return m
}

func (m SelectManager) ToSql() string {
	return m.Ast.Visit(m.Visitor)
}
