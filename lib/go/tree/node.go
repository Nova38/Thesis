package tree

import (
	"slices"
	"strings"

	"github.com/samber/lo"
	"github.com/samber/oops"
)

const (
	// PathSeparator is the path separator
	PathSeparator = "."
	DefaultSize   = 10
)

type Node[T any] struct {
	Path          string
	SubPaths      map[string]Node[T]
	AllowSubPaths bool
	Value         *T
}

func NewTree[T any]() *Node[T] {
	return &Node[T]{
		Path:          "",
		SubPaths:      make(map[string]Node[T], DefaultSize),
		AllowSubPaths: true,
		Value:         new(T),
	}
}

// error() Is a helper to construct an oops error
func (n *Node[T]) error() (err oops.OopsErrorBuilder) {
	return oops.With("tree", n)
}

func (n *Node[T]) AddPath(path string, value *T) (err error) {
	// Split the path
	split := strings.Split(path, ".")

	// Add the path
	return n.addPath(split, value)
}

// addPath adds a path to the tree, if the path is a nested path, it will
// recursively add the path to the tree.
//
// - If the node does not allow sub-paths, it will return an error if it tries to add a sub-path
func (n *Node[T]) addPath(split []string, value *T) (err error) {
	// Check if the path is empty
	if len(split) == 0 {
		n.Value = value
		return
	}

	if !n.AllowSubPaths {
		return n.error().Errorf("cannot add sub-path to node that does not allow sub-paths")
	}

	// Check to see if the sub-path exists
	subPath, ok := n.SubPaths[split[0]]
	if !ok {
		n.SubPaths[split[0]] = Node[T]{
			Path:          split[0],
			SubPaths:      make(map[string]Node[T], DefaultSize),
			AllowSubPaths: true,
			Value:         nil,
		}
	}

	// Add the sub-path
	if err := subPath.addPath(split[1:], value); err != nil {
		return n.error().Wrap(err)
	}

	return nil
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
	split := strings.Split(path, ".")

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

// ─────────────────────────────────────────────────────────────────────────────

// Merge merges a variable number of trees into the current tree.  It prevents
// the merging of sub-paths from the sources trees if there is already a path
// in the current tree whose value is not nil.
//
//   - Copies the pointers of the trees into the current tree, so either pass in
//     a copy of the tree or be aware that the nodes could still be modified in
//     their source tree
//   - If a path exists in the current tree and the value is not nil, it will
//     not be overwritten,
//   - If the path does not exist in the tree it will be added and copied from
//     the source tree
//   - If shadowing is enabled the merge for sub-paths will be skipped if the
//     path already exists
//     in the current tree and the value is not nil
func (n *Node[T]) Merge(trees ...*Node[T]) {
	for _, tree := range trees {
		n.merge(tree)
	}
}

func (n *Node[T]) merge(tree *Node[T]) {
	// Check if the tree is nil or if the current node's value is not nil
	// if so, return
	if tree == nil || n.Value != nil {
		return
	}

	// Check if the tree has a value, if so copy it
	if tree.Value != nil {
		n.Value = tree.Value
	}

	// Check if the tree has sub-paths
	if len(tree.SubPaths) > 0 {
		// Check if the current node has sub-paths
		if len(n.SubPaths) == 0 {
			// Set the sub-paths
			n.SubPaths = tree.SubPaths
		} else {
			// Merge the sub-paths
			for _, subPath := range tree.SubPaths {

				// Check if the sub-path exists in the current tree
				// if not, add it
				n2, ok := n.SubPaths[subPath.Path]
				if !ok {
					n.SubPaths[subPath.Path] = subPath
				} else {
					n2.merge(&subPath)
				}

			}
		}
	}
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
func (n *Node[T]) Visit(options VisitorOptions, visitFn func(node *Node[T])) {
	// Visit the current node
	n.visit(VisitorOptions{}, visitFn)
}

// visit recursively visits each node in the tree and calls the visitor function in a depth first manner
//
//   - It optionally visits the paths in a deterministic order
func (n *Node[T]) visit(options VisitorOptions, visitFn func(node *Node[T])) {
	visitFn(n)

	if options.Deterministic {
		keys := lo.Keys(n.SubPaths)
		slices.Sort(keys)

		for _, key := range keys {
			subPath := n.SubPaths[key]
			visitFn(&subPath)
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
