package Red_Black_Tree

import "fmt"

type Color bool

const (
	RED   Color = false
	BLACK Color = true
)

type RedBlackTree struct {
	Root *Node
}

type Node struct {
	value  int
	color  Color
	left   *Node
	right  *Node
	parent *Node
	Height int
}

func (node *Node) NewNode(Value int, key int) *Node {
	return &Node{value: Value}
}
func FindHeight(node *Node) int {
	if node == nil {
		return -1
	}
	lefth := FindHeight(node.left)
	righth := FindHeight(node.right)
	if lefth > righth {
		return lefth + 1
	} else {
		return righth + 1
	}
}

func (tree *RedBlackTree) CreateTree(root *Node, arr []int) {
	//tree := new(RedBlackTree)
	if root == nil {
		tree.Insert(arr[0])
	}
	for i := 1; i < len(arr); i++ {
		//node := root.NewNode(arr[i], arr[i])
		tree.Insert(arr[i])

	}

}

func minValueNode(node *Node) *Node {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}

func (tree *RedBlackTree) Search(key int) *Node {
	return search(tree.Root, key)
}

func search(node *Node, value int) *Node {
	if node == nil || value == node.value {
		return node
	}

	if value < node.value {
		return search(node.left, value)
	}
	return search(node.right, value)
}

func (tree *RedBlackTree) leftRotate(node *Node) {
	child := node.right
	node.right = child.left
	if node.right != nil {
		node.right.parent = node
	}
	child.parent = node.parent

	if node.parent == nil {
		tree.Root = child
	} else if node == node.parent.left {
		node.parent.left = child
	} else {
		node.parent.right = child
	}
	child.left = node
	node.parent = child
}

func (tree *RedBlackTree) rightRotate(node *Node) {
	child := node.left
	node.left = child.right
	if node.left != nil {
		node.left.parent = node
	}
	child.parent = node.parent

	if node.parent == nil {
		tree.Root = child
	} else if node == node.parent.left {
		node.parent.left = child
	} else {
		node.parent.right = child
	}
	child.right = node
	node.parent = child

}

func flipColors(node *Node) {
	node.color = RED
	node.right.color = BLACK
	node.left.color = BLACK
}

func (tree *RedBlackTree) Insert(val int) {
	node := &Node{value: val}
	var parent *Node
	current := tree.Root
	for current != nil {
		parent = current
		if node.value < current.value {
			current = current.left
		} else {
			current = current.right
		}
	}
	node.parent = parent
	if parent == nil {
		tree.Root = node
	} else if node.value < parent.value {
		parent.left = node
	} else {
		parent.right = node
	}
	tree.fixInsert(node)
}

func (tree *RedBlackTree) fixInsert2(node *Node) {
	var parent *Node
	var grandparent *Node
	for node != tree.Root && node.color == RED && node.parent.color == RED {
		parent = node.parent
		grandparent = parent.parent
		if parent == grandparent.left {
			uncle := grandparent.right
			if uncle != nil && uncle.color == RED {
				// uncle red
				grandparent.color = RED
				parent.color = BLACK
				uncle.color = BLACK
				node = grandparent
			} else {
				//uncle black
				if node == parent.right {
					tree.leftRotate(parent)
					node = parent
					parent = node.parent
				}
				tree.rightRotate(grandparent)
				swap(parent.color, grandparent.color)
				node = parent
			}
		} else {
			uncle := grandparent.left
			if uncle != nil && uncle.color == RED {
				grandparent.color = RED
				parent.color = BLACK
				uncle.color = BLACK
				node = grandparent
			} else {
				if node == parent.left {
					tree.rightRotate(parent)
					node = parent
					parent = node.parent
				}
				tree.leftRotate(grandparent)
				swap(parent.color, grandparent.color)
				node = parent
			}
		}
	}
	tree.Root.color = BLACK
}

func swap(col1 Color, col2 Color) {
	col1, col2 = col2, col1
}

func (tree *RedBlackTree) fixInsert(node *Node) {
	for node != tree.Root && node.parent.color == RED {
		if node.parent == node.parent.parent.left {
			uncle := node.parent.parent.right
			if uncle != nil && uncle.color == RED {
				// Случай 1: Дядя красный
				node.parent.color = BLACK
				uncle.color = BLACK
				node.parent.parent.color = RED
				node = node.parent.parent
			} else {
				// Случай 2: Дядя чёрный
				if node == node.parent.right {
					// Случай 2a: Узел - правый ребёнок
					node = node.parent
					tree.leftRotate(node)
				}
				// Случай 2b: Узел - левый ребёнок
				node.parent.color = BLACK
				node.parent.parent.color = RED
				tree.rightRotate(node.parent.parent)
			}
		} else {
			uncle := node.parent.parent.left
			if uncle != nil && uncle.color == RED {
				// Случай 1: Дядя красный (симметрично)
				node.parent.color = BLACK
				uncle.color = BLACK
				node.parent.parent.color = RED
				node = node.parent.parent
			} else {
				// Случай 2: Дядя чёрный (симметрично)
				if node == node.parent.left {
					// Случай 2a: Узел - левый ребёнок (симметрично)
					node = node.parent
					tree.rightRotate(node)
				}
				// Случай 2b: Узел - правый ребёнок (симметрично)
				node.parent.color = BLACK
				node.parent.parent.color = RED
				tree.leftRotate(node.parent.parent)
			}
		}
	}
	tree.Root.color = BLACK
}

func (tree *RedBlackTree) transplant(u, v *Node) {
	if u.parent == nil {
		tree.Root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}
func (tree *RedBlackTree) Delete(key int) {
	nodeToDelete := tree.Search(key)
	if nodeToDelete != nil {
		tree.deleteNode(nodeToDelete)
	}
}

func (tree *RedBlackTree) deleteNode(nodeToDelete *Node) {
	var nodeToReplace, replacementChild *Node

	nodeToReplace = nodeToDelete // Узел для удаления

	originalColor := nodeToReplace.color

	if nodeToReplace.left == nil {
		replacementChild = nodeToReplace.right
		tree.transplant(nodeToReplace, nodeToReplace.right)
	} else if nodeToReplace.right == nil {
		replacementChild = nodeToReplace.left
		tree.transplant(nodeToReplace, nodeToReplace.left)
	} else {
		// Узел имеет оба поддерева
		nodeToReplace = minValueNode(nodeToReplace.right)
		originalColor = nodeToReplace.color
		replacementChild = nodeToReplace.right

		if nodeToReplace.parent == nodeToDelete {
			if replacementChild != nil {
				replacementChild.parent = nodeToReplace
			}
		} else {
			tree.transplant(nodeToReplace, nodeToReplace.right)
			nodeToReplace.right = nodeToDelete.right
			nodeToReplace.right.parent = nodeToReplace
		}

		tree.transplant(nodeToDelete, nodeToReplace)
		nodeToReplace.left = nodeToDelete.left
		nodeToReplace.left.parent = nodeToReplace
		nodeToReplace.color = nodeToDelete.color
	}

	if originalColor == BLACK {
		tree.fixDelete(replacementChild)
	}
}

func (tree *RedBlackTree) fixDelete(deletedNode *Node) {
	for deletedNode != tree.Root && deletedNode.color == BLACK {
		if deletedNode == deletedNode.parent.left {
			// deletedNode — левый потомок
			brother := deletedNode.parent.right // brother — брат deletedNode

			if brother.color == RED {
				// Случай 1: brother — красный
				brother.color = BLACK
				deletedNode.parent.color = RED
				tree.leftRotate(deletedNode.parent)
				brother = deletedNode.parent.right // обновляем brother
			}

			// Случай 2: brother — черный и оба его потомка черные
			if brother.left.color == BLACK && brother.right.color == BLACK {
				brother.color = RED
				deletedNode = deletedNode.parent // поднимаемся на уровень выше
			} else {
				// Случай 3: brother — черный и левый потомок brother — красный
				if brother.right.color == BLACK {
					brother.left.color = BLACK
					brother.color = RED
					tree.rightRotate(brother)
					brother = deletedNode.parent.right // обновляем brother
				}

				// Случай 4: brother — черный и правый потомок brother — красный
				brother.color = deletedNode.parent.color
				deletedNode.parent.color = BLACK
				brother.right.color = BLACK
				tree.leftRotate(deletedNode.parent)
				deletedNode = tree.Root // завершаем цикл
			}
		} else {
			// deletedNode — правый потомок
			brother := deletedNode.parent.left // brother — брат deletedNode

			if brother.color == RED {
				// Случай 1: brother — красный
				brother.color = BLACK
				deletedNode.parent.color = RED
				tree.rightRotate(deletedNode.parent)
				brother = deletedNode.parent.left // обновляем brother
			}

			// Случай 2: brother — черный и оба его потомка черные
			if brother.right.color == BLACK && brother.left.color == BLACK {
				brother.color = RED
				deletedNode = deletedNode.parent // поднимаемся на уровень выше
			} else {
				// Случай 3: brother — черный и правый потомок brother — красный
				if brother.left.color == BLACK {
					brother.right.color = BLACK
					brother.color = RED
					tree.leftRotate(brother)
					brother = deletedNode.parent.left // обновляем brother
				}

				// Случай 4: brother — черный и левый потомок brother — красный
				brother.color = deletedNode.parent.color
				deletedNode.parent.color = BLACK
				brother.left.color = BLACK
				tree.rightRotate(deletedNode.parent)
				deletedNode = tree.Root // завершаем цикл
			}
		}
	}
	deletedNode.color = BLACK // Корень должен быть черным
}

func PrintTree2(root *Node) {
	printSubTree2(root, 0)
}
func printSubTree2(node *Node, space int) {
	if node == nil {
		return
	}
	space += 2
	printSubTree2(node.right, space)
	fmt.Println("")
	for i := 2; i < space; i++ {
		fmt.Print("  ")
	}
	if node.color == RED {
		fmt.Println(node.value, "RED")
	} else {
		fmt.Println(node.value, "BLACK")
	}
	//fmt.Println(node.value)

	printSubTree2(node.left, space)
}

func InorderTraversal(node *Node) {
	if node != nil {
		InorderTraversal(node.left)
		fmt.Printf("%d\t", node.value)
		InorderTraversal(node.right)
	}
}

func Test(arr []int) {
	tree := new(RedBlackTree)
	//var root *Node
	tree.CreateTree(nil, arr)
	//tree.Insert(3)

	//tree.Delete(6)
	PrintTree2(tree.Root)
	InorderTraversal(tree.Root)
}
