package tree

import "strings"

type Node[T any] struct {
	Path          string
	SubPaths      map[string]*Node[T]
	AllowSubPaths bool
	Value         *T
}

func NewTree[T any]() *Node[T] {
	return &Node[T]{
		SubPaths:      make(map[string]*Node[T]),
		AllowSubPaths: true,
	}
}

func (n *Node[T]) AddPath(path string, value *T) {
	// Split the path
	split := strings.Split(path, ".")

	// Add the path
	n.addPath(split, value)
}

func (n *Node[T]) addPath(split []string, value *T) {
	// Check if the path is empty
	if len(split) == 0 {
		n.Value = value
		return
	}

	// Get the first path
	first := split[0]

	// Check if the subpath exists
	subPath, ok := n.SubPaths[first]
	if !ok {
		// Create the subpath
		subPath = &Node[T]{
			SubPaths:      make(map[string]*Node[T]),
			AllowSubPaths: true,
		}
		n.SubPaths[first] = subPath
	}

	// Add the subpath
	subPath.addPath(split[1:], value)
}
