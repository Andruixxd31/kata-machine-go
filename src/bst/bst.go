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

	if target > cur.Val {
		return bst.find(cur.Right, target)
	} else if target < cur.Val {
		return bst.find(cur.Left, target)
	} else {
		return true
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
