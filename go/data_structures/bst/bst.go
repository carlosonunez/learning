package data_structures

import "fmt"

type BinarySearchTree struct {
	item  int32
	left  *BinarySearchTree
	right *BinarySearchTree
}

func (tree *BinarySearchTree) InternalPrintTree(depth int, direction string) {
	// Print the current node...
	if depth == 0 {
		fmt.Printf("ROOT: %d\n", tree.item)
	} else {
		fmt.Printf("[Depth: %d, Direction: %s] %d\n", depth, direction, tree.item)
	}

	// ...then recurse into the left or right branch!
	if tree.left != nil {
		tree.left.InternalPrintTree(depth+1, "left")
	}
	if tree.right != nil {
		tree.right.InternalPrintTree(depth+1, "right")
	}
}

// printTree prints the tree starting at its root. Its internal method above
// takes care of reaching its leaves.
func (tree *BinarySearchTree) printTree() {
	tree.InternalPrintTree(0, "root")
}
