package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type monitorLog struct {
	time              string
	tp                float64
	desired_octet_scg int
	desired_octet_mcg int
	ddds_scg          int
	ddds_mcg          int
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
	log := [100]monitorLog{}

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
		} else if strings.Contains(scanner.Text(), "DDDS Count") {
			res_ddds := strings.Split(standardizeSpaces(scanner.Text()), " ")

			log[index].desired_octet_scg, _ = strconv.Atoi(res_ddds[4])
			log[index].desired_octet_mcg, _ = strconv.Atoi(res_ddds[9])

		} else if strings.Contains(scanner.Text(), "Desired Octet") {
			res_ddds := strings.Split(standardizeSpaces(scanner.Text()), " ")

			log[index].ddds_scg, _ = strconv.Atoi(res_ddds[4])
			log[index].ddds_mcg, _ = strconv.Atoi(res_ddds[9])

		} else if strings.Contains(scanner.Text(), "Dowlink RX") {

			res_tp := strings.Split(standardizeSpaces(scanner.Text()), " ")

			log[index].tp, _ = strconv.ParseFloat(res_tp[10], 64)

			log[index].tp, _ = strconv.ParseFloat(res_tp[10], 64)

			if strings.Compare(res_tp[11], "Mbps") == 0 {
				log[index].tp *= 0.001
			} else if strings.Compare(res_tp[11], "Kbps") == 0 {
				log[index].tp *= 0.000001
			} else if strings.Compare(res_tp[11], "bps") == 0 {
				log[index].tp = 0
			}

			avg += log[index].tp

			if log[index].tp < 10 {
				index++
			}

		}

	}
	for i := 0; i < index; i++ {
		fmt.Printf("time : %s\n desired_scg : %d desired_mcg : %d\n ddds_scg : %d ddds_mcg : %d\n  tp : %0.2fMbps\n-----------------\n",
			log[i].time, log[i].desired_octet_scg, log[i].desired_octet_mcg, log[i].ddds_scg, log[i].ddds_mcg, log[i].tp)
	}
	fmt.Printf("avg : %05.2fMbps\n", float64(avg)/float64(index))
}
