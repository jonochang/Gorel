package ast

type Ascending struct {
	Unary
}

func (n Ascending) Visit(v Visitor) (s string) {
	s = v.GetAscending(n)
	return
}

func (n Ascending) IsAscending() bool  { return true }
func (n Ascending) IsDescending() bool { return false }
func (n Ascending) Direction() string  { return "ASC" }
//func (n Ascending) Reverse() (Node) { return Descending{Unary{n.Expression}} }
