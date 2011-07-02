package ast

import (
	"testing"
	"db"
)

const (
	DB_SOCK   = "/tmp/mysql.sock"
	DB_USER   = "gorel_test"
	DB_PASSWD = "abc123"
	DB_NAME   = "gorel_test"
)

func TestGetLiteral(t *testing.T) {
	v := new(ToSql)
	l := new(Literal)

	t.Log("Test string")
	l.Value = "test"
	s := v.GetLiteral(*l)
	if s != "\"test\"" {
		t.Errorf("failed test string")
	}

	t.Log("Test boolean")
	l.Value = false
	s = v.GetLiteral(*l)
	if s != "false" {
		t.Errorf("failed test boolean")
	}

	t.Log("Test int")
	l.Value = 123
	s = v.GetLiteral(*l)
	if s != "123" {
		t.Errorf("failed test int")
	}

	t.Log("Test float")
	l.Value = 123.3
	s = v.GetLiteral(*l)
	if s != "123.3" {
		t.Errorf("failed test float")
	}

	t.Log("Test Array of ints")
	l.Value = []int{1, 2, 3, 4}
	s = v.GetLiteral(*l)
	if s != "1,2,3,4" {
		t.Log(s)
		t.Errorf("failed test float")
	}

	t.Log("Test Array of ints")
	l.Value = []int{1, 2, 3, 4}
	s = v.GetLiteral(*l)
	if s != "1,2,3,4" {
		t.Log(s)
		t.Errorf("failed test float")
	}

	t.Log("Test Array of strings")
	l.Value = []string{"a", "bc", "d", "ef1"}
	s = v.GetLiteral(*l)
	if s != "\"a\",\"bc\",\"d\",\"ef1\"" {
		t.Log(s)
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

	n1 := &Literal{1}
	n2 := &Literal{2}
	n3 := &Literal{3}
	n.Children = []Node{n1, n2, n3}

	s := v.GetAnd(*n)
	if s != "1 AND 2 AND 3" {
		t.Log(s)
		t.Errorf("failed to get And")
	}
}

//-----------------Binary----------------
func TestGetBetween(t *testing.T) {
	v := new(ToSql)
	n := new(Between)
	n.Left = &Literal{1}
	n.Right = And{[]Node{
		&Literal{0},
		&Literal{3}}}

	s := v.GetBetween(*n)
	if s != "1 BETWEEN 0 AND 3" {
		t.Log(s)
		t.Errorf("failed to get Between ")
	}
}

func TestGetNotEqual(t *testing.T) {
	v := new(ToSql)
	n := new(NotEqual)
	n.Left = &Literal{1}

	s := v.GetNotEqual(*n)
	if s != "1 IS NOT NULL" {
		t.Log(s)
		t.Errorf("failed to get NotEqual null right")
	}

	n.Right = &Literal{2}
	s = v.GetNotEqual(*n)
	if s != "1 != 2" {
		t.Log(s)
		t.Errorf("failed to get NotEqual 1 != 2")
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

	n.Left = &Literal{1}
	n.Right = &Literal{2}

	s := v.GetOr(*n)
	if s != "1 OR 2" {
		t.Log(s)
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

	n.Left = &Literal{1}
	n.Right = &Literal{2}

	s := v.GetGreaterThan(*n)
	if s != "1 > 2" {
		t.Log(s)
		t.Errorf("failed to get GreaterThan ")
	}
}

func TestGetGreaterThanOrEqual(t *testing.T) {
	v := new(ToSql)
	n := new(GreaterThanOrEqual)

	n.Left = &Literal{2}
	n.Right = &Literal{1}

	s := v.GetGreaterThanOrEqual(*n)
	if s != "2 >= 1" {
		t.Log(s)
		t.Errorf("failed to get GreaterThanOrEqual ")
	}
}

func TestGetLessThan(t *testing.T) {
	v := new(ToSql)
	n := new(LessThan)

	n.Left = &Literal{1}
	n.Right = &Literal{2}

	s := v.GetLessThan(*n)
	if s != "1 < 2" {
		t.Log(s)
		t.Errorf("failed to get LessThan ")
	}
}

func TestGetLessThanOrEqual(t *testing.T) {
	v := new(ToSql)
	n := new(LessThanOrEqual)

	n.Left = &Literal{2}
	n.Right = &Literal{1}

	s := v.GetLessThanOrEqual(*n)
	if s != "2 <= 1" {
		t.Log(s)
		t.Errorf("failed to get LessThanOrEqual ")
	}
}

func TestGetMatches(t *testing.T) {
	v := new(ToSql)
	n := new(Matches)

	n.Left = &Literal{2}
	n.Right = &Literal{1}

	s := v.GetMatches(*n)
	if s != "2 LIKE 1" {
		t.Log(s)
		t.Errorf("failed to get Matches ")
	}
}

func TestGetDoesNotMatch(t *testing.T) {
	v := new(ToSql)
	n := new(DoesNotMatch)

	n.Left = &Literal{2}
	n.Right = &Literal{1}

	s := v.GetDoesNotMatch(*n)
	if s != "2 NOT LIKE 1" {
		t.Log(s)
		t.Errorf("failed to get DoesNotMatch ")
	}
}

func TestGetNotIn(t *testing.T) {
	v := new(ToSql)
	n := new(NotIn)

	n.Left = &Literal{"abc"}
	n.Right = &Literal{[]string{"1", "2", "3", "4"}}
	s := v.GetNotIn(*n)
	if s != "\"abc\" NOT IN (\"1\",\"2\",\"3\",\"4\")" {
		t.Log(s)
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
	connection, _ := db.MySQLNewConnection(DB_SOCK, DB_USER, DB_PASSWD, DB_NAME)
	tableSchema, _ := connection.GetTable("Users")
	table := Table{tableSchema, []Node{}}
	n := TableAlias{Binary{table, SqlLiteral{"u"}}}

	v := ToSql{connection}
	s := v.GetTableAlias(n)
	if s != "`Users` `u`" {
		t.Log(s)
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

	n.Left = &Literal{1}
	n.Right = &Literal{2}

	s := v.GetEquality(*n)
	if s != "1 = 2" {
		t.Log(s)
		t.Errorf("failed to get Equality")
	}
}

func TestGetIn(t *testing.T) {
	v := new(ToSql)
	n := new(In)

	n.Left = &Literal{"abc"}
	n.Right = &Literal{[]string{"1", "2", "3", "4"}}

	s := v.GetIn(*n)
	if s != "\"abc\" IN (\"1\",\"2\",\"3\",\"4\")" {
		t.Log(s)
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

//-----------------InfixOperation----------------
func TestGetMultiplication(t *testing.T) {
	v := new(ToSql)
	n := new(Multiplication)

	n.Left = &Literal{1}
	n.Right = &Literal{2}

	s := v.GetMultiplication(*n)
	if s != "1 * 2" {
		t.Log(s)
		t.Errorf("failed to get Multiplication ")
	}
}

func TestGetDivision(t *testing.T) {
	v := new(ToSql)
	n := new(Division)

	n.Left = &Literal{1}
	n.Right = &Literal{2}

	s := v.GetDivision(*n)
	if s != "1 / 2" {
		t.Log(s)
		t.Errorf("failed to get Division ")
	}
}

func TestGetAddition(t *testing.T) {
	v := new(ToSql)
	n := new(Addition)

	n.Left = &Literal{1}
	n.Right = &Literal{2}

	s := v.GetAddition(*n)
	if s != "1 + 2" {
		t.Log(s)
		t.Errorf("failed to get Addition ")
	}
}

func TestGetSubtraction(t *testing.T) {
	v := new(ToSql)
	n := new(Subtraction)

	n.Left = &Literal{1}
	n.Right = &Literal{2}

	s := v.GetSubtraction(*n)
	if s != "1 - 2" {
		t.Log(s)
		t.Errorf("failed to get Subtraction ")
	}
}

//-----------------Join----------------
func TestGetInnerJoin(t *testing.T) {
	v := new(ToSql)
	v.Connection, _ = db.MySQLNewConnection(DB_SOCK, DB_USER, DB_PASSWD, DB_NAME)
  table := Table{db.TableSchema{Name:"users"}, []Node{}}
  on := On{Unary{Equality{Binary{&Literal{1}, &Literal{2}}}}}
	n := InnerJoin{&BaseJoin{table, on}}
	s := v.GetInnerJoin(n)
	if s != "INNER JOIN `users` ON 1 = 2" {
		t.Log(s)
		t.Errorf("failed to get InnerJoin ")
	}
}

func TestGetOuterJoin(t *testing.T) {
	v := new(ToSql)
	v.Connection, _ = db.MySQLNewConnection(DB_SOCK, DB_USER, DB_PASSWD, DB_NAME)
  table := Table{db.TableSchema{Name:"users"}, []Node{}}
  on := On{Unary{Equality{Binary{&Literal{1}, &Literal{2}}}}}
	n := OuterJoin{&BaseJoin{table, on}}

	s := v.GetOuterJoin(n)
	if s != "LEFT OUTER JOIN `users` ON 1 = 2" {
		t.Log(s)
		t.Errorf("failed to get OuterJoin ")
	}
}

func TestGetStringJoin(t *testing.T) {
	v := new(ToSql)
	n := StringJoin{"test"}

	s := v.GetStringJoin(n)
	if s != "test" {
		t.Log(s)
		t.Errorf("failed to get StringJoin ")
	}
}

//-----------------Unary----------------
func TestGetNot(t *testing.T) {
	v := new(ToSql)
	n := new(Not)

	n.Expression = Equality{Binary{&Literal{1}, &Literal{2}}}

	s := v.GetNot(*n)
	if s != "NOT (1 = 2)" {
		t.Log(s)
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

	n.Expression = &Literal{10}

	s := v.GetOffset(*n)
	if s != "OFFSET 10" {
		t.Log(s)
		t.Errorf("failed to get Offset ")
	}
}

func TestGetLimit(t *testing.T) {
	v := new(ToSql)
	n := new(Limit)

	n.Expression = &Literal{100}

	s := v.GetLimit(*n)
	if s != "LIMIT 100" {
		t.Log(s)
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

	n.Expression = Equality{Binary{&Literal{1}, &Literal{2}}}

	s := v.GetHaving(*n)
	if s != "HAVING (1 = 2)" {
		t.Log(s)
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
	eq1 := Equality{Binary{&Literal{1}, &Literal{2}}}
	eq2 := Equality{Binary{&Literal{3}, &Literal{4}}}
	and := And{[]Node{eq1, eq2}}
	n.Expression = and
	s := v.GetGrouping(*n)
	if s != "(1 = 2 AND 3 = 4)" {
		t.Log(s)
		t.Errorf("failed to get Grouping ")
	}
}

func TestGetOn(t *testing.T) {
	v := new(ToSql)
	n := new(On)

	n.Expression = Equality{Binary{&Literal{1}, &Literal{2}}}

	s := v.GetOn(*n)
	if s != "ON 1 = 2" {
		t.Log(s)
		t.Errorf("failed to get On ")
	}
}

// ----- not generated ----

func TestGetField(t *testing.T) {
	connection, _ := db.MySQLNewConnection(DB_SOCK, DB_USER, DB_PASSWD, DB_NAME)
	tableSchema, _ := connection.GetTable("Users")
	table := Table{tableSchema, []Node{}}

	n := Field{table, table.ColumnMap["id"]}
	v := ToSql{connection}

	s := v.GetField(n)
	if s != "`Users`.`id`" {
		t.Log(s)
		t.Errorf("failed to get Field ")
	}
}

func TestGetSelectCore(t *testing.T) {
	n := new(SelectCore)

	connection, _ := db.MySQLNewConnection(DB_SOCK, DB_USER, DB_PASSWD, DB_NAME)
	tableSchema, _ := connection.GetTable("Users")
	table := Table{tableSchema, []Node{}}
	table_alias := TableAlias{Binary{table, SqlLiteral{"u"}}}
	f := Field{table_alias, table.ColumnMap["id"]}

	n.Projections = []Node{SqlLiteral{"*"}}
	n.Source = JoinSource{table_alias, make([]Node, 0)}
	n.Wheres = []Node{f.Eq(&Literal{1})}

	v := ToSql{connection}
	s := v.GetSelectCore(*n)
	if s != "SELECT * FROM `Users` `u`  WHERE `u`.`id` = 1" {
		t.Log(s)
		t.Errorf("failed to get Select Core ")
	}

}

func TestVisitGetTable(t *testing.T) {
	connection, _ := db.MySQLNewConnection(DB_SOCK, DB_USER, DB_PASSWD, DB_NAME)
	table, _ := connection.GetTable("Users")
	n := Table{table, []Node{}}

	v := ToSql{connection}
	s := v.GetTable(n)
	if s != "`Users`" {
		t.Log(s)
		t.Errorf("failed to get table ")
	}
}

func TestVisitGetTableAlias(t *testing.T) {
	connection, _ := db.MySQLNewConnection(DB_SOCK, DB_USER, DB_PASSWD, DB_NAME)
	table, _ := connection.GetTable("Users")
	n := Table{table, []Node{}}
	s := n.GetNameAlias()
	if s != "Users_0" {
		t.Log(s)
		t.Errorf("failed to get table ")
	}
	n.CreateTableAlias()
	s = n.GetNameAlias()
	if s != "Users_1" {
		t.Log(s)
		t.Errorf("failed to get table ")
	}
}

func TestVisitGetJoinSource(t *testing.T) {
	connection, _ := db.MySQLNewConnection(DB_SOCK, DB_USER, DB_PASSWD, DB_NAME)
	tableSchema, _ := connection.GetTable("Users")
	table := Table{tableSchema, []Node{}}
	table_alias := TableAlias{Binary{table, SqlLiteral{"u"}}}

	n := JoinSource{table_alias, make([]Node, 0)}
	v := ToSql{connection}
	s := v.GetJoinSource(n)
	if s != "FROM `Users` `u` " {
		t.Log(s)
		t.Errorf("failed to get join source ")
	}
}

func TestVisitGetSqlLiteral(t *testing.T) {
	v := new(ToSql)
	n := SqlLiteral{"test"}
	s := v.GetSqlLiteral(n)
	if s != "test" {
		t.Log(s)
		t.Errorf("failed to get select statement ")
	}
}

func TestVisitGetSelectStatement(t *testing.T) {
	connection, _ := db.MySQLNewConnection(DB_SOCK, DB_USER, DB_PASSWD, DB_NAME)
	tableSchema, _ := connection.GetTable("Users")
	table := Table{tableSchema, []Node{}}
	table_alias := TableAlias{Binary{table, SqlLiteral{"u"}}}
	f := Field{table_alias, table.ColumnMap["id"]}

	n := NewSelectStatement()
	if len(n.Cores) == 0 {
		t.Errorf("Expected a core in new select statement")
	}

	c := n.Cores[len(n.Cores)-1].(SelectCore)
	c.Projections = []Node{SqlLiteral{"*"}}
	c.Source = JoinSource{table_alias, make([]Node, 0)}
	c.Wheres = []Node{f.Eq(&Literal{1})}

	n.Cores[len(n.Cores)-1] = c
	n.Orders = append(n.Orders, f)
	n.Limit = Limit{Unary{&Literal{1}}}
	n.Offset = Offset{Unary{&Literal{1}}}
	v := ToSql{connection}
	s := v.GetSelectStatement(n)
	if s != "SELECT * FROM `Users` `u`  WHERE `u`.`id` = 1 ORDER BY `u`.`id` LIMIT 1 OFFSET 1" {
		t.Log(s)
		t.Errorf("failed to get select statement ")
	}
}

func TestVisitGetWith(t *testing.T) {
	v := new(ToSql)
	n := new(With)
	s := v.GetWith(*n)
	if s != "FROM `Users` `u` " {
		t.Log(s)
		t.Errorf("failed to get with")
	}
}
