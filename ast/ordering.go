package ast

type Ordering interface {
	IsAscending() bool
	IsDescending() bool
	Direction() string
	Reverse() Ordering
	Node
}
