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
		mysql := ast.MySQL{ast.ToSql{connection}}
		v = &mysql
	case db.Adapter_PostgreSQL:
	default:
		tosql := ast.ToSql{connection}
		v = &tosql
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
	l := ast.NewSqlLiteral(alias)
	n = ast.NewTableAliasByNode(grouping, nil, &l)
	return
}

func (m *SelectManager) From(table ast.Node) *SelectManager {
	c := m.Ast.Cores[len(m.Ast.Cores)-1].(*ast.SelectCore)
	js := c.Source.(*ast.JoinSource)
	switch val := table.(type) {
	case ast.InnerJoin, ast.OuterJoin, ast.StringJoin:
		r := js.JoinOn
		r = append(r, val)
		js.JoinOn = r
	default:
		js.Source = val
	}

	return m
}

func (m *SelectManager) Join(table ast.Node) *SelectManager {
	return m.InnerJoin(table)
}

func (m *SelectManager) InnerJoin(table ast.Node) *SelectManager {
	c := m.Ast.Cores[len(m.Ast.Cores)-1].(*ast.SelectCore)
	js := c.Source.(*ast.JoinSource)

	r := js.JoinOn
	r = append(r, ast.InnerJoin{&ast.BaseJoin{table, nil}})
	js.JoinOn = r

	return m
}

func (m *SelectManager) OuterJoin(table ast.Node) *SelectManager {
	c := m.Ast.Cores[len(m.Ast.Cores)-1].(*ast.SelectCore)
	js := c.Source.(*ast.JoinSource)

	r := js.JoinOn
	r = append(r, ast.OuterJoin{&ast.BaseJoin{table, nil}})
	js.JoinOn = r

	return m
}

func (m *SelectManager) Project(any interface{}) *SelectManager {
	switch val := any.(type) {
	case ast.Node:
		m.project(val)
	case string:
		m.project(ast.NewSqlLiteral(any.(string)))
	case bool:
		m.project(ast.NewSqlLiteral(strconv.Btoa(any.(bool))))
	case int:
		m.project(ast.NewSqlLiteral(strconv.Itoa(any.(int))))
	case int64:
		m.project(ast.NewSqlLiteral(strconv.Itoa64(any.(int64))))
	case float32:
		m.project(ast.NewSqlLiteral(strconv.Ftoa32(any.(float32), 'f', 0)))
	case float64:
		m.project(ast.NewSqlLiteral(strconv.Ftoa64(any.(float64), 'f', 0)))
	}
	return m
}

func (m *SelectManager) project(n ast.Node) *SelectManager {
	c := m.Ast.Cores[len(m.Ast.Cores)-1].(*ast.SelectCore)
	c.Projections = append(c.Projections, n)
	m.Ast.Cores[len(m.Ast.Cores)-1] = c
	return m
}

func (m *SelectManager) On(n ast.Node) *SelectManager {
	c := m.Ast.Cores[len(m.Ast.Cores)-1].(*ast.SelectCore)
	js := c.Source.(*ast.JoinSource)

	last := js.JoinOn[len(js.JoinOn)-1]
	last_join := last.(ast.Join)
	on := ast.On{ast.Unary{n}}
	last_join.SetRight(on)

	js.JoinOn[len(js.JoinOn)-1] = last_join

	return m
}

func (m *SelectManager) Order(n ...ast.Node) *SelectManager {
	m.Ast.Orders = append(m.Ast.Orders, n...)
	return m
}

func (m *SelectManager) Group(n ...ast.Node) *SelectManager {
	c := m.Ast.Cores[len(m.Ast.Cores)-1].(*ast.SelectCore)
	groups := make([]ast.Node, 0)

	for i := 0; i < len(n); i++ {
		field := n[i]
		group := ast.Group{ast.Unary{field}}
		groups = append(groups, group)
	}

	c.Groups = append(c.Groups, groups...)
	return m
}

func (m *SelectManager) Offset(amount int) *SelectManager {
	n := new(ast.Offset)
	n.Expression = &ast.Literal{amount}
	m.Ast.Offset = n
	return m
}

func (m *SelectManager) RemoveOffset() *SelectManager {
	m.Ast.Offset = nil
	return m
}

func (m *SelectManager) Limit(amount int) *SelectManager {
	n := new(ast.Limit)
	n.Expression = &ast.Literal{amount}
	m.Ast.Limit = n
	return m
}

func (m *SelectManager) RemoveLimit() *SelectManager {
	m.Ast.Limit = nil
	return m
}

func (m SelectManager) ToSql() string {
	return m.Ast.Visit(m.Visitor)
}
