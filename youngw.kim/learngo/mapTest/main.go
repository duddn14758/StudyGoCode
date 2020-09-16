package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	duddn := map[string]string{"name": "youngwoo", "age": "27", "food": "pizza"}
	favFood := []string{"Pizza", "Beef", "Hamburger"}
	young := person{name: "youngwoo", age: 20, favFood: favFood}
	fmt.Println(young.name, young.age, young.favFood)
	fmt.Println(young)

	for key, value := range duddn {
		fmt.Println(key, value)
	}
}
