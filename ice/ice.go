package ice

// Struct and interface example

import "fmt"

type vehicle interface {
	Start() string
	Break() string
}

type car struct {
	brand    string
	velocity string
}

func (myCar car) Start() string {
	return myCar.brand + " has been started!"
}

// GenCar exported
func GenCar() {

	newCar := car{"Toyota", "100 KM/H"}

	fmt.Println(newCar.Start())
}
