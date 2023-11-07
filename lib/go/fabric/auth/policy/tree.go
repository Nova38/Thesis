package policy

type Node[T any] struct {
	Path          string
	SubPaths      map[string]*Node[T]
	AllowSubPaths bool
	Policy        *T
}

func NewTree[T any]() *Node[T] {
	return &Node[T]{
		SubPaths:      make(map[string]*Node[T]),
		AllowSubPaths: true,
	}
}
