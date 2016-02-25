package data_structures

import (
	"testing"
)

var treeWithoutChildren BinarySearchTree = BinarySearchTree{
	item: 3,
}

var treeWithSingleLevel BinarySearchTree = BinarySearchTree{
	item: 3,
	left: &BinarySearchTree{
		item: 2,
	},
	right: &BinarySearchTree{
		item: 4,
	},
}

var treeWithTwoLevels BinarySearchTree = BinarySearchTree{
	item: 3,
	left: &BinarySearchTree{
		item: 2,
		left: &BinarySearchTree{
			item: 1,
		},
		right: &BinarySearchTree{
			item: 2,
		},
	},
	right: &BinarySearchTree{
		item: 4,
		left: &BinarySearchTree{
			item: 4,
		},
		right: &BinarySearchTree{
			item: 6,
		},
	},
}

var unbalancedTreeWithTwoLevels BinarySearchTree = BinarySearchTree{
	item: 3,
	left: &BinarySearchTree{
		item: 2,
		left: &BinarySearchTree{
			item: 4,
		},
		right: &BinarySearchTree{
			item: 2,
		},
	},
	right: &BinarySearchTree{
		item: 4,
		left: &BinarySearchTree{
			item: 4,
		},
		right: &BinarySearchTree{
			item: 1,
		},
	},
}

// Copy tests
// ==============
func TestCloneOnEmptyTree() {
	tree := treeWithoutChildren.clone()
	if tree != treeWithoutChildren {
		t.Errorf("Trees are not equal.")
	}
}

// Printing tests
// ==========================================
func TestPrintTreeRootOnly(t *testing.T) {
	treeWithoutChildren.printTree()
}

func TestPrintTreeSingleLevel(t *testing.T) {
	treeWithSingleLevel.printTree()
}

func TestPrintTreeMultiLevel(t *testing.T) {
	treeWithTwoLevels.printTree()
}

// Insert tests
// =================
func TestInsertIntoEmptyTree(t *testing.T) {
	emptyTree := &BinarySearchTree{}
	emptyTree.insert(3)
	emptyTree.printTree()
}

func TestInsertIntoSingletonTree(t *testing.T) {
	tree := treeWithoutChildren
	tree.insert(6)
	tree.insert(1)
	tree.insert(2)
	tree.insert(1)
	tree.printTree()
}

func TestInsertIntoMultilevelTree(t *testing.T) {
	tree := treeWithTwoLevels
	tree.insert(10)
	tree.insert(-4)
	tree.printTree()
}

// Lookup tests
// =========================
func TestLookupOnEmptyTree(t *testing.T) {
	result := treeWithoutChildren.lookup(10)
	if result {
		t.Errorf("Expected false; got true")
	}
}

func TestLookupOnSingletonTree(t *testing.T) {
	result := treeWithSingleLevel.lookup(3)
	if !result {
		t.Errorf("Expected true, but got false.")
	}

	result = treeWithSingleLevel.lookup(2500)
	if result {
		t.Errorf("Expected false, but got true.")
	}
}

func TestLookupOnMultilevelTree(t *testing.T) {
	result := treeWithSingleLevel.lookup(3)
	if !result {
		t.Errorf("Expected true, but got false.")
	}

	result = treeWithSingleLevel.lookup(6)
	if !result {
		t.Errorf("Expected true, but got false.")
	}

	result = treeWithSingleLevel.lookup(2500)
	if result {
		t.Errorf("Expected false, but got true.")
	}
}
