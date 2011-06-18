package gorel

import (
  "fmt"
  "strconv"
  "strings"
)

type ToSql struct {
}

func (c ToSql) GetLiteral(n Literal) (s string) {
  switch val := n.value.(type) {
    case string:
      s = val
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
  }
  return
}

func (b ToSql) VisitNodes(nodes []Node) (s string) {
  s = ""
  results := make([]string, 0)
  for i:=0; i < len(nodes); i++ {
    if (nodes[i] != nil) {
      results = append(results, nodes[i].Visit(b))
    }
  }
  s = strings.Join(results, ", ")
  return
}

//-----------------Binary----------------
func (c ToSql) GetBetween(n Between) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}

func (c ToSql) GetNotEqual(n NotEqual) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}

func (c ToSql) GetAssignment(n Assignment) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}

func (c ToSql) GetOr(n Or) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}

func (c ToSql) GetAnd(n And) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}

func (c ToSql) GetAs(n As) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}

func (c ToSql) GetGreaterThan(n GreaterThan) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}

func (c ToSql) GetGreaterThanOrEqual(n GreaterThanOrEqual) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}

func (c ToSql) GetLessThan(n LessThan) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}

func (c ToSql) GetLessThanOrEqual(n LessThanOrEqual) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}

func (c ToSql) GetMatches(n Matches) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}

func (c ToSql) GetDoesNotMatch(n DoesNotMatch) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}

func (c ToSql) GetNotIn(n NotIn) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}

func (c ToSql) GetOrdering(n Ordering) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}

func (c ToSql) GetValues(n Values) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}

func (c ToSql) GetDeleteStatement(n DeleteStatement) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}

func (c ToSql) GetTableAlias(n TableAlias) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}

func (c ToSql) GetExcept(n Except) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}

func (c ToSql) GetIntersect(n Intersect) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}

func (c ToSql) GetUnion(n Union) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}

func (c ToSql) GetUnionAll(n UnionAll) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}


//-----------------Equality----------------
func (c ToSql) GetEquality(n Equality) (s string) {
  ls := n.left.Visit(c)
  rs := n.right.Visit(c)
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}

func (c ToSql) GetIn(n In) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}


//-----------------Function----------------
func (c ToSql) GetCount(n Count) (s string) {
  expressions := c.VisitNodes(n.expressions)
  alias := n.alias.Visit(c)
  distinct := n.distinct
  s = fmt.Sprintf("%v * %v * %v", expressions, alias, distinct)
  return s

}

func (c ToSql) GetSum(n Sum) (s string) {
  expressions := c.VisitNodes(n.expressions)
  alias := n.alias.Visit(c)
  distinct := n.distinct
  s = fmt.Sprintf("%v * %v * %v", expressions, alias, distinct)
  return s

}

func (c ToSql) GetExists(n Exists) (s string) {
  expressions := c.VisitNodes(n.expressions)
  alias := n.alias.Visit(c)
  distinct := n.distinct
  s = fmt.Sprintf("%v * %v * %v", expressions, alias, distinct)
  return s

}

func (c ToSql) GetMax(n Max) (s string) {
  expressions := c.VisitNodes(n.expressions)
  alias := n.alias.Visit(c)
  distinct := n.distinct
  s = fmt.Sprintf("%v * %v * %v", expressions, alias, distinct)
  return s

}

func (c ToSql) GetMin(n Min) (s string) {
  expressions := c.VisitNodes(n.expressions)
  alias := n.alias.Visit(c)
  distinct := n.distinct
  s = fmt.Sprintf("%v * %v * %v", expressions, alias, distinct)
  return s

}

func (c ToSql) GetAvg(n Avg) (s string) {
  expressions := c.VisitNodes(n.expressions)
  alias := n.alias.Visit(c)
  distinct := n.distinct
  s = fmt.Sprintf("%v * %v * %v", expressions, alias, distinct)
  return s

}


//-----------------Join----------------
func (c ToSql) GetInnerJoin(n InnerJoin) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}

func (c ToSql) GetOuterJoin(n OuterJoin) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}

func (c ToSql) GetStringJoin(n StringJoin) (s string) {
  ls := ""
  if (n.left != nil) {
    ls = n.left.Visit(c)
  }
  
  rs := ""
  if (n.right != nil) {
    rs = n.right.Visit(c)
  }
  
  s = fmt.Sprintf("%v * %v", ls, rs)
  return
}


//-----------------Unary----------------
func (c ToSql) GetNot(n Not) (s string) {
  expr := n.expression.Visit(c)
  s = expr
  return s

}

func (c ToSql) GetLock(n Lock) (s string) {
  expr := n.expression.Visit(c)
  s = expr
  return s

}

func (c ToSql) GetOffset(n Offset) (s string) {
  expr := n.expression.Visit(c)
  s = expr
  return s

}

func (c ToSql) GetLimit(n Limit) (s string) {
  expr := n.expression.Visit(c)
  s = expr
  return s

}

func (c ToSql) GetTop(n Top) (s string) {
  expr := n.expression.Visit(c)
  s = expr
  return s

}

func (c ToSql) GetHaving(n Having) (s string) {
  expr := n.expression.Visit(c)
  s = expr
  return s

}

func (c ToSql) GetUnqualifiedColumn(n UnqualifiedColumn) (s string) {
  expr := n.expression.Visit(c)
  s = expr
  return s

}

func (c ToSql) GetGroup(n Group) (s string) {
  expr := n.expression.Visit(c)
  s = expr
  return s

}

func (c ToSql) GetGrouping(n Grouping) (s string) {
  expr := n.expression.Visit(c)
  s = expr
  return s

}

func (c ToSql) GetOn(n On) (s string) {
  expr := n.expression.Visit(c)
  s = expr
  return s

}

