package BinarySearchTree

import (
	"fmt"
	"math"
	"strings"
)

type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Height int
}

// insert вставляет новый узел в бинарное дерево поиска
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
	root.updateHeight()
	return root
}
func CreateTree(root *Node, arr []int) *Node {
	root = Insert(nil, arr[0])
	for i := 1; i < len(arr); i++ {
		root = Insert(root, arr[i])
	}
	return root
}

func getSuccesor(node *Node) *Node {
	node = node.Left
	for node != nil && node.Right != nil {
		node = node.Right
	}
	return node
}

func (node *Node) Delete(key int) *Node {
	if node == nil {
		return nil
	} else if key < node.Value {
		node.Left = node.Left.Delete(key)
	} else if key > node.Value {
		node.Right = node.Right.Delete(key)
	} else {
		if node.Left == nil {
			tmp := node.Right
			node = nil
			return tmp
		}
		if node.Right == nil {
			tmp := node.Left
			node = nil
			return tmp
		}
		successor := getSuccesor(node)
		node.Value = successor.Value
		successor.Right = node.Delete(successor.Value)
	}
	return node
}

func (n *Node) Visualize(level int) string {
	if n == nil {
		return ""
	}

	indent := strings.Repeat(" ", level)
	leftStr := n.Left.Visualize(level + 1)
	rightStr := n.Right.Visualize(level + 1)

	// Выравнивание поддеревьев для более симметричного вида
	maxLen := int(math.Max(float64(len(leftStr)), float64(len(rightStr))))
	leftStr = fmt.Sprintf("%-*s", maxLen, leftStr)
	rightStr = fmt.Sprintf("%-*s", maxLen, rightStr)

	return fmt.Sprintf("%s%d\n%s%s\n%s%s", indent, n.Value, indent, leftStr, indent, rightStr)
}

func PreorderTraversal(node *Node) {
	if node != nil {
		fmt.Print(node.Value, " ")
		PreorderTraversal(node.Left)
		PreorderTraversal(node.Right)
	}
}

func InorderTraversal(node *Node) {
	if node != nil {
		InorderTraversal(node.Left)
		fmt.Printf("%d\t", node.Value)
		InorderTraversal(node.Right)
	}
}

func PostorderTraversal(node *Node) {
	if node != nil {
		PostorderTraversal(node.Left)
		PostorderTraversal(node.Right)
		fmt.Print(node.Value, " ")
	}
}

func (node *Node) Search(val int) bool {
	if node == nil {
		return false
	}
	if node.Value == val {
		return true
	}

	if node.Value > val {
		node = node.Left
		return node.Search(val)
	}
	node = node.Right
	return node.Search(val)
}
func PrintTree(root *Node) {
	printSubTree(root, 0)
}
func printSubTree(node *Node, space int) {
	if node == nil {
		return
	}
	space += 2
	printSubTree(node.Right, space)
	fmt.Println()
	for i := 2; i < space; i++ {
		fmt.Print("-")
	}
	fmt.Println(node.Value)
	printSubTree(node.Left, space)
}

func (node *Node) GetHeight() int {
	if node == nil {
		return -1
	}
	return node.Height
}
func (node *Node) updateHeight() {
	if node == nil {
		return
	}
	if node.Left.GetHeight() > node.Right.GetHeight() {
		node.Height = node.Left.GetHeight() + 1
	} else {
		node.Height = node.Right.GetHeight() + 1
	}

}

func FindHeight(node *Node) int {
	if node == nil {
		return -1
	}
	lefth := FindHeight(node.Left)
	righth := FindHeight(node.Right)
	if lefth > righth {
		return lefth + 1
	} else {
		return righth + 1
	}
}

//TODO: Доделать Delete()
