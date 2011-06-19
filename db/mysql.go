package db

import (
	"reflect"
	"mysql"
	"os"
	"fmt"
)

type MySQL struct {
	Database string
	db       *mysql.Client
}

func MySQLNewConnection(sock string, user string, password string, database string) (Connection, os.Error) {
	var err os.Error
	var db *mysql.Client
	db, err = mysql.DialUnix(sock, user, password, database)
	if err != nil {
		err = db.Close()
		return nil, err
	}

	conn := new(MySQL)
	conn.Database = database
	conn.db = db
	return conn, err
}

func (c MySQL) GetQueryResult(query string) (res *mysql.Result, err os.Error) {

	err = c.db.Query(query)

	if err != nil {
		return
	}

	res, err = c.db.UseResult()

	return res, err
}

func (c MySQL) GetTable(name string) (t TableSchema, err os.Error) {
	var query = "SELECT * FROM " + name + " LIMIT 1;"
	var res *mysql.Result
	res, err = c.GetQueryResult(query)

	if err != nil {
		return
	}

	fields := res.FetchFields()
	t.Database = c.Database
	t.Name = name
	t.ColumnMap = make(map[string]ColumnSchema)
	t.Connection = c

	for i := 0; i < len(fields); i++ {
		field := fields[i]
		var tipe string
		var kind reflect.Kind

		switch field.Type {
		// Signed/unsigned ints
		case mysql.FIELD_TYPE_TINY,
			mysql.FIELD_TYPE_SHORT,
			mysql.FIELD_TYPE_YEAR,
			mysql.FIELD_TYPE_INT24,
			mysql.FIELD_TYPE_LONG,
			mysql.FIELD_TYPE_LONGLONG:
			if field.Flags&mysql.FLAG_UNSIGNED > 0 {
				tipe = "uint"
				kind = reflect.Uint
			} else {
				tipe = "int"
				kind = reflect.Int
			}

		// Floats and doubles
		case mysql.FIELD_TYPE_FLOAT,
			mysql.FIELD_TYPE_DOUBLE:
			tipe = "double"
			kind = reflect.Float32

		// String
		case mysql.FIELD_TYPE_VARCHAR,
			mysql.FIELD_TYPE_VAR_STRING,
			mysql.FIELD_TYPE_STRING:
			tipe = "string"
			kind = reflect.String

		// Decimal
		case mysql.FIELD_TYPE_DECIMAL,
			mysql.FIELD_TYPE_NEWDECIMAL:
			tipe = "decimal"
			kind = reflect.Float32

		// Anything else
		default:
			tipe = "Unknown"
			kind = reflect.Invalid
		}

		col := new(ColumnSchema)
		col.Database = t.Database
		col.Relation = t.Name
		col.Name = field.Name
		col.Type = tipe
		col.GoType = kind
		col.Limit = int(field.Length)
		t.ColumnMap[col.Name] = *col
	}

	err = res.Free()

	if err != nil {
		return
	}

	return
}

func (c MySQL) QuoteColumnName(name string) (s string) {
	s = fmt.Sprintf("`%v`", name)
	return
}

func (c MySQL) QuoteTableName(name string) (s string) {
	s = fmt.Sprintf("`%v`", name)
	return
}
