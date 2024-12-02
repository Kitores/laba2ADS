package Red_Black_Tree

type Color bool

const (
	red   Color = false
	black Color = true
)

type Node struct {
	key    int
	value  int
	color  Color
	left   *Node
	right  *Node
	parent *Node
}

func (node *Node) NewNode(Value int, key int) *Node {
	return &Node{value: Value, key: key}
}

func (node *Node) isRed() bool {
	return node != nil && node.color == red
}

func (node *Node) isBlack() bool {
	return node != nil && node.color == black
}

func rotateLeft(node *Node) *Node {
	right := node.right
	node.right = right.left
	if right.left == nil {
	}
	return
}

func flipColors(node *Node) {
	node.color = red
	node.right.color = black
	node.left.color = black
}

func (node *Node) Insert(Value int, key int) *Node {
	if node == nil {
		return node.NewNode(Value, key)
	}
	if node.right.isRed() {
		node.left.Insert(Value, key)
	}
	if node.isBlack() {
		node.right.Insert(Value, key)
	}
	return node.NewNode(Value, key)
}

func (root *Node) Insert(value int) {
	newNode := &Node{value: value, color: red}
	root.bstInsert()
	root.fixInsert()
}

func Insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{Value: value}
	}
	if value < root.Value {
		root.Left = Insert(root.Left, value)
	} else if value > root.Value {
		root.Right = Insert(root.Right, value)
	} else {
		root.Value = value // Обновляем значение, если ключ уже существует
	}
	return root
}
