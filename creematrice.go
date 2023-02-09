package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func createMat2(n int) [][]int {
	//This function creates a square slice matrix of size n,
	//containing j in every cell
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, n)
	}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			mat[i][j] = rand.Intn(10)
		}
	}
	return mat
}

func ecritDansFichier(mat [][]int, nameFile string) {
	file, err := os.Create(nameFile)
	defer file.Close() // on ferme automatiquement à la fin de notre programme
	for i := 0; i < len(mat); i++ {
		s := ""
		for j := 0; j < len(mat); j++ {
			s = s + strconv.Itoa(mat[i][j]) + " "
		}

		_, err = file.WriteString(s) // écrire dans le fichier
		s = "\n"
		_, err = file.WriteString(s)
	}
	fmt.Print(err)

}

func main() {
	ecritDansFichier(createMat2(100), "matriceB.txt")
}
