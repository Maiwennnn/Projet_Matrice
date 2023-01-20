package main

import (
	"fmt"
	"strings"
)

/*
func ReadFile(fileName string) string {

	fmt.Printf("\n\nReading a file in Go lang\n")
	//fileName := "test.txt"

	// The ioutil package contains inbuilt
	// methods like ReadFile that reads the
	// filename and returns the contents.
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", fileName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data) // conversion de byte en string
	s := string(data)
	return s
}

func ecritDansFichier(mat [][]int, nameFile string) {
	file, err := os.OpenFile(nameFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
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
}*/

func main() {
	mata := ReadFile("matriceA.txt")
	matb := ReadFile("matriceB.txt")

	mata2 := strings.Split(mata, " ")
	fmt.Printf("Print de la matrice sous forme de tableau de char \n")
	for i := 0; i < len(mata2); i++ {
		fmt.Printf(mata2[i])
	}
	mata3 := make([]string, 1)
	for i := 0; i < len(mata2); i++ {
		mata3 := append(strings.Split(mata2[i], "\n"))
	}
	print(len(mata3))
	for i := 0; i < len(mata3); i++ {
		fmt.Printf(mata3[i])
	}

	matb2 := strings.Split(matb, " ")
	fmt.Printf("Print de la matrice sous forme de tableau de char \n")
	for i := 0; i < len(matb2); i++ {
		fmt.Printf(matb2[i])
	}
	matb3 := make([]string, 1)
	for i := 0; i < len(matb2); i++ {
		matb3 := append(strings.Split(matb2[i], "\n"))
	}
	print(len(matb3))
	for i := 0; i < len(matb3); i++ {
		fmt.Printf(matb3[i])
	}

	//matC := matProduct()

}
