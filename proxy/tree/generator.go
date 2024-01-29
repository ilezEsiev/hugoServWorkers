package tree

type Node struct {
	Key    int
	Height int
	Left   *Node
	Right  *Node
}

type AVLTree struct {
	Root *Node
}

func (t *AVLTree) Insert(value int) {
	t.Root = insert(t.Root, value)
}

func NewAVLTree() *AVLTree {
	outTree := &AVLTree{}
	for i := 0; i < 5; i++ {
		outTree.Insert(i)
	}
	return outTree
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.Height
}

func rebalance(node *Node) *Node {
	updateHeight(node)
	balance := getBalance(node)

	if balance > 1 {
		if getBalance(node.Left) < 0 {
			node.Left = leftRotate(node.Left)
		}
		return rightRotate(node)
	}
	if balance < -1 {
		if getBalance(node.Right) > 0 {
			node.Right = rightRotate(node.Right)
		}
		return leftRotate(node)
	}
	return node
}

func insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{Key: value, Height: 1}
	}
	if value < node.Key {
		node.Left = insert(node.Left, value)
	} else if value > node.Key {
		node.Right = insert(node.Right, value)
	}
	return rebalance(node)
}

func updateHeight(node *Node) {
	node.Height = max(height(node.Left), height(node.Right)) + 1
}

func getBalance(node *Node) int {
	return height(node.Left) - height(node.Right)
}

func leftRotate(x *Node) *Node {
	y := x.Right
	T2 := y.Left
	y.Left = x
	x.Right = T2
	updateHeight(x)
	updateHeight(y)
	return y
}

func rightRotate(y *Node) *Node {
	x := y.Left
	T2 := x.Right
	x.Right = y
	y.Left = T2
	updateHeight(y)
	updateHeight(x)
	return x
}
