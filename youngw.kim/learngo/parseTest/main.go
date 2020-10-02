package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello Hi Anyeong"
	fmt.Println(s)

	res1 := strings.Split(s, " ")
	fmt.Println(res1)

	res2 := strings.SplitAfter(s, " ")
	fmt.Println(res2)
}
