package AVL_Tree

import (
	"fmt"
)

type AVLTree struct {
	root *Node
}

type Node struct {
	Key    int
	Value  int
	Height int
	Left   *Node
	Right  *Node
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
	//node.Height = 1 + int(math.Max(float64(node.Left.GetHeight()), float64(node.Right.GetHeight())))
}

func (node *Node) min() *Node {
	if node.Left == nil {
		return node
	}
	return node.Left.min()
}

func (node *Node) rotateRight() *Node {
	newRoot := node.Left
	node.Left = newRoot.Right
	newRoot.Right = node
	node.updateHeight()
	newRoot.updateHeight()
	return newRoot
}
func (node *Node) rotateLeft() *Node {
	newRoot := node.Right
	node.Right = newRoot.Left
	newRoot.Left = node
	node.updateHeight()
	newRoot.updateHeight()
	return newRoot
}

func (node *Node) BalanceFactor() int {
	if node == nil {
		return 0
	}
	return node.Right.GetHeight() - node.Left.GetHeight()
}

func balance(node *Node) *Node {
	node.updateHeight()
	balanceFlag := node.BalanceFactor()

	if balanceFlag > 1 {
		if node.Right.BalanceFactor() < 0 {
			node.Right = node.Right.rotateRight()
		}
		return node.rotateLeft()
	} else if balanceFlag < -1 {
		if node.Left.BalanceFactor() > 0 {
			node.Left = node.Left.rotateLeft()
		}
		return node.rotateRight()
	}
	return node
}

func CreateTree(root *Node, arr []int) *Node {
	if root == nil {
		root = root.Insert(0, arr[0])
	}
	//fmt.Println(root)
	for i := 1; i < len(arr); i++ {
		root = root.Insert(i, arr[i])
		//fmt.Println(arr[i], root)
	}
	return root
}

func (node *Node) Insert(key int, newVal int) *Node {
	if node == nil {
		node = &Node{Key: key, Value: newVal}
		return node
	}
	if newVal < node.Value { //Insert in left subtree
		node.Left = node.Left.Insert(key, newVal)
	} else if newVal > node.Value { //Insert in right subtree
		node.Right = node.Right.Insert(key, newVal)
	} else {
		node.Value = newVal // Update exciting node
		return node
	}
	return balance(node)
}

func Balance(root *Node) *Node {
	balance(root)
	return root
}
func (node *Node) Remove(value int) *Node {
	if node == nil {
		return nil
	}
	if value < node.Value {
		node.Left = node.Left.Remove(value)
	} else if value > node.Value {
		node.Right = node.Right.Remove(value)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		} else {
			minNode := node.Right.min()
			node.Value = minNode.Value
			node.Right = node.Right.Remove(minNode.Value)
		}
	}
	return balance(node)
}

func (node *Node) Search(value int) bool {
	if node == nil {
		return false
	}
	if value == node.Value {
		return true
	} else if value < node.Value {
		return node.Left.Search(value)
	} else {
		return node.Right.Search(value)
	}
}

func InorderTraversal(node *Node) {
	if node != nil {
		InorderTraversal(node.Left)
		fmt.Printf("%d\t", node.Value)
		InorderTraversal(node.Right)
	}
}

func Traversal(root *Node) {
	if root != nil {
		Traversal(root.Left)
		fmt.Printf("%d\t", root.Value)
		Traversal(root.Right)
	}
}

func PrintTree(root *Node) {
	printSubTree(root, 0)
}
func printSubTree(node *Node, space int) {
	if node == nil {
		return
	}
	space += 2
	//fmt.Printf("%v\t   ---- Val %d, Key %d", node, node.Value, node.Key)
	printSubTree(node.Right, space)
	fmt.Println()
	for i := 2; i < space; i++ {
		fmt.Print("             ")
	}
	fmt.Println(node.Value)
	printSubTree(node.Left, space)
}

func PrintTree2(root *Node) {
	printSubTree2(root, 0)
}
func printSubTree2(node *Node, space int) {
	if node == nil {
		return
	}
	space += 2
	printSubTree2(node.Right, space)
	fmt.Println("")
	for i := 2; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Println(node.Value)

	printSubTree2(node.Left, space)
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

func Test(arr []int) {
	var node *Node
	//var root *Node
	node = CreateTree(nil, arr)
	//tree.Insert(3)

	//tree.Delete(6)
	PrintTree2(node)
	InorderTraversal(node)
	fmt.Println(FindHeight(node))
}
