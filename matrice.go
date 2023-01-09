package main

import "fmt"

func main() {

	mat := createMat(3, 1)
	matProduct(mat, mat)
}

func printMat(mat [][]int) {
	//This function displays any matrix on the terminal
	fmt.Println("***Matrice***")
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			fmt.Print((mat)[i][j], " ")
		}
		fmt.Println("")
	}
}
func createMat(n int, k int) [][]int {
	//This function creates a square slice matrix of size n,
	//containing j in every cell
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, n)
	}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			mat[i][j] = k
		}
	}
	fmt.Println("La matrice suivante a été crée : ")
	printMat(mat)
	return mat
}

func createEmptyMat(n int) [][]int {
	//This function creates an empty square slice matrix of size n
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, n)
	}
	fmt.Println("Une matrice vide a été créée ")
	return mat
}

func matProduct(mat1 [][]int, mat2 [][]int) [][]int {
	//this function calculates, returns and displays the product of mat1 and mat2
	fmt.Println("***Produit de deux matrices***")
	res := createEmptyMat(len(mat1))
	for i := 0; i < len(mat1); i++ {
		for j := 0; j < len(mat1); j++ {
			for k := 0; k < len(mat1); k++ {
				res[i][j] += mat1[i][k] * mat2[k][j]
			}
		}
	}
	printMat(res)
	return res
}
