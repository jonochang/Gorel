package ast

import (
	"db"
)

type Field struct {
	Table  Table
	Column db.ColumnSchema
}

func (n Field) Visit(v Visitor) (s string) {
	s = v.GetField(n)
	return
}
