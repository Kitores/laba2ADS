package main

import (
	AVL_Tree "TreesLaba2/internal/tree/AVL-Tree"
	"TreesLaba2/pkg"
	"fmt"
)

func main() {
	//arr := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	arr := pkg.RandArray(20, 12)
	fmt.Println(arr)

	//var r = BinarySearchTree.Node{}
	//var root *BinarySearchTree.Node = &r
	//
	//var root *BinarySearchTree.Node
	//root = BinarySearchTree.CreateTree(root, arr)
	//BinarySearchTree.PrintTree(root)
	//
	//root.InorderTraversal()

	var root *AVL_Tree.Node
	root = AVL_Tree.CreateTree(root, arr)
	//AVL_Tree.PrintTree(root)
	//fmt.Println(root)
	AVL_Tree.InorderTraversal(root)
	AVL_Tree.PrintTree2(root)
	fmt.Println("--------------------------------------------------------")
	root = root.Remove(74)
	//AVL_Tree.Balance(root)
	AVL_Tree.PrintTree2(root)
	//test := root.Remove(13)
	//root = AVL_Tree.Balance(root)
	fmt.Println("val = ", root.Value, "balance Factor = ", root.BalanceFactor())
	//fmt.Println(root.Visualize(0))
}
