package bst

type BST struct {
	Root *Node
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (bst *BST) Find(node Node) bool {

	return bst.find(bst.Root, node.Val)
}

func (bst *BST) find(cur *Node, target int) bool {
	if cur == nil {
		return false
	}

	if target == cur.Val {
		return true
	}

	if target > cur.Val {
		return bst.find(cur.Right, target)
	} else {
		return bst.find(cur.Left, target)
	}
}

func (bst *BST) Insert(node *Node) bool {
	return bst.insert(bst.Root, node)
}

func (bst *BST) insert(cur *Node, node *Node) bool {
	if cur == nil {
		return false
	}

	if node.Val > cur.Val {
		if cur.Right == nil {
			cur.Right = node
			return true
		} else {
			return bst.insert(cur.Right, node)
		}
	} else {
		if cur.Left == nil {
			cur.Left = node
			return true
		} else {
			return bst.insert(cur.Left, node)
		}
	}
}

func (bst *BST) Delete(node *Node) bool {
	var deleted bool
	bst.Root, deleted = bst.delete(bst.Root, node)
	return deleted

}

func (bst *BST) delete(cur *Node, node *Node) (*Node, bool) {
	if cur == nil {
		return nil, false
	}

	if node.Val < cur.Val {
		cur.Left, _ = bst.delete(cur.Left, node)
	} else if node.Val > cur.Val {
		cur.Right, _ = bst.delete(cur.Right, node)
	} else {
		// Found the node
		if cur.Left == nil && cur.Right == nil {
			// case 1: leaf
			return nil, true
		}
		if cur.Left == nil {
			// case 2: one child (right)
			return cur.Right, true
		}
		if cur.Right == nil {
			// case 2: one child (left)
			return cur.Left, true
		}

		// case 3: two children
		successor := findMin(cur.Right)
		cur.Val = successor.Val
		cur.Right, _ = bst.delete(cur.Right, successor)
	}

	return cur, true
}

func findMin(n *Node) *Node {
	for n.Left != nil {
		n = n.Left
	}
	return n
}
