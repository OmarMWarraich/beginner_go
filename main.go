package main

import (
	"fmt"
)

/* func main() {
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
} */

// func main() {

// 		/* loops */

//     age:= 30

//     if age >= 18 {
//         fmt.Println("You are an adult")
//     } else if age >= 13 {
//         fmt.Println("You are a teenager")
//     } else {
//         fmt.Println("You are a child")
//     }

//     day := "Tuesday"

//     switch day {
//     case "Monday":
// 			fmt.Println("start of the week")
// 		case "Tuesday", "Wednesday", "Thursday":
// 			fmt.Println("midweek")
// 		case "Friday":
// 			fmt.Println("TGIF (Thank God It's Friday)")
// 		default:
// 			fmt.Println("it's the weekend")
//     }

// 		for i:= 0; i < 5; i++ {
// 			fmt.Println("This is i", i)
// 		}

// 		/* There are no while loops in go*/

// 		counter := 0
// 		for counter < 3 {
// 			fmt.Println("This is the counter", counter)
// 			counter++
// 		}

// 		/* arrays & slices */

// 		numbers := [5]int{10,20,30,40,50}
// 		fmt.Printf("this is length of the array %v\n", len(numbers))
// 		fmt.Printf("this is out array %v\n", numbers)
// 		fmt.Printf("this is index 1 %v\n", numbers[1])
// 		numbers[1] = 90
// 		fmt.Printf("this is index 1 %v\n", numbers[1])
// 		fmt.Printf("this is last value %v\n", numbers[len(numbers)-1])

// 		/* matrix 2 x 3 */

// 		matrix := [2][3]int{
// 			{1,2,3},
// 			{4,5,6},
// 		}
// 		fmt.Printf("this is our matrix %v\n", matrix)

// 		/* slices are dynamic arrays */
// 		/* allNumbers := numbers[:]
// 		firstThreee := numbers[0:3] */

// 		fruits := []string{"apple", "banana", "strawberry"}
// 		fmt.Printf("These are my fruits: %v", fruits)
// 		fmt.Println()

// 		fruits = append(fruits, "kiwi")
// 		fmt.Printf("These are my fruits: %v", fruits)
// 		fmt.Println()

// 		moreFruits := []string{"blueberries", "pineapple"}
// 		fruits = append(fruits, moreFruits...)
// 		fmt.Printf("These are my fruits: %v", fruits)
// 		fmt.Println()

// 		for index, value := range numbers {
// 			fmt.Printf("index: %d and value: %d", index, value)
// 			fmt.Println()
// 		}

// 		capitalCities := map[string]string {
// 			"Pakistan": "Islamabad",
// 			"China": "Beijing",
// 			"USA": "Washington DC" ,
// 		}

// 		fmt.Println(capitalCities["USA"])

// 		capital, exists := capitalCities["Germany"]

// 		if exists {
// 			fmt.Println("this is the capital", capital)
// 		} else {
// 			fmt.Println("does not exist")
// 		}

// 		delete(capitalCities, "USA")
// 		fmt.Printf("this is the updated map: %v\n", capitalCities)
// }

// /* functions */

// func add (a int, b int) int {
// 	return a + b
// }

// func calculateSumAndProduct(a, b int) (int, int) {
// 	return a + b, a * b
// }

// defined struct: not anonymous struct
type Person struct {
	Name string
	Age int
}

func main() {
	person := Person{Name: "John", Age: 25}
	fmt.Printf("This is out person %+v\n", person)

	// anonymous struct
	employee := struct {
		name string
		id int
	}{
		name: "alice",
		id: 123,
	}
	fmt.Printf("This is out employee %+v\n", employee)

	// nested struct

	type Address struct {
		Street string
		City string
	}

	type Contact struct {
		Name string
		Address Address
		Phone string
	}

	contact := Contact{
		Name: "Marc",
		Address: Address{
			Street: "123 Main street",
			City: "Anytown",
		},
	}

	fmt.Printf("This is out contact %+v\n", contact)
	fmt.Println("name before: ", person.Name)
	modifyPersonName(person)
	fmt.Println("name after: ", person.Name)

	fmt.Println("name before: ", person.Name)
	modifyPointerPersonName(&person)
	fmt.Println("name after: ", person.Name)

	x := 20
	ptr := &x
	fmt.Printf("value of x: %d and address of x %p\n", x, ptr)
	*ptr = 30
	fmt.Printf("value of new x: %d and address of x %p\n", x, ptr)

}

func modifyPersonName(person Person){
	person.Name = "Melkey"
	fmt.Println("inside scope: new name", person.Name)
}

func modifyPointerPersonName(person *Person){
	person.Name = "Melkey"
	fmt.Println("inside scope: new name", person.Name)
}