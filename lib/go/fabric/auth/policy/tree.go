package policy

const (
	// PathSeparator is the path separator
	PathSeparator = "."
	DefaultSize   = 10
)

type Node[T any] struct {
	FullPath      string
	Path          string
	SubPaths      map[string]*Node[T]
	AllowSubPaths bool
	Value         *T
}

func NewTree[T any]() *Node[T] {
	return &Node[T]{
		FullPath:      "",
		Path:          "",
		SubPaths:      make(map[string]*Node[T], DefaultSize),
		AllowSubPaths: true,
		Value:         new(T),
	}
}
