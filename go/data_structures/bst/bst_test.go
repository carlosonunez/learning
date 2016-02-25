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

func TestPrintTreeRootOnly(t *testing.T) {
	treeWithoutChildren.printTree()
}

func TestPrintTreeSingleLevel(t *testing.T) {
	treeWithSingleLevel.printTree()
}

func TestPrintTreeMultiLevel(t *testing.T) {
	treeWithTwoLevels.printTree()
}
