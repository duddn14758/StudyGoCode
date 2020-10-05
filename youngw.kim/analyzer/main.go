package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func main() {
	var logFileName string
	index := 0
	avg := 0.0

	tp := [100]float64{}
	//fmt.Scan(&logFileName)
	logFileName = "test.log"
	f, err := os.Open(logFileName)
	if err != nil {
		fmt.Print("There has been an error!: ", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "bps") {
			res := strings.Split(standardizeSpaces(scanner.Text()), " ")
			fmt.Println("---------------")
			for i, _ := range res {
				fmt.Printf("res[%d] : %s\n", i, res[i])
			}
			//fmt.Println(strconv.ParseFloat(res[1], 64))
			tp[index], _ = strconv.ParseFloat(res[1], 64)
			//tp = append(strconv.ParseFloat(res[1], 64), index)

			if strings.Compare(res[2], "Mbps") == 0 {
				tp[index] *= 0.001
			}
			avg += tp[index]

			index++
		}
	}
	for i := 0; i < index; i++ {
		fmt.Printf("%0.3fGbps\n", tp[i])
	}
	fmt.Printf("avg : %0.3fGbps\n", float64(avg)/float64(index))
}
