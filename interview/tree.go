package interview

//TreeNode  definition for a binary tree node.
type TreeNode struct {
	V int       //value of node
	L *TreeNode //left child
	R *TreeNode //right child
}

//TreeNodeDoubleLinked  definition for a binary tree node who has an extra link to its parent node
type TreeNodeDoubleLinked struct {
	V int                   //value of node
	L *TreeNodeDoubleLinked //left child
	R *TreeNodeDoubleLinked //right child
	P *TreeNodeDoubleLinked //parent node
}

//BuildBST builds a Binary Search Tree
func BuildBST(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	root := &TreeNode{V: vals[0]}
	for _, v := range vals[1:] {
		insertBSTNode(root, v)
	}
	return root
}

func findMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.L == nil {
		return root
	}
	return findMin(root.L)
}

func findNode(root *TreeNode, v int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.V == v {
		return root
	}

	if v < root.V {
		return findNode(root.L, v)
	}

	return findNode(root.R, v)
}

func insertBSTNode(root *TreeNode, v int) {
	if v < root.V {
		if root.L == nil {
			root.L = &TreeNode{V: v}
			return
		}

		insertBSTNode(root.L, v)
	} else {
		if root.R == nil {
			root.R = &TreeNode{V: v}
			return
		}

		insertBSTNode(root.R, v)
	}
}

//BuildBSTDoubleLinked builds a Binary Search Tree double linked
func BuildBSTDoubleLinked(vals []int) *TreeNodeDoubleLinked {
	if len(vals) == 0 {
		return nil
	}
	root := &TreeNodeDoubleLinked{V: vals[0]}
	for _, v := range vals[1:] {
		insertBSTNodeDoubleLinked(root, v)
	}
	return root
}

func findNodeDoubleLinked(root *TreeNodeDoubleLinked, v int) *TreeNodeDoubleLinked {
	if root == nil {
		return nil
	}

	if root.V == v {
		return root
	}

	if v < root.V {
		return findNodeDoubleLinked(root.L, v)
	}

	return findNodeDoubleLinked(root.R, v)
}

func findMinDoubleLinked(root *TreeNodeDoubleLinked) *TreeNodeDoubleLinked {
	if root == nil {
		return nil
	}

	if root.L == nil {
		return root
	}
	return findMinDoubleLinked(root.L)
}

func insertBSTNodeDoubleLinked(root *TreeNodeDoubleLinked, v int) {
	if v < root.V {
		if root.L == nil {
			root.L = &TreeNodeDoubleLinked{V: v, P: root}
			return
		}

		insertBSTNodeDoubleLinked(root.L, v)
	} else {
		if root.R == nil {
			root.R = &TreeNodeDoubleLinked{V: v, P: root}
			return
		}

		insertBSTNodeDoubleLinked(root.R, v)
	}
}
