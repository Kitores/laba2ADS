package main

import (
	"TreesLaba2/internal/measurements"
	"TreesLaba2/internal/plotting"
	"fmt"
)

func main() {
	//arr := []int{8, 3, 10, 1, 6, 14, 4, 7, 13, 23, 5, 2, 52}
	////arr := pkg.RandArray(100, 12)
	////fmt.Println(arr)
	//Red_Black_Tree.Test(arr)

	//trees := []string{"BST", "AVL", "RBT"}
	trees := []string{"RBT"}

	for _, tree := range trees {
		arrX, arrY := measurements.Calculate(tree, 12)
		fmt.Println(arrX, arrY)
		plotting.CreateLineChartByPlotter(arrX, arrY, tree)
		plotting.CreateLineChart(arrX, arrY, tree)
	}

}
