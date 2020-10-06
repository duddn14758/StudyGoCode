package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type log_tp struct {
	time string
	tp   float64
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func splitTimeCase(r rune) bool {
	return r == ' ' || r == '.'
}

func main() {
	var logFileName string
	index := 0
	avg := 0.0

	log := [100]log_tp{}

	//tp := [100]float64{}
	//fmt.Scan(&logFileName)
	logFileName = "test.log"
	f, err := os.Open(logFileName)
	if err != nil {
		fmt.Print("There has been an error!: ", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "2020-") {
			// parsing time val
			res_time := strings.FieldsFunc(scanner.Text(), splitTimeCase)
			log[index].time = res_time[1]

			// parsing throughput val
			if scanner.Scan() != true {
				fmt.Println("No more lines")
				break
			}
			res_tp := strings.Split(standardizeSpaces(scanner.Text()), " ")

			log[index].tp, _ = strconv.ParseFloat(res_tp[10], 64)

			if strings.Compare(res_tp[11], "Mbps") == 0 {
				log[index].tp *= 0.001
			} else if strings.Compare(res_tp[11], "Kbps") == 0 {
				log[index].tp *= 0.000001
			}
			avg += log[index].tp

			if log[index].tp < 10 {
				index++
			}
		}
	}
	for i := 0; i < index; i++ {
		fmt.Printf("%s %0.3fGbps\n", log[i].time, log[i].tp)
	}
	fmt.Printf("avg : %0.3fGbps\n", float64(avg)/float64(index))
}
