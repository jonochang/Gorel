package ast

type Descending struct {
	Unary
}

func (n Descending) Visit(v Visitor) (s string) {
	s = v.GetDescending(n)
	return
}

func (n Descending) IsAscending() bool  { return false }
func (n Descending) IsDescending() bool { return true }
func (n Descending) Direction() string  { return "DESC" }
func (n Descending) Reverse() Ordering  { return Ascending{Unary{n.Expression}} }
