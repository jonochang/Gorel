package ast

import (
	"fmt"
	"strconv"
	"strings"
	"reflect"
	"db"
)

type ToSql struct {
	Connection db.Connection
}

func (c *ToSql) GetLiteral(n Literal) (s string) {
	s = c.Convert(n.Value)
	return
}

func (c *ToSql) Convert(unknown interface{}) (s string) {
	switch val := unknown.(type) {
	case string:
		s = strconv.Quote(val)
	case bool:
		s = strconv.Btoa(val)
	case int:
		s = strconv.Itoa(val)
	case int64:
		s = strconv.Itoa64(val)
	case uint:
		s = strconv.Uitoa(val)
	case uint64:
		s = strconv.Uitoa64(val)
	case float32:
		s = strconv.Ftoa32(val, 'f', -1)
	case float64:
		s = strconv.Ftoa64(val, 'f', -1)
	default:
		s = c.ConvertUnknown(unknown)
	}
	return
}

func (c *ToSql) ConvertUnknown(unknown interface{}) (s string) {
	unknownVal := reflect.ValueOf(unknown)
	kind := unknownVal.Kind()
	switch kind {
	case reflect.Array, reflect.Slice:
		results := make([]string, 0)
		for i := 0; i < unknownVal.Len(); i++ {
			results = append(results, c.ConvertVal(unknownVal.Index(i)))
		}
		s = strings.Join(results, ",")
	}
	return
}
func (c *ToSql) ConvertVal(val reflect.Value) (s string) {
	switch val.Kind() {
	case reflect.Bool:
		s = c.Convert(val.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		s = c.Convert(val.Int())
	case reflect.Float32, reflect.Float64:
		s = c.Convert(val.Float())
	case reflect.Array, reflect.Slice:
		if val.Len() > 0 {
			s = c.Convert(val.Slice(0, val.Len()))
		}
	case reflect.String:
		s = c.Convert(val.String())
	default:
		panic(fmt.Sprintf("Cannot convert %v", val.Kind()))

	}
	return
}

func (b *ToSql) VisitNodes(nodes []Node) (s []string) {
	s = make([]string, 0)
	for i := 0; i < len(nodes); i++ {
		if nodes[i] != nil {
			s = append(s, nodes[i].Visit(b))
		}
	}
	return
}

func (b *ToSql) VisitNodesString(nodes []Node, sep string) (s string) {
	results := b.VisitNodes(nodes)
	s = strings.Join(results, sep)
	return
}

//-----------------And----------------
func (c *ToSql) GetAnd(n And) (s string) {
	results := c.VisitNodes(n.Children)
	s = strings.Join(results, " AND ")
	return
}


//-----------------Binary----------------
func (c *ToSql) GetBetween(n Between) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("%v BETWEEN %v", ls, rs)
	return
}

func (c *ToSql) GetNotEqual(n NotEqual) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	if n.Right == nil {
		s = fmt.Sprintf("%v IS NOT NULL", ls)
	} else {
		s = fmt.Sprintf("%v != %v", ls, rs)
	}
	return
}

func (c *ToSql) GetAssignment(n Assignment) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("%v * %v", ls, rs)
	return
}

func (c *ToSql) GetOr(n Or) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("%v OR %v", ls, rs)
	return
}

func (c *ToSql) GetAs(n As) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("%v AS %v", ls, rs)
	return
}

func (c *ToSql) GetGreaterThan(n GreaterThan) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("%v > %v", ls, rs)
	return
}

func (c *ToSql) GetGreaterThanOrEqual(n GreaterThanOrEqual) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("%v >= %v", ls, rs)
	return
}

func (c *ToSql) GetLessThan(n LessThan) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("%v < %v", ls, rs)
	return
}

func (c *ToSql) GetLessThanOrEqual(n LessThanOrEqual) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("%v <= %v", ls, rs)
	return
}

func (c *ToSql) GetMatches(n Matches) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("%v LIKE %v", ls, rs)
	return
}

func (c *ToSql) GetDoesNotMatch(n DoesNotMatch) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("%v NOT LIKE %v", ls, rs)
	return
}

func (c *ToSql) GetNotIn(n NotIn) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("%v NOT IN (%v)", ls, rs)
	return
}

func (c *ToSql) GetAscending(n Ascending) (s string) {
	expr := ""
	if n.Expression != nil {
		expr = n.Expression.Visit(c)
	}

	s = fmt.Sprintf("%v ASC", expr)
	return
}

func (c *ToSql) GetDescending(n Descending) (s string) {
	expr := ""
	if n.Expression != nil {
		expr = n.Expression.Visit(c)
	}

	s = fmt.Sprintf("%v DESC", expr)
	return
}

func (c *ToSql) GetValues(n Values) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("%v * %v", ls, rs)
	return
}

func (c *ToSql) GetDeleteStatement(n DeleteStatement) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("%v * %v", ls, rs)
	return
}

func (c *ToSql) GetTableAlias(n TableAlias) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Alias != nil {
		source := n.Alias.Visit(c)
		rs = c.Connection.QuoteTableName(source)
	}

	s = fmt.Sprintf("%v %v", ls, rs)
	return
}

func (c *ToSql) GetExcept(n Except) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("%v * %v", ls, rs)
	return
}

func (c *ToSql) GetIntersect(n Intersect) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("%v * %v", ls, rs)
	return
}

func (c *ToSql) GetUnion(n Union) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("%v * %v", ls, rs)
	return
}

func (c *ToSql) GetUnionAll(n UnionAll) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("%v * %v", ls, rs)
	return
}


//-----------------Equality----------------
func (c *ToSql) GetEquality(n Equality) (s string) {
	ls := n.Left.Visit(c)
	rs := n.Right.Visit(c)
	s = fmt.Sprintf("%v = %v", ls, rs)
	return
}

func (c *ToSql) GetIn(n In) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("%v IN (%v)", ls, rs)
	return
}


//-----------------Function----------------
func (c *ToSql) GetCount(n CountNode) (s string) {
	expressions := c.VisitNodesString(n.Expressions, ", ")

	alias := ""
	if n.Alias != nil {
		alias = fmt.Sprintf("AS %v", n.Alias.Visit(c))
	}

	distinct := ""
	if n.Distinct {
		distinct = "DISTINCT "
	}

	s = fmt.Sprintf("COUNT(%v%v) %v", distinct, expressions, alias)
	return s

}

func (c *ToSql) GetSum(n SumNode) (s string) {
	expressions := c.VisitNodesString(n.Expressions, ", ")
	alias := ""
	if n.Alias != nil {
		alias = fmt.Sprintf("AS %v", n.Alias.Visit(c))
	}
	s = fmt.Sprintf("SUM(%v) %v", expressions, alias)
	return s
}

func (c *ToSql) GetExists(n Exists) (s string) {
	expressions := c.VisitNodesString(n.Expressions, ", ")
	alias := ""
	if n.Alias != nil {
		alias = n.Alias.Visit(c)
	}
	distinct := n.Distinct
	s = fmt.Sprintf("%v * %v * %v", expressions, alias, distinct)
	return s

}

func (c *ToSql) GetMax(n MaxNode) (s string) {
	expressions := c.VisitNodesString(n.Expressions, ", ")
	alias := ""
	if n.Alias != nil {
		alias = fmt.Sprintf("AS %v", n.Alias.Visit(c))
	}
	s = fmt.Sprintf("MAX(%v) %v", expressions, alias)
	return s
}

func (c *ToSql) GetMin(n MinNode) (s string) {
	expressions := c.VisitNodesString(n.Expressions, ", ")
	alias := ""
	if n.Alias != nil {
		alias = fmt.Sprintf("AS %v", n.Alias.Visit(c))
	}
	s = fmt.Sprintf("MIN(%v) %v", expressions, alias)
	return s

}

func (c *ToSql) GetAvg(n AvgNode) (s string) {
	expressions := c.VisitNodesString(n.Expressions, ", ")
	alias := ""
	if n.Alias != nil {
		alias = fmt.Sprintf("AS %v", n.Alias.Visit(c))
	}
	s = fmt.Sprintf("AVG(%v) %v", expressions, alias)
	return s

}


//-----------------InfixOperation----------------
func (c *ToSql) GetMultiplication(n Multiplication) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("%v * %v", ls, rs)
	return
}

func (c *ToSql) GetDivision(n Division) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("%v / %v", ls, rs)
	return
}

func (c *ToSql) GetAddition(n Addition) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("%v + %v", ls, rs)
	return
}

func (c *ToSql) GetSubtraction(n Subtraction) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("%v - %v", ls, rs)
	return
}


//-----------------Join----------------
func (c *ToSql) GetInnerJoin(n InnerJoin) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("INNER JOIN %v %v", ls, rs)
	return
}

func (c *ToSql) GetOuterJoin(n OuterJoin) (s string) {
	ls := ""
	if n.Left != nil {
		ls = n.Left.Visit(c)
	}

	rs := ""
	if n.Right != nil {
		rs = n.Right.Visit(c)
	}

	s = fmt.Sprintf("LEFT OUTER JOIN %v %v", ls, rs)
	return
}

func (c *ToSql) GetStringJoin(n StringJoin) (s string) {
	s = n.name
	return
}


//-----------------Unary----------------
func (c *ToSql) GetNot(n Not) (s string) {
	expr := ""
	if n.Expression != nil {
		expr = n.Expression.Visit(c)
	}
	s = fmt.Sprintf("NOT (%v)", expr)
	return s

}

func (c *ToSql) GetLock(n Lock) (s string) {
	expr := "*"
	if n.Expression != nil {
		n.Expression.Visit(c)
	}
	s = expr
	return s

}

func (c *ToSql) GetOffset(n Offset) (s string) {
	expr := ""
	if n.Expression != nil {
		expr = n.Expression.Visit(c)
	}
	s = fmt.Sprintf("OFFSET %v", expr)
	return s

}

func (c *ToSql) GetLimit(n Limit) (s string) {
	expr := ""
	if n.Expression != nil {
		expr = n.Expression.Visit(c)
	}
	s = fmt.Sprintf("LIMIT %v", expr)
	return s

}

func (c *ToSql) GetTop(n Top) (s string) {
	expr := "*"
	if n.Expression != nil {
		n.Expression.Visit(c)
	}
	s = expr
	return s

}

func (c *ToSql) GetHaving(n Having) (s string) {
	expr := ""
	if n.Expression != nil {
		expr = n.Expression.Visit(c)
	}
	s = fmt.Sprintf("HAVING (%v)", expr)
	return s

}

func (c *ToSql) GetUnqualifiedColumn(n UnqualifiedColumn) (s string) {
	expr := "*"
	if n.Expression != nil {
		n.Expression.Visit(c)
	}
	s = expr
	return s

}

func (c *ToSql) GetGroup(n Group) (s string) {
	s = ""
	if n.Expression != nil {
		s = n.Expression.Visit(c)
	}
	return s

}

func (c *ToSql) GetGrouping(n Grouping) (s string) {
	expr := ""
	if n.Expression != nil {
		expr = n.Expression.Visit(c)
	}
	s = fmt.Sprintf("(%v)", expr)
	return s

}

func (c *ToSql) GetOn(n On) (s string) {
	expr := ""
	if n.Expression != nil {
		expr = n.Expression.Visit(c)
	}
	s = fmt.Sprintf("ON %v", expr)
	return s

}

func (c *ToSql) GetField(n Field) (s string) {
	quoted_column_name := c.Connection.QuoteColumnName(n.Column.Name)
	quoted_table_name := c.Connection.QuoteTableName(n.Table.GetName())
	if n.Table.HasAlias() {
		quoted_table_name = c.Connection.QuoteTableName(n.Table.GetNameAlias())
	}

	s = fmt.Sprintf("%v.%v", quoted_table_name, quoted_column_name)
	return
}

func (c *ToSql) GetSelectCore(n SelectCore) (s string) {
	results := make([]string, 0)
	results = append(results, "SELECT")

	if len(n.Projections) > 0 {
		projections := c.VisitNodes(n.Projections)
		projections_string := strings.Join(projections, ", ")
		results = append(results, projections_string)
	}

	if n.Source != nil {
		source := n.Source.Visit(c)
		results = append(results, source)
	}

	if len(n.Wheres) > 0 {
		wheres := c.VisitNodes(n.Wheres)
		where_string := strings.Join(wheres, " AND ")
		results = append(results, "WHERE "+where_string)
	}

	if len(n.Groups) > 0 {
		groups := c.VisitNodes(n.Groups)
		group_string := strings.Join(groups, ", ")
		results = append(results, "GROUP BY "+group_string)
	}

	s = strings.Join(results, " ")
	return
}

func (c *ToSql) GetTable(n Table) (s string) {
	quoted_table_name := c.Connection.QuoteTableName(n.Name)
	if n.HasAlias() {
		quoted_table_alias := c.Connection.QuoteTableName(n.GetNameAlias())
		s = fmt.Sprintf("%v %v", quoted_table_name, quoted_table_alias)
	} else {
		s = quoted_table_name
	}
	return
}

func (c *ToSql) GetJoinSource(n JoinSource) (s string) {
	if n.Source == nil ||
		n.JoinOn == nil {
		return ""
	}

	source := n.Source.Visit(c)
	joins := c.VisitNodes(n.JoinOn)
	joins_string := strings.Join(joins, " ")
	s = fmt.Sprintf("FROM %v %v", source, joins_string)
	return
}

func (c *ToSql) GetSqlLiteral(n SqlLiteral) (s string) {
	s = n.Value
	return
}

func (c *ToSql) GetSelectStatement(n SelectStatement) (s string) {
	results := make([]string, 0)

	if n.With != nil {
		with := n.With.Visit(c)
		results = append(results, with)
	}

	if len(n.Cores) > 0 {
		cores := c.VisitNodes(n.Cores)
		cores_string := strings.Join(cores, " ")
		results = append(results, cores_string)
	}

	if len(n.Orders) > 0 {
		orders := c.VisitNodes(n.Orders)
		orders_string := strings.Join(orders, ", ")
		results = append(results, "ORDER BY "+orders_string)
	}

	if n.Limit != nil {
		limit := n.Limit.Visit(c)
		results = append(results, limit)
	}

	if n.Offset != nil {
		offset := n.Offset.Visit(c)
		results = append(results, offset)
	}

	if n.Lock != nil {
		lock := n.Lock.Visit(c)
		results = append(results, lock)
	}

	s = strings.Join(results, " ")
	return
}

func (c *ToSql) GetWith(n With) (s string) {
	return
}
