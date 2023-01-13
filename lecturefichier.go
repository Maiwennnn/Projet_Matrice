package main

import (
	"fmt"
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
	fmt.Printf("Print de la matrice sous forme de tableau de char \n")
	for i := 0; i < len(data2); i++ {
		fmt.Printf(data2[i])
	}
	data3 := make([]string, 1)
	for i := 0; i < len(data2); i++ {
		data3 := append(strings.Split(data2[i], "\n"))
	}
	print(len(data3))
	for i := 0; i < len(data3); i++ {
		fmt.Printf(data3[i])
	}
}
