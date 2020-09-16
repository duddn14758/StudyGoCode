package main

import (
	"fmt"
	"strings"

	"github.com/youngw.kim/learngo/stringTest/some"
)

func multiply(a, b int) int {
	return a * b
}

func add(a, b int) int {
	return a + b
}

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("lenAndUpper function return value") // defer : running lines after this function returns
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
	//return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func enumerating(numbers ...int) int {
	defer fmt.Println("numbers return")
	total := 0
	for index, number := range numbers {
		fmt.Println(index, number)
		total += number
	}
	return total
}

func canIdrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 19 {
		return false
	}
	return true
}

func canIdrinkSwitch(age int) bool {
	/*
		switch koreanAge := age + 2; {
		case koreanAge < 20:
			return false
		case koreanAge == 20:
			return true
		case koreanAge > 50:
			return false
		}
		return true
	*/
	/*
		switch {
		case age < 18:
			return false
		case age == 18:
			return true
		case age > 50:
			return false
		}
		return false
	*/
	switch age {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func main() {
	//const name = "young"
	var name string = "youngw" // name := "youngw" --> shortcut only works at in function
	name = "0woo"
	fmt.Println(name)
	some.SayHello()
	fmt.Println(multiply(2, 2))
	fmt.Println(add(3, 4))
	fLen, fName := lenAndUpper("young")
	//fLen, _ := lenAndUpper("young")
	fmt.Println(fName, fLen)
	repeatMe("turtle", "lion", "dolphin", "tiger", "chicken")
	fmt.Println(enumerating(1, 2, 3, 4, 5))
	fmt.Println(canIdrink(16))
	fmt.Println(canIdrinkSwitch(16))
}
