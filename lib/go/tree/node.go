package tree

import (
	"log/slog"
	"slices"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
	"github.com/samber/oops"
)

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

// error() Is a helper to construct an oops error
func (n *Node[T]) error() (err oops.OopsErrorBuilder) {
	return oops.With("tree", n)
}

// ─────────────────────────────────────────────────────────────────────────────

// GetPath gets a path from the tree, if the path is a nested path, it will
// recursively get the path from the tree.
//
//   - If the path does not exist in the tree it will return the closest path
//     that does exist. (i.e. if the path is "a.b.c" and "a.b" exists, it will
//     return "a.b")
//   - if the value of the path is nil, it will return the closest parent path
//     that does have a value. (i.e. if the path is "a.b.c" and "a.b" exists,
//     but "a.b.c" does not exist, it will return node "a.b")
//   - This allows for nodes to have sub-paths with a set value, but not
//     have to add value itself.
//   - If the path is empty, it will return the root node.
func (n *Node[T]) GetPath(path string) (node *Node[T]) {
	// Split the path
	split := strings.Split(path, PathSeparator)

	return n.getPath(nil, split)
}

// getPath - gets a path from the tree, if the path is a nested path, it will
// recursively get the path from the tree.
//
//   - If the path does not exist in the tree it will return the closest path
//     that is defined
//   - if the value of the path is nil, it will return the closest parent path
//     that does have a value
//   - The initial parent node should be nil, it is used for recursion
func (n *Node[T]) getPath(p *Node[T], split []string) (node *Node[T]) {
	// Check if the path is empty
	if len(split) == 0 {
		if n.Value == nil {
			return p
		}
		return n
	}

	// Check to see if the sub-path exists
	// if not, return nil
	subPath, ok := n.SubPaths[split[0]]
	if !ok {
		return n
	}

	// Get the sub-path
	return subPath.getPath(n, split[1:])
}

// ─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

func (n *Node[T]) Print() {
	log.Print(&n)
}

// ─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

type VisitorOptions struct {
	// ResolvePaths will resolve the paths to the closest parent path that has a value
	ResolvePaths bool
	// Deterministic will visit the paths in a deterministic order (i.e. the sorted order of the paths)
	Deterministic bool

	// TODO: Add more options
}

// Visit visits each node in the tree and calls the visitor function in a depth first manner
func (n *Node[T]) Visit(options VisitorOptions, visitFn func(node *Node[T]) string) {
	// Visit the current node
	n.visit(VisitorOptions{}, visitFn)
}

// visit recursively visits each node in the tree and calls the visitor function in a depth first manner
//
//   - It optionally visits the paths in a deterministic order
func (n *Node[T]) visit(options VisitorOptions, visitFn func(node *Node[T]) string) {
	visitFn(n)

	if options.Deterministic {
		keys := lo.Keys(n.SubPaths)
		slices.Sort(keys)

		for _, key := range keys {
			subPath := n.SubPaths[key]
			p := visitFn(subPath)
			slog.Info("visit", "p", p)
			subPath.visit(options, visitFn)
		}
	} else {
		for _, subPath := range n.SubPaths {
			subPath.visit(options, visitFn)
		}
	}
}

// VisitPaths visits all nodes in the supplied paths and calls the visitor function in a depth first manner
// It optionally resolves the paths to the closest parent path that has a value
//
//   - If the path is empty, it will return the root node.
//   - Groups the paths by the first path segment, then visits each group in order of the path segments
func (n *Node[T]) VisitPaths(paths []string, resolve bool, visitor func(node *Node[T])) {
	// Visit the sub-paths
	for _, path := range paths {
		if subPath, ok := n.SubPaths[path]; ok {
			subPath.VisitPaths(paths, resolve, visitor)
		}
	}
}
