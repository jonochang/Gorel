package db

type Adapter int

const (
	Adapter_MySQL Adapter = iota
	Adapter_PostgreSQL
)
