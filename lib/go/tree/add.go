package tree

import (
	"log/slog"
	"strings"

	lop "github.com/samber/lo/parallel"
)

func (n *Node[T]) AddPath(path string, value *T) (err error) {
	// Split the path
	slog.Info("StartAdd", "path", path)

	split := strings.Split(path, ".")

	base := split[0]

	subNode, ok := n.SubPaths[base]
	if !ok {
		subNode = &Node[T]{
			FullPath:      base,
			Path:          base,
			SubPaths:      make(map[string]*Node[T], DefaultSize),
			AllowSubPaths: true,
			Value:         nil,
		}
	}
	n.SubPaths[base] = subNode

	return subNode.addPath(base, split, value)
}

// addPath adds a path to the tree, if the path is a nested path, it will
// recursively add the path to the tree.
//
//   - If the node does not allow sub-paths, it will return an error if it tries
//     to add a sub-path
func (n *Node[T]) addPath(pathPrefix string, split []string, value *T) (err error) {
	// a := split[0]
	// if len(split) == 1 {
	// b := split[1]

	// spew.Dump(n)

	slog.Info("addPath", "split", split)

	// Check if the path is empty
	if len(split) == 1 {

		n.Value = value
		return
	}

	pathPrefix = pathPrefix + "." + split[1]

	if !n.AllowSubPaths {
		return n.error().Errorf("cannot add sub-path to node that does not allow sub-paths")
	}
	if n.SubPaths == nil {
		n.SubPaths = make(map[string]*Node[T], DefaultSize)
	}

	// Check if the sub-path exists
	subPath, ok := n.SubPaths[split[1]]
	if !ok {
		// Create the sub-path
		subPath = &Node[T]{
			FullPath:      pathPrefix,
			Path:          split[1],
			SubPaths:      make(map[string]*Node[T], DefaultSize),
			AllowSubPaths: true,
			Value:         nil,
		}
	}

	// Recurse to the sub-path
	err = subPath.addPath(pathPrefix, split[1:], value)
	if err != nil {
		return err
	}
	n.SubPaths[split[1]] = subPath

	return
}

// ─────────────────────────────────────────────────────────────────────────────

//

type AddEntry[T any] struct {
	Path  string
	Value *T
}

// BulkAddPaths - Adds a list of paths to the tree at the same time.
//
//   - It does this by splitting up the paths into nested groups and then adding
//     making the map entries for each group synchronously. Then it will add the
//     values asynchronously followed by each groups sub-paths synchrony.
//   - As we are writing synchronously to the map, we do not need to worry about
//     attempts to invalidate the map while we are writing to its sub-paths.
func (n *Node[T]) BulkAddPaths(entries []AddEntry[T]) (err error) {
	// Split the path
	slog.Info("StartBulkAdd", "path", entries)

	groups := lop.GroupBy(entries, func(entry AddEntry[T]) string {
		return strings.SplitN(entry.Path, ".", 2)[0]
	})

	// Ensure the root node has a sub-path map
	if n == nil {
		return n.error().Errorf("cannot add paths to nil node")
	} else if n.SubPaths == nil {
		n.SubPaths = make(map[string]*Node[T], len(groups))
	}

	// Add the groups synchronously
	for group := range groups {
		n.addGroupMap("", group)
	}

	// Add the values asynchronously
	for group, groupEntries := range groups {

		g, ok := n.SubPaths[group]
		if !ok {
			return n.error().Errorf("cannot find group %s", group)
		}
		go g.addGroupEntries(group, groupEntries)
	}

	return nil
}

// addGroupMap adds a  map entry for a group of paths to the tree, meant to be
// called synchronously.
func (n *Node[T]) addGroupMap(prefix, group string) {
	subNode, ok := n.SubPaths[group]
	if !ok {
		subNode = &Node[T]{
			FullPath:      prefix + group,
			Path:          group,
			SubPaths:      make(map[string]*Node[T], DefaultSize),
			AllowSubPaths: true,
			Value:         nil,
		}
	}
	n.SubPaths[group] = subNode
}

// addGroupEntries adds a subentry for each member of the group
func (n *Node[T]) addGroupEntries(prefix string, entries []AddEntry[T]) {

	slog.Info("Making The Maps Entries", "prefix", prefix, "entries", entries)
	for _, entry := range entries {
		n.addGroupMap(prefix, entry.Path)
	}

	slog.Info("Adding Entries", "prefix", prefix, "entries", entries)

}
