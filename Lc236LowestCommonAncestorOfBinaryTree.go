package LeetcodeExercise

//TreeNode  definition for a binary tree node.
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	hits := 0
	return afterOrderTravel(root, p, q, &hits)
}

func afterOrderTravel(root, p, q *TreeNode, hits *int) *TreeNode {
	if root.left != nil {
		ancestor := afterOrderTravel(root.left, p, q, hits)
		if ancestor != nil {
			return ancestor
		}
	}
	if root.right != nil {
		ancestor := afterOrderTravel(root.right, p, q, hits)
		if ancestor != nil {
			return ancestor
		}
	}
	if root.val == p.val || root.val == q.val {
		*hits++
	}
	if *hits == 2 {
		ancestor := root //allow a node as it's own decendant
		return ancestor
	}
	return nil //continue find
}
