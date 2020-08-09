package fileio

// FIle I/O example

import (
	"fmt"
	"os"
	"reflect"
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

	file2, err := os.Create("/home/carlosngv/newFile.go")
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
	fmt.Println("TYPE:", reflect.TypeOf(fileInfo.Size()))
	defer file.Close()
	data := make([]byte, fileInfo.Size())
	file.Read(data) // Reads the data from file to byte array
	fmt.Println("The data is", string(data))

}

// ChangePermission exported
func ChangePermission() {
	err := os.Chmod("stringFile", 0777) // Gives user, groups and others file permissions.
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Permissions changed!")
	}
}
