package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	request := [2]int{0}
	fmt.Println("input your log file Name")
	logFileName := "testLog.log"
	fmt.Println("input file name is", logFileName)

	f, err := os.Open(logFileName)
	if err != nil {
		fmt.Print("There has been an error!: ", err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	fmt.Println("File read Success")

	//Alogs := []log_type_A{}
	//Blogs := []log_type_B{}

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "FIRST_CASE_PARSING_MSG") {
			request[0]++
			nowIndex := strings.IndexAny(scanner.Text(), ":")
			fmt.Println(nowIndex)
			fmt.Println(scanner.Text())
			//fmt.Println((strings.Split(scanner.Text(), ":")))

		} else if strings.Contains(scanner.Text(), "SECOND_CASE_PARSING_MSG") {
			request[1]++
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error occured")
		}
	}
	fmt.Println(request)
}
