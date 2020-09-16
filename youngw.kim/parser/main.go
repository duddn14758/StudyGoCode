package main

import (
	"bufio"
	"fmt"
	//"io/ioutil"
	"os"
	"strings"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	request := [2]int{0}
	fmt.Println("input your log file Name")
	//logFileName, _ := reader.ReadString('\n')
	//logFileName = "C:/Go/src/github.com/youngw.kim/parser/" + logFileName
	logFileName := "testLog.log"
	fmt.Println("input file name is", logFileName)

	/*
		dat, err := ioutil.ReadFile(logFileName)
		if err != nil {
			fmt.Print("ioutil has been an error!: ", err)
		}
		fmt.Println(dat)
	*/
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
