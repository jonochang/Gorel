package ast

import (
	"db"
)

type Table struct {
	db.TableSchema
  //aliases []TableAlias
}

func (n Table) Visit(v Visitor) (s string) {
	s = v.GetTable(n)
	return
}

func (n Table) GetNameAlias() (s string) {
  return n.Alias
}

func (n Table) GetName() (s string) {
  return n.Name
}

//func (n Table) GetTableAlias() (ta TableAlias) {
//  length := len(n.Aliases)
//  ta.Left = n
//  ta.Right = fmt.Sprintf("%v_%v", n.Name, length+1)
//  ta.Left.Alias = ta.Right //set alias for fields
//  return ta
//}
