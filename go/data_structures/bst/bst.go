package data_structures

import (
	"fmt"
)

// A standard binary search tree structure.
type BinarySearchTree struct {
	item  int32
	left  *BinarySearchTree
	right *BinarySearchTree
}

// InternalPrintTree traverses each node of the tree and prints it along
// with its depth and direction relative to their parents.
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

// InternalIsEmpty detects whether the tree is empty.
func (tree *BinarySearchTree) InternalIsEmpty() bool {
	if (tree.item == 0) && (tree.left == nil) && (tree.right == nil) {
		return true
	} else {
		return false
	}
}

// --- public ---

// Creates a copy of a BinarySearchTree struct since Golang does not
// natively support direct copying of structures that contain anything
// other than fundamental types (bools, ints, strings, etc.)
// https://groups.google.com/forum/#!topic/golang-nuts/JYII32waCL4
func (tree *BinarySearchTree) clone() BinarySearchTree {
	newTree := BinarySearchTree{}
	if tree.item != nil {
		newTree.item = tree.item
	}
	if tree.left != nil {
		newTree.left = tree.left.clone()
	}
	if tree.right != nil {
		newTree.right = tree.right.clone()
	}
	return newTree
}

// printTree prints the tree starting at its root. Its internal method above
// takes care of reaching its leaves.
func (tree *BinarySearchTree) printTree() {
	tree.InternalPrintTree(0, "root")
}

// lookup searches for a value within the tree.
// Returns a bool corresponding to whether the value was found or not.
func (tree *BinarySearchTree) lookup(val int32) bool {
	if tree.InternalIsEmpty() {
		return false
	}

	if tree.item == val {
		return true
	} else if (val < tree.item) && (tree.left) != nil {
		tree.left.lookup(val)
	} else if (val > tree.item) && (tree.right) != nil {
		tree.right.lookup(val)
	}

	// If we've gotten to this point, abandon hope.
	return false
}

// add inserts a node into the BST given.
// Lower values are on the left; higher values are on the right.
func (tree *BinarySearchTree) insert(val int32) {
	// The tree is empty if the node we're on has no left or right children.
	if tree.InternalIsEmpty() {
		tree.item = val
	} else if val == tree.item {
		fmt.Println("ERROR: Item already exists in tree.")
	} else if val < tree.item {
		// If the value provided is lower than the current item, go left
		// Otherwise, go right. If the left or right trees are not yet
		// defined, create them now.
		if tree.left == nil {
			tree.left = &BinarySearchTree{
				item: val,
			}
		} else {
			tree.left.insert(val)
		}
	} else if val > tree.item {
		if tree.right == nil {
			tree.right = &BinarySearchTree{
				item: val,
			}
		} else {
			tree.right.insert(val)
		}
	}
}
