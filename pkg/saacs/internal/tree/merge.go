package tree

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
					n2.merge(subPath)
				}

			}
		}
	}
}
