package gorel

import (
	"testing"
)

func TestMySQL_GetLiteral(t *testing.T) {
	v := new(MySQL)
	l := new(Literal)

	t.Log("Test string")
	l.value = "te\"st"
	s := v.GetLiteral(*l)
	if s != "\"te\\\"st\"" {
		t.Log(s)
    t.Errorf("failed test string")
	}

}
