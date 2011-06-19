package ast

import (
  "db"
)

type Table struct {
  db.TableSchema
}

func (n Table) Visit(v Visitor) (s string) {
  s = v.GetTable(n)
  return
}

