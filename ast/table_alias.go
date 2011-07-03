package ast

type TableAlias struct {
	Left  Node
	Table *Table
	Alias *SqlLiteral
}

func (n TableAlias) Visit(v Visitor) (s string) {
	s = v.GetTableAlias(n)
	return
}

func NewTableAlias(table *Table, alias *SqlLiteral) TableAlias {
	return TableAlias{table, table, alias}
}

func NewTableAliasByNode(left Node, table *Table, alias *SqlLiteral) TableAlias {
	return TableAlias{left, table, alias}
}

func (n TableAlias) GetNameAlias() (s string) {
	return n.Alias.Value
}

func (n TableAlias) GetName() (s string) {
	if n.Table != nil {
		s = n.Table.Name
	}
	return
}

func (n TableAlias) HasAlias() bool {
	return true
}

func (n TableAlias) GetField(id string) (field Field) {
	if n.Table != nil {
		field = NewField(n, n.Table.ColumnMap[id])
	} else {
		panic("Can't get field when there is no table") //TODO calculate field from sub select
	}
	return
}
