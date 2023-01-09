package main

import "fmt"

func main() {

	var mat = *remplitMatrice(3, 1)
	for i := 0; i < len(mat); i++ {

	}

}

func remplitMatrice(n int, p int) *[3][3]int {
	var mat *[3][3]int
	for i := 0; i < len(*mat); i++ {
		for j := 0; j < len(*mat); j++ {
			(*mat)[i][j] = p
			fmt.Print((*mat)[i][j], " ")
		}
		fmt.Println("")
	}

	return mat
}
