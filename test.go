package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

/*func check(e error) {
	if e != nil {
		panic(e)
	}
}*/

func ReadFile(fileName string) {

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
}

/*func WriteFile(text string, file *os.File) {
	if _, err := file.WriteString(text); err != nil {
		panic(err)
	}
}*/

func main() {
	fmt.Println("J'aime Go")
	fmt.Println("Il est ", time.Now())
	ReadFile("test.txt")

	/*file,err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer file.Close()
	check(err)
	WriteFile("Brocolis\n", file)*/

}
