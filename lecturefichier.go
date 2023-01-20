package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*func ReadFile(fileName string) string {

	fmt.Printf("\n\nReading a file in Go lang\n")
	data = iotuil.Readfile(fileName)
	fmt.Printf("\nFile Name: %s", fileName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)
	s := string(data)
	return s
}*/

func main() {
	//fmt.Println("J'aime Go")
	//fmt.Println("Il est ", time.Now())
	//ReadFile("test.txt")
	/*fmt.Println("J'aime Go")
	fmt.Println("Il est ", time.Now())*/
	//data := ReadFile("matrice_de_test.txt \n")
	data := "1 1 1\n1 1 1\n1 1 1"
	print("\n", data)
	data2 := strings.Split(data, "\n")
	fmt.Printf("\n Détection de la taille de la matrice \n")
	n := len(data2)
	fmt.Println("La matrice est de taille", n)

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
}
