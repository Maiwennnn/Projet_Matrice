package main

import "fmt"

func main() {

	mat := make([][]int, 3)
	for i := range mat {
		mat[i] = make([]int, 3)
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			mat[i][j] = 1
		}
	}
	printMat(mat)

	res := make([][]int, 3)
	for i := range res {
		res[i] = make([]int, 3)
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			for k := 0; k < len(mat); k++ {
				res[i][j] += mat[i][k] * mat[k][j]
			}
		}
	}
	printMat(res)
}

func printMat(mat [][]int) {
	fmt.Println("***Matrice***")
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			fmt.Print((mat)[i][j], " ")
		}
		fmt.Println("")
	}
}
