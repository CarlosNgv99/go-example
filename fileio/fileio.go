package fileio

// FIle I/O example

import (
	"fmt"
	"os"
)

// FIle I/O using os package

// CreateFile exported
func CreateFile() {
	file, err := os.Create("myfile")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := []byte("Hello World")
	file.Write(data)

	file2, err := os.Create("stringFile")
	if err != nil {
		panic(err)
	}
	defer file2.Close()
	file2.WriteString("File String!")
}

// ReadFile exported
func ReadFile() {
	file, err := os.Open("myfile")
	if err != nil {
		panic(err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println("SIZE:", fileInfo.Size())
	defer file.Close()
	data := make([]byte, fileInfo.Size())
	file.Read(data) // Reads the data from file to byte array
	fmt.Println("The data is", string(data))

}
