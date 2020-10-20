package main

import "fmt"

var dx = [4]int{0, -1, 0, 1}
var dy = [4]int{-1, 0, 1, 0}

type queue struct {
	x, y int
}

func main() {
	for i := 0; i < 4; i++ {
		fmt.Println("(", dx[i], ", ", dy[i], ")")
	}
}
