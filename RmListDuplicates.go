package interview

/** It's an interview question asked by Alibaba.
Q: There is an ordered single linked list, please remove the all the duplicates.
Example: 3 2 2 1 -> 3 1
*/

type Node struct {
	V    int
	Next *Node
}

//RmListDuplicates will remove any element if it appears more than once in the linked list.
//algorithm: go through the list, find non duplicate elment, append to the new list, then find next non duplicate element
func RmListDuplicates(n *Node) *Node {
	h := new(Node)
	tail := h

	for {
		if n == nil {
			break //reach list end
		}
		p := n.Next
		dups := 0 //count of duplicates of the same value
		for {
			if p == nil || p.V != n.V {
				break //reach list end
			}

			dups++
			p = p.Next //forward p, check next element
		}
		if dups == 0 {
			n.Next = nil  //don't take nodes after n
			tail.Next = n //append the non duplicate element
			tail = tail.Next
		}
		n = p
	}

	return h.Next
}
