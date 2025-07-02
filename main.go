package main

import (
	"fmt"
)

func main() {
    // Variables

    var name string = "Omar"
    fmt.Printf("This is my name %s\n", name)

    age := 43
    fmt.Printf("This is my agr %d\n", age)

    var city string = "Rawalpindi"
    fmt.Printf("This is my city %s\n", city)

    country := "Pakistan"
    fmt.Printf("This is my country %s\n", country)

    var continent, planet string = "Asia", "Earth"
    fmt.Printf("This is my continent %s and this is my planet %s\n", continent, planet)

    var (
        isEmployed bool = true
        salary int = 50000
        position string = "developer"
    )

    fmt.Printf("isEmployed : %t this is my salary : %d and this is my position: %s\n", isEmployed, salary, position)

    // Zero Values
    var defaultInt int
    var defaultFloat float64
    var defaultString string
    var defaultBool bool

    fmt.Printf("int: %d float: %f string %s and bool %t\n", defaultInt, defaultFloat, defaultString, defaultBool)

    // Constants and Enums

    const pi = 3.14

    const (    
        Monday = 1
        Tuesday = 2
        Wednesday = 3
    )      

    fmt.Printf("Monday: %d - Tuesday %d Wednesday %d\n", Monday, Tuesday, Wednesday)

    const typedAge int = 25
    const untypedAge = 25

		fmt.Println(typedAge == untypedAge)

		const (
			Jan = iota + 1 //1
			Feb
			Mar
			Apr
		)

		fmt.Printf("jan - %d feb - %d mar - %d apr - %d\n", Jan, Feb, Mar, Apr )

		fmt.Println("This is the result",add(1,2))

		result:= add(3,4)
		fmt.Println("This is the result", result)

		sum, product := calculateSumAndProduct(10, 10)
		fmt.Printf("this is sum: %d and this is product: %d\n", sum, product)

		// sum, _ := calculateSumAndProduct(10, 10)
		// fmt.Printf("this is sum: %d\n", sum)
}

func add (a int, b int) int {
	return a + b
}

func calculateSumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}