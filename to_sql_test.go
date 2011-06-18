package gorel

import (
	"testing"
)

func TestGetLiteral(t *testing.T) {
	v := new(ToSql)
	l := new(Literal)

	t.Log("Test string")
	l.value = "test"
	s := v.GetLiteral(*l)
	if s != "test" {
		t.Errorf("failed test string")
	}

	t.Log("Test boolean")
	l.value = false
	s = v.GetLiteral(*l)
	if s != "false" {
		t.Errorf("failed test boolean")
	}

	t.Log("Test int")
	l.value = 123
	s = v.GetLiteral(*l)
	if s != "123" {
		t.Errorf("failed test int")
	}

	t.Log("Test float")
	l.value = 123.3
	s = v.GetLiteral(*l)
	if s != "123.3" {
		t.Errorf("failed test float")
	}
}

func TestVisitNilNodes(t *testing.T) {
	v := new(ToSql)
	s := v.VisitNodes(nil)

	if len(s) != 0 {
		t.Errorf("failed to visit nil nodes")
	}

	nodes := make([]Node, 3)
	s = v.VisitNodes(nodes)
	if len(s) != 0 {
		t.Errorf("failed to skip nil nodes")
	}
}

//-----------------And----------------
func TestGetAnd(t *testing.T) {
	v := new(ToSql)
	n := new(And)

	n1 := Literal{1}
	n2 := Literal{2}
	n3 := Literal{3}
	n.children = []Node{n1, n2, n3}

	s := v.GetAnd(*n)
	if s != "1 AND 2 AND 3" {
		t.Errorf("failed to get And")
	}
}

//-----------------Binary----------------
func TestGetBetween(t *testing.T) {
	v := new(ToSql)
	n := new(Between)

	s := v.GetBetween(*n)
	if s != "" {
		t.Errorf("failed to get Between ")
	}
}
func TestGetNotEqual(t *testing.T) {
	v := new(ToSql)
	n := new(NotEqual)

	s := v.GetNotEqual(*n)
	if s != "" {
		t.Errorf("failed to get NotEqual ")
	}
}
func TestGetAssignment(t *testing.T) {
	v := new(ToSql)
	n := new(Assignment)

	s := v.GetAssignment(*n)
	if s != "" {
		t.Errorf("failed to get Assignment ")
	}
}
func TestGetOr(t *testing.T) {
	v := new(ToSql)
	n := new(Or)

	s := v.GetOr(*n)
	if s != "" {
		t.Errorf("failed to get Or ")
	}
}
func TestGetAs(t *testing.T) {
	v := new(ToSql)
	n := new(As)

	s := v.GetAs(*n)
	if s != "" {
		t.Errorf("failed to get As ")
	}
}
func TestGetGreaterThan(t *testing.T) {
	v := new(ToSql)
	n := new(GreaterThan)

	s := v.GetGreaterThan(*n)
	if s != "" {
		t.Errorf("failed to get GreaterThan ")
	}
}
func TestGetGreaterThanOrEqual(t *testing.T) {
	v := new(ToSql)
	n := new(GreaterThanOrEqual)

	s := v.GetGreaterThanOrEqual(*n)
	if s != "" {
		t.Errorf("failed to get GreaterThanOrEqual ")
	}
}
func TestGetLessThan(t *testing.T) {
	v := new(ToSql)
	n := new(LessThan)

	s := v.GetLessThan(*n)
	if s != "" {
		t.Errorf("failed to get LessThan ")
	}
}
func TestGetLessThanOrEqual(t *testing.T) {
	v := new(ToSql)
	n := new(LessThanOrEqual)

	s := v.GetLessThanOrEqual(*n)
	if s != "" {
		t.Errorf("failed to get LessThanOrEqual ")
	}
}
func TestGetMatches(t *testing.T) {
	v := new(ToSql)
	n := new(Matches)

	s := v.GetMatches(*n)
	if s != "" {
		t.Errorf("failed to get Matches ")
	}
}
func TestGetDoesNotMatch(t *testing.T) {
	v := new(ToSql)
	n := new(DoesNotMatch)

	s := v.GetDoesNotMatch(*n)
	if s != "" {
		t.Errorf("failed to get DoesNotMatch ")
	}
}
func TestGetNotIn(t *testing.T) {
	v := new(ToSql)
	n := new(NotIn)

	s := v.GetNotIn(*n)
	if s != "" {
		t.Errorf("failed to get NotIn ")
	}
}
func TestGetOrdering(t *testing.T) {
	v := new(ToSql)
	n := new(Ordering)

	s := v.GetOrdering(*n)
	if s != "" {
		t.Errorf("failed to get Ordering ")
	}
}
func TestGetValues(t *testing.T) {
	v := new(ToSql)
	n := new(Values)

	s := v.GetValues(*n)
	if s != "" {
		t.Errorf("failed to get Values ")
	}
}
func TestGetDeleteStatement(t *testing.T) {
	v := new(ToSql)
	n := new(DeleteStatement)

	s := v.GetDeleteStatement(*n)
	if s != "" {
		t.Errorf("failed to get DeleteStatement ")
	}
}
func TestGetTableAlias(t *testing.T) {
	v := new(ToSql)
	n := new(TableAlias)

	s := v.GetTableAlias(*n)
	if s != "" {
		t.Errorf("failed to get TableAlias ")
	}
}
func TestGetExcept(t *testing.T) {
	v := new(ToSql)
	n := new(Except)

	s := v.GetExcept(*n)
	if s != "" {
		t.Errorf("failed to get Except ")
	}
}
func TestGetIntersect(t *testing.T) {
	v := new(ToSql)
	n := new(Intersect)

	s := v.GetIntersect(*n)
	if s != "" {
		t.Errorf("failed to get Intersect ")
	}
}
func TestGetUnion(t *testing.T) {
	v := new(ToSql)
	n := new(Union)

	s := v.GetUnion(*n)
	if s != "" {
		t.Errorf("failed to get Union ")
	}
}
func TestGetUnionAll(t *testing.T) {
	v := new(ToSql)
	n := new(UnionAll)

	s := v.GetUnionAll(*n)
	if s != "" {
		t.Errorf("failed to get UnionAll ")
	}
}

//-----------------Equality----------------
func TestGetEquality(t *testing.T) {
	v := new(ToSql)
	n := new(Equality)

	ll := new(Literal)
	ll.value = 1
	n.left = ll

	rl := new(Literal)
	rl.value = 2
	n.right = rl

	s := v.GetEquality(*n)
	if s != "1 = 2" {
		t.Errorf("failed to get Equality")
	}
}

func TestGetIn(t *testing.T) {
	v := new(ToSql)
	n := new(In)

	s := v.GetIn(*n)
	if s != "" {
		t.Errorf("failed to get In ")
	}
}

//-----------------Function----------------
func TestGetCount(t *testing.T) {
	v := new(ToSql)
	n := new(Count)

	s := v.GetCount(*n)
	if s != "" {
		t.Errorf("failed to get Count ")
	}
}
func TestGetSum(t *testing.T) {
	v := new(ToSql)
	n := new(Sum)

	s := v.GetSum(*n)
	if s != "" {
		t.Errorf("failed to get Sum ")
	}
}
func TestGetExists(t *testing.T) {
	v := new(ToSql)
	n := new(Exists)

	s := v.GetExists(*n)
	if s != "" {
		t.Errorf("failed to get Exists ")
	}
}
func TestGetMax(t *testing.T) {
	v := new(ToSql)
	n := new(Max)

	s := v.GetMax(*n)
	if s != "" {
		t.Errorf("failed to get Max ")
	}
}
func TestGetMin(t *testing.T) {
	v := new(ToSql)
	n := new(Min)

	s := v.GetMin(*n)
	if s != "" {
		t.Errorf("failed to get Min ")
	}
}
func TestGetAvg(t *testing.T) {
	v := new(ToSql)
	n := new(Avg)

	s := v.GetAvg(*n)
	if s != "" {
		t.Errorf("failed to get Avg ")
	}
}

//-----------------Join----------------
func TestGetInnerJoin(t *testing.T) {
	v := new(ToSql)
	n := new(InnerJoin)

	s := v.GetInnerJoin(*n)
	if s != "" {
		t.Errorf("failed to get InnerJoin ")
	}
}
func TestGetOuterJoin(t *testing.T) {
	v := new(ToSql)
	n := new(OuterJoin)

	s := v.GetOuterJoin(*n)
	if s != "" {
		t.Errorf("failed to get OuterJoin ")
	}
}
func TestGetStringJoin(t *testing.T) {
	v := new(ToSql)
	n := new(StringJoin)

	s := v.GetStringJoin(*n)
	if s != "" {
		t.Errorf("failed to get StringJoin ")
	}
}

//-----------------Unary----------------
func TestGetNot(t *testing.T) {
	v := new(ToSql)
	n := new(Not)

	s := v.GetNot(*n)
	if s != "" {
		t.Errorf("failed to get Not ")
	}
}
func TestGetLock(t *testing.T) {
	v := new(ToSql)
	n := new(Lock)

	s := v.GetLock(*n)
	if s != "" {
		t.Errorf("failed to get Lock ")
	}
}
func TestGetOffset(t *testing.T) {
	v := new(ToSql)
	n := new(Offset)

	s := v.GetOffset(*n)
	if s != "" {
		t.Errorf("failed to get Offset ")
	}
}
func TestGetLimit(t *testing.T) {
	v := new(ToSql)
	n := new(Limit)

	s := v.GetLimit(*n)
	if s != "" {
		t.Errorf("failed to get Limit ")
	}
}
func TestGetTop(t *testing.T) {
	v := new(ToSql)
	n := new(Top)

	s := v.GetTop(*n)
	if s != "" {
		t.Errorf("failed to get Top ")
	}
}
func TestGetHaving(t *testing.T) {
	v := new(ToSql)
	n := new(Having)

	s := v.GetHaving(*n)
	if s != "" {
		t.Errorf("failed to get Having ")
	}
}
func TestGetUnqualifiedColumn(t *testing.T) {
	v := new(ToSql)
	n := new(UnqualifiedColumn)

	s := v.GetUnqualifiedColumn(*n)
	if s != "" {
		t.Errorf("failed to get UnqualifiedColumn ")
	}
}
func TestGetGroup(t *testing.T) {
	v := new(ToSql)
	n := new(Group)

	s := v.GetGroup(*n)
	if s != "" {
		t.Errorf("failed to get Group ")
	}
}
func TestGetGrouping(t *testing.T) {
	v := new(ToSql)
	n := new(Grouping)

	s := v.GetGrouping(*n)
	if s != "" {
		t.Errorf("failed to get Grouping ")
	}
}
func TestGetOn(t *testing.T) {
	v := new(ToSql)
	n := new(On)

	s := v.GetOn(*n)
	if s != "" {
		t.Errorf("failed to get On ")
	}
}
