package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
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
	fmt.Println("J'aime Go")
	fmt.Println("Il est ", time.Now())
	//ReadFile("test.txt")
	/*fmt.Println("J'aime Go")
	fmt.Println("Il est ", time.Now())*/
	//data := ReadFile("matrice_de_test.txt \n")
	data := "1 1 1\n1 1 1\n1 1 1"
	print("\n", data)
	data2 := strings.Split(data, " ")
	fmt.Printf("\n Print de la matrice sous forme de tableau de char \n")
	for i := 0; i < len(data2); i++ {
		fmt.Printf(data2[i])
	}

	fmt.Printf("\n Détection de la taille de la matrice \n")
	n := 0
	for i := 0; i < len(data2); i++ {
		if data2[i] == "\n" {
			n = i
			break
		}
	}

	data3 := make([][]string, n)
	for i := range data3 {
		data3[i] = make([]string, n)
	}

	fmt.Printf("On met tous les chiffres dans un tableau\n")
	tab := make([]string, n*n)
	for i := range data2 {
		splitage := strings.Split(data2[i], "\n")
		for j := range splitage {
			tab = append(tab, splitage[j])
		}
	}

	for i := range tab {
		fmt.Printf(tab[i] + " ")
	}

	fmt.Printf("\nOn convertit les strings du tableau en entiers\n")
	tabint := make([]int, n)
	fmt.Print("len(tab):", len(tab))
	conversion, err := strconv.Atoi(tab[0])
	tabint[0] = conversion
	fmt.Print("\n tabint[0] : ", tabint[0])
	if err != nil {
		fmt.Println("Error during conversion")
	}
	/*for i := range tab {
		fmt.Print("On entre dans la première boucle")
		fmt.Print("i=", i)
		conversion, err := strconv.Atoi(tab[i])
		tabint[i] = conversion
		fmt.Print(tabint[i])
		if err != nil {
			fmt.Println("Error during conversion")
		}
	}/*
	for i := range tabint {
		fmt.Print(tabint[i])
	}

	/*for i := 0; i < len(data2); i++ {
		for j :=0; j<n; j++{
			data3[]
		}
	}*/
}
