package gorel

import (
	"ast"
)

type SelectManager struct {
	ast Node
	ctx SelectCore
}

func NewSelectManager(table ast.Table) (m SelectManager) {
	m = new(SelectManager)
	m.ast = ast.NewSelectStatement()
	m.ctx = ast.Cores[0]
	return
}

func (m SelectManager) As(alias string) (n Node) {
	return
}
