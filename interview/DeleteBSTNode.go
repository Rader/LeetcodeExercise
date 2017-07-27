package interview

/*
Delete a node from a BST(Binary Search Tree)
*/

func delBSTNodeDoubleLinked(root, n *TreeNodeDoubleLinked) *TreeNodeDoubleLinked {
	if n.R != nil {
		min := findMinDoubleLinked(n.R)
		n.V = min.V //swap node N and min
		n = min     //now min is the node to delete
	}
	//delete node
	if n.P == nil { // we are deleting the root, and apparently it has no children
		return nil
	}

	//check whether n is the left or right child
	if n.P.L == n {
		n.P.L = nil
		n.P = nil
	} else {
		n.P.R = nil
		n.P = nil
	}
	return root
}
