package ast

import (
	"db"
)

type Field struct {
	Table  TableNameAlias
	Column db.ColumnSchema
	*ExpressionFunctions
}

func (n Field) Visit(v Visitor) (s string) {
	s = v.GetField(n)
	return
}

func NewField(table TableNameAlias, column db.ColumnSchema) Field {
	return Field{table, column, &ExpressionFunctions{}}
}

func (f Field) Count() CountNode { return f.count(f, false) }

func (f Field) CountDistinct() CountNode { return f.count(f, true) }

func (f Field) Sum() SumNode { return f.sum(f) }

func (f Field) Max() MaxNode { return f.max(f) }

func (f Field) Min() MinNode { return f.min(f) }

func (f Field) Avg() AvgNode { return f.avg(f) }

func (f Field) Ascending() Ascending { return Ascending{Unary{f}} }

func (f Field) Descending() Descending { return Descending{Unary{f}} }
