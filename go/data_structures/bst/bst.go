package data_structures

import "fmt"

type Tree struct {
	item  int32
	left  *Tree
	right *Tree
}

func (tree *Tree) InternalPrintTree(depth int, direction string) {
	// Print the current node...
	if depth == 0 {
		fmt.Println("ROOT: %s", tree.item)
	} else {
		fmt.Println("[Depth: %d, Direction: %s] %s", depth, direction, tree.item)
	}

	// ...then recurse into the left or right branch!
	if tree.left != nil {
		InternalPrintTree(&tree.left, depth+1, "left")
	} else if tree.right != nil {
		InternalPrintTree(&tree.right, depth+1, "right")
	}
}

func (tree *Tree) printTree() {
	InternalPrintTree(&tree, 0, "root")
}
