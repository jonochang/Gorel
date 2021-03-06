package ast

type TableNameAlias interface {
	GetNameAlias() string
	GetName() string
	HasAlias() bool
	GetField(string) Field
}
