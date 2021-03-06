package ast

import (
	"db"
	"fmt"
)

type Table struct {
	db.TableSchema
	Aliases []Node
}

func (n Table) Visit(v Visitor) (s string) {
	s = v.GetTable(n)
	return
}

func (n Table) GetNameAlias() (s string) {
	length := len(n.Aliases)
	s = fmt.Sprintf("%v_%v", n.Name, length)
	return
}

func (n Table) GetName() string {
	return n.Name
}

func (n Table) HasAlias() bool {
	return len(n.Aliases) > 0
}

func (n *Table) CreateTableAlias() (ta TableAlias) {
	l := NewSqlLiteral(n.GetNameAlias())
	ta = NewTableAlias(n, &l)
	n.Aliases = append(n.Aliases, ta)
	return ta
}

func (n Table) CurrentTableAlias() TableAlias {
	return n.Aliases[len(n.Aliases)-1].(TableAlias)
}

func (n Table) GetField(id string) Field {
	return NewField(n, n.ColumnMap[id])
}
