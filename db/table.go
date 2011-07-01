package db

import (
	"reflect"
)

type ColumnSchema struct {
	Database  string
	Relation  string
	Name      string
	Type      string
	GoType    reflect.Kind
	Limit     int
	Precision int
	Scale     int
	Null      bool
	Default   interface{}
}

type TableSchema struct {
	Database   string
	Name       string
	ColumnMap  map[string]ColumnSchema
	Connection Connection
}
