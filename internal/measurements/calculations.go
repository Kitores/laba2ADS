package measurements

import (
	AVL_Tree "TreesLaba2/internal/tree/AVL-Tree"
	"TreesLaba2/internal/tree/BinarySearchTree"
	Red_Black_Tree "TreesLaba2/internal/tree/Red-Black-Tree"
	"TreesLaba2/pkg"
	"fmt"
)

func Calculate(typeOfTree string, seed int64) (arrX []float64, arrY []int) {
	var Height int

	//keysQuantities := []int{20000, 40000, 60000, 80000, 100000}

	//keysQuantities := []int{10, 20, 30, 40, 50, 70, 100, 200, 300}
	for quantity := 1000; quantity < 100000; quantity += 1000 {
		//for _, quantity := range keysQuantities {

		//fmt.Println(quantity)
		switch typeOfTree {
		case "BST":
			//arr := pkg.IncreasingArray(quantity)
			arr := pkg.RandArray(quantity, seed)
			Height = BSTcalculate(arr)
		case "AVL":
			//arr := pkg.RandArray(quantity, seed)
			arr := pkg.IncreasingArray(quantity)
			Height = AVLcalculate(arr)
		case "RBT":
			arr := pkg.IncreasingArray(quantity)
			Height = RBTcalculate(arr)
		}
		arrY = append(arrY, Height)
		arrX = append(arrX, float64(quantity))
	}
	return arrX, arrY
}

func BSTcalculate(arr []int) int {
	var root *BinarySearchTree.Node
	root = BinarySearchTree.CreateTree(root, arr)
	//BinarySearchTree.PrintTree(root)
	//return root.GetHeight() + 1
	return BinarySearchTree.FindHeight(root)
}

func AVLcalculate(arr []int) int {
	var root *AVL_Tree.Node
	root = AVL_Tree.CreateTree(root, arr)
	//AVL_Tree.PrintTree2(root)
	//return root.GetHeight() + 1
	fmt.Println(root.GetHeight(), " ", AVL_Tree.FindHeight(root))
	return AVL_Tree.FindHeight(root)
}

func RBTcalculate(arr []int) int {
	tree := new(Red_Black_Tree.RedBlackTree)
	tree.CreateTree(nil, arr)
	//AVL_Tree.PrintTree2(root)
	//return root.GetHeight() + 1
	fmt.Println(Red_Black_Tree.FindHeight(tree.Root))
	return Red_Black_Tree.FindHeight(tree.Root)

	//TODO:
}
