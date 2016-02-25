package data_structures

import "testing"

var treeWithoutChildren Tree = Tree{
	item: 3,
}

var treeWithLeftOnly Tree = Tree{
	item: 3,
	left: &Tree{
		item: 2,
	},
}

var treeWithRightOnly Tree = Tree{
	item: 3,
	right: &Tree{
		item: 4,
	},
}

var treeWithSingleLevel Tree = Tree{
	item: 3,
	left: &Tree{
		item: 2,
	},
	right: &Tree{
		item: 4,
	},
}

var treeWithTwoLevels Tree = Tree{
	item: 3,
	left: &Tree{
		item: 2,
		left: &Tree{
			item: 1,
		},
		right: &Tree{
			item: 2,
		},
	},
	right: Tree{
		item: 4,
		left: &Tree{
			item: 4,
		},
		right: &Tree{
			item: 6,
		},
	},
}

var unbalancedTreeWithTwoLevels Tree = Tree{
	item: 3,
	left: &Tree{
		item: 2,
		left: &Tree{
			item: 4,
		},
		right: &Tree{
			item: 2,
		},
	},
	right: &Tree{
		item: 4,
		left: &Tree{
			item: 4,
		},
		right: &Tree{
			item: 1,
		},
	},
}

func TestPrintTreeRootOnly(t *testing.T) {
	printTree(&treeWithoutChildren)
}
