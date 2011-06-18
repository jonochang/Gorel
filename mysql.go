package gorel

import (
	"strconv"
	//"fmt"
)

type MySQL struct {
	ToSql
}

func (c MySQL) GetLiteral(n Literal) (s string) {
	switch val := n.value.(type) {
	case string:
		//s = fmt.Sprintf("'%v'", val)
		//TODO: use custom MySQL Quoting or Binary Protocol
		s = strconv.Quote(val)
	case bool:
		s = strconv.Btoa(val)
	}

	return
}
