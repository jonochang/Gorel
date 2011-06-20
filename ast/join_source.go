package ast

type JoinSource struct {
	Source Node
	JoinOn []Node
}

func (n JoinSource) Visit(v Visitor) (s string) {
	s = v.GetJoinSource(n)
	return
}
