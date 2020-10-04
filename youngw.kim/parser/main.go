package main

import (
	"bufio"
	"fmt"
	"github.com/StudyGoCode/youngw.kim/parser/types"
	"os"
	"strings"
)

func main() {
	request := []int{0}
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

	//Blogs := []log_type_B{}

	for scanner.Scan() { // read file per line
		if strings.Contains(scanner.Text(), types.LOG_TYPES[0]) {
			LogF := types.First_case_log{}
			request[0]++
			//nowIndex := strings.IndexAny(scanner.Text(), "[")
			//fmt.Println(nowIndex)
			fmt.Println("Original message : ", scanner.Text())
			res := strings.SplitAfter(scanner.Text(), "]")
			for _, value := range res {
				fmt.Println(value)

			}
			LogF.SetVal(scanner.Text())
			//fmt.Println((strings.Split(scanner.Text(), ":")))

		} else if strings.Contains(scanner.Text(), types.LOG_TYPES[1]) {
			request[1]++
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error occured")
		}
		break
	}
	fmt.Println(request)
}
