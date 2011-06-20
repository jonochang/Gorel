package gorel

import (
	"db"
	"ast"
	"os"
)

type Table struct {
	ast.Table
}

func (t Table) Field(name string) (f ast.Field) {
	column := t.ColumnMap[name]
	f.Table = t.Table
	f.Column = column
	return
}

func GetTable(name string, c db.Connection) (t Table, err os.Error) {
	t.TableSchema, err = c.GetTable(name)
	return
}
