package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadFile(fileName string, c chan string) {

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
	c <- s
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

func createEmptyMat(n int) [][]int {
	//This function creates an empty square slice matrix of size n
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, n)
	}
	//fmt.Println("Une matrice vide a été créée ")
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
	//fmt.Println("La matrice suivante a été crée : ")
	//printMat(mat)
	return mat
}
func matProduct1(mat1 [][]int, mat2 [][]int, c chan []int, i int) {
	//this function calculates, returns and displays the product of mat1 and mat2
	//fmt.Println("***Produit de deux matrices***")
	res := make([]int, len(mat1))

	for j := 0; j < len(mat1); j++ {
		for k := 0; k < len(mat1); k++ {
			res[j] += mat1[i][k] * mat2[k][j]
		}
		//fmt.Print(res[j], " ")
	}
	//fmt.Print("")
	c <- res

}

func matriceEnInt(data string) [][]int {
	data2 := strings.Split(strings.ReplaceAll(data, "\r\n", "\n"), "\n")
	fmt.Println("string split a la ligne")
	fmt.Println(data2)
	fmt.Printf("\n Détection de la taille de la matrice \n")
	n := len(data2)
	fmt.Println("La matrice est de taille", n)
	fmt.Print(data2[0])
	fmt.Print(data2[1])
	fmt.Print(data2[2])

	//On fait une matrice de strings
	data3 := make([][]string, n)
	for i := range data3 {
		data3[i] = make([]string, n)
	}

	for i := range data2 {
		splitage := strings.Split(data2[i], " ")
		for j := range splitage {
			data3[i][j] = splitage[j]
		}
	}

	fmt.Printf("\n Print de la matrice sous forme de matrice de char \n")
	for i := 0; i < len(data3); i++ {
		for j := 0; j < len(data3); j++ {
			fmt.Printf(data3[i][j])
		}
		fmt.Println("")
	}

	fmt.Println("On convertit les strings du tableau en entiers")

	matrice_finale := make([][]int, n)
	for i := range matrice_finale {
		matrice_finale[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			conversion, err := strconv.Atoi(data3[i][j])
			fmt.Println("Conversion = ", conversion)
			matrice_finale[i][j] = conversion
			if err != nil {
				fmt.Println("Error during conversion")
			}
		}
	}
	fmt.Printf("\n Print de la matrice finale (en int) \n")
	for i := 0; i < len(matrice_finale); i++ {
		for j := 0; j < len(matrice_finale); j++ {
			fmt.Print(matrice_finale[i][j], " ")
		}
		fmt.Println("")
	}

	return matrice_finale

}

func main() {
	var start time.Time
	start = time.Now()

	/*cA := make(chan string)
	cB := make(chan string)
	go ReadFile("matriceA.txt", cA)
	go ReadFile("matriceB.txt", cB)
	matA := <-cA
	matB := <-cB*/

	matAint := createMat(100, 5)
	matBint := createMat(100, 1)

	//matAint := matriceEnInt(matA)
	//matBint := matriceEnInt(matB)

	channel_second := make([]chan []int, len(matAint))
	for i := range channel_second {
		channel_second[i] = make(chan []int)
	}

	matC := make([][]int, len(matAint))
	for i := 0; i < len(matAint); i++ {
		go matProduct1(matAint, matBint, channel_second[i], i)
		matC[i] = <-channel_second[i]
	}
	//printMat(matC)

	go ecritDansFichier(matC, "matriceC.txt")
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Print("time", elapsed)

}
