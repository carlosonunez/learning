package data_structures

import "testing"

var treeWithoutChildren Tree = Tree{
	item: 3,
}

var treeWithLeftOnly Tree = Tree{
	item: 3,
	left: Tree{
		item: 2,
	},
}

var treeWithRightOnly Tree = Tree{
	item: 3,
	left: nil,
	right: Tree{
		item:  "4",
		left:  nil,
		right: nil,
	},
}

var treeWithSingleLevel Tree = Tree{
	item: 3,
	left: Tree{
		item:  "2",
		left:  nil,
		right: nil,
	},
	right: Tree{
		item:  "4",
		left:  nil,
		right: nil,
	},
}

var treeWithTwoLevels Tree = Tree{
	item: 3,
	left: Tree{
		item: 2,
		left: Tree{
			item:  "1",
			left:  nil,
			right: nil,
		},
		right: Tree{
			item:  "2",
			left:  nil,
			right: nil,
		},
	},
	right: Tree{
		item: 4,
		left: Tree{
			item:  "4",
			left:  nil,
			right: nil,
		},
		right: Tree{
			item:  "6",
			left:  nil,
			right: nil,
		},
	},
}

var unbalancedTreeWithTwoLevels Tree = Tree{
	item: 3,
	left: Tree{
		item: 2,
		left: Tree{
			item:  "4",
			left:  nil,
			right: nil,
		},
		right: Tree{
			item:  "2",
			left:  nil,
			right: nil,
		},
	},
	right: Tree{
		item: 4,
		left: Tree{
			item:  "4",
			left:  nil,
			right: nil,
		},
		right: Tree{
			item:  "1",
			left:  nil,
			right: nil,
		},
	},
}

func TestPrintTreeRootOnly(t *testing.T) {
	printTree(treeWithoutChildren)
}
