package data_structures

import "fmt"

type Tree struct {
	item  int32
	left  Tree
	right Tree
}

func (tree *Tree, depth *int, direction *string) InternalPrintTree() {
	if depth == 0 {
		message := fmt.Sprintf("ROOT: %s", tree.item)
	} else {
		message := fmt.Sprintf("[Depth: %d, Direction: %s] %s", depth, direction, tree.item)
	}
	fmt.Println(message)
	if tree.left != nil {
		InternalPrintTree(tree.left, depth+1, "left")
	} else if tree.right != nil {
		InternalPrintTree(tree.right, depth+1, "right")
	}
}

func (tree *Tree) printTree() {
	InternalPrintTree(tree, 0, "root")
}
