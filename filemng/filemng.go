package filemng

import (
	"fmt"
	"log"
	"os"
)

func CreateFile() {

	fmt.Println("Creating file...")
	file, err := os.Create("example.text")

	if err != nil {
		log.Fatalf("An error has occured. Please try again.")
	}

	defer file.Close()

	len, err := file.WriteString("This is a demonstration of a file creation!")

	if err != nil {
		log.Fatalf("An error has occured. Please try again.")
	}

	fmt.Println("File name:", file.Name())
	fmt.Println("Length: ", len)
}

func readFile() {

}

func init() {
	CreateFile()
	readFile()
}
