package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

func main() {
	str := ""
	str1 := "AAAAAAAAAA"
	// 1.
	start := time.Now()
	for i := 0; i < 100000; i++ {
		str += str1
	}
	elapsed := time.Since(start)
	fmt.Printf("strlen(%d) : %v\n", len(str), elapsed)
	// 2.
	var b bytes.Buffer
	str = ""
	start = time.Now()
	for i := 0; i < 100000; i++ {
		b.WriteString(str1)
	}
	str = b.String()
	elapsed = time.Since(start)
	fmt.Printf("strlen(%d) : %v\n", len(str), elapsed)
	// 3.
	str = ""
	start = time.Now()
	for i := 0; i < 100000; i++ {
		str = fmt.Sprintf("%s%s", str, str1)
	}
	elapsed = time.Since(start)
	fmt.Printf("strlen(%d) : %v\n", len(str), elapsed)
	// 4.
	str = ""
	mySlice := []string{}
	for i := 0; i < 100000; i++ {
		mySlice = append(mySlice, str1)
	}
	start = time.Now()
	str = strings.Join(mySlice, "")
	elapsed = time.Since(start)
	fmt.Printf("strlen(%d) : %v\n", len(str), elapsed)
}
