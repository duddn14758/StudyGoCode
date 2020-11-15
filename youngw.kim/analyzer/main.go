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
	scg_tp            float64
	mcg_tp            float64
	total_tp          float64
	fresh_buffer      int
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

func bitAliasMbps(unit string, tp float64) float64 {
	if strings.Compare(unit, "Gbps") == 0 {
		tp *= 1000
	} else if strings.Compare(unit, "Kbps") == 0 {
		tp *= 0.001
	} else if strings.Compare(unit, "bps") == 0 {
		tp = 0
	}
	return tp
}

func bitAliasGbps(unit string, tp float64) float64 {
	if strings.Compare(unit, "Mbps") == 0 {
		tp *= 0.001
	} else if strings.Compare(unit, "Kbps") == 0 {
		tp *= 0.000001
	} else if strings.Compare(unit, "bps") == 0 {
		tp = 0
	}
	return tp
}

func main() {
	var logFileName string
	index := 0
	avg_total := 0.0
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

			log[index].total_tp, _ = strconv.ParseFloat(res_tp[10], 64)

			log[index].total_tp = bitAliasMbps(res_tp[11], log[index].total_tp)

			avg_total += log[index].total_tp

		} else if strings.Contains(scanner.Text(), "SCG TX") {
			res_tp := strings.Split(standardizeSpaces(scanner.Text()), " ")

			log[index].scg_tp, _ = strconv.ParseFloat(res_tp[4], 64)
			log[index].mcg_tp, _ = strconv.ParseFloat(res_tp[10], 64)

			log[index].scg_tp = bitAliasMbps(res_tp[5], log[index].scg_tp)
			log[index].mcg_tp = bitAliasMbps(res_tp[11], log[index].mcg_tp)

			index++
		}

	}
	fmt.Println("    time | desired_scg | desired_mcg | ddds_scg | ddds_mcg | scg_tp | mcg_tp")
	for i := 0; i < index; i++ {
		fmt.Printf("%s  %12d  %12d   %8d   %8d   %6.2f   %6.2f\n",
			log[i].time, log[i].desired_octet_scg, log[i].desired_octet_mcg, log[i].ddds_scg, log[i].ddds_mcg, log[i].scg_tp, log[i].mcg_tp)
	}
	fmt.Printf("avg_total : %05.2fMbps\n", float64(avg_total)/float64(index))
}
