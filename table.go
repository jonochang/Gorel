package gorel

import (
	"db"
	"ast"
	"fmt"
	"os"
)

type Table struct {
  ast.Table
}

func (t Table) Field(name string) (f ast.Field) {
	column := t.ColumnMap[name]
	quoted_column_name := t.Connection.QuoteColumnName(column.Name)
	quoted_table_name := t.Connection.QuoteTableName(t.Name)
	if t.Alias != "" {
		quoted_table_name = t.Connection.QuoteTableName(t.Alias)
	}

	quoted_name := fmt.Sprintf("%v.%v", quoted_table_name, quoted_column_name)
	f.Name = quoted_name
	return
}

func GetTable(name string, c db.Connection) (t Table, err os.Error) {
	t.TableSchema, err = c.GetTable(name)
	return
}
