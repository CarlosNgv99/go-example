package main

import (
	"fmt"
	"go-example/fileio"
	_ "os"
	_ "time"
)

func main() {
	/*fmt.Println("\n---- Interface Example -------")
	ice.GenCar()

	fmt.Println("\n---- File Creation Example -------")
	filemng.CreateFile()*/

	fmt.Println("\n---- File I/O Example -------")
	fileio.CreateFile()
	fileio.ReadFile()
	//fileio.ChangePermission()

	/*grades := []int{122, 23, 432}
	value := grades[1:] // first 2 elements
	slice1 := make([]int, 3, 100)

	fmt.Println("\n---- Slice1 -------")
	fmt.Println("Value:", slice1)
	fmt.Println("Capacity:", cap(slice1))
	fmt.Println("Length:", len(slice1))
	fmt.Println("\n")

	fmt.Println("---- Value -------")
	fmt.Println("Value:", value)
	fmt.Println("\n")

	fmt.Println("---- Grades -------")
	fmt.Println("Value:", grades)
	fmt.Println("Length:", len(grades))
	fmt.Println("\n")

	fmt.Println("-------- MAPS --------")

	mapExample := make(map[string]int)
	mapExample = map[string]int{
		"Bye":  12,
		"Good": 32,
	}
	mapExample["Hello"] = 34
	delete(mapExample, "Hello")
	fmt.Println(mapExample)
	fmt.Println("\n")

	fmt.Println("-------- STRUCT --------")
	man := Man{"Carlos", 12, "Male", Job{"Apple", "Developer", 45000}}
	fmt.Println(man)*/

}

/*

type Job struct {
	company  string
	position string
	salary   int
}

type Man struct {
	name   string
	age    int
	gender string
	Job
}
*/
