package db

import (
	"os"
)

type Connection interface {
	GetTable(name string) (t TableSchema, err os.Error)
	QuoteColumnName(name string) (s string)
	QuoteTableName(name string) (s string)
}
