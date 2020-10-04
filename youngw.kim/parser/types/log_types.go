package types

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	TIME = iota
	FIRST
	SECOND
	THIRD
)

var LOG_TYPES = []string{"FIRST_CASE_PARSING_MSG", "SECOND_CASE_PARSING_MSG"}

type Log_type_AA struct {
	first, second, third int
}

type Log_type_BB struct {
	fourth, fifth, sixth int
}

type Log_type_CC struct {
	seventh, eighth int
}

type First_case_log struct {
	LogA Log_type_AA
	LogB Log_type_BB
	LogC Log_type_CC
}

type log_type_A struct {
	name               string
	AB, BC, CD, DE, EF int
}

type log_type_B struct {
	name               string
	aa, bb, cc, dd, ee int
}

func Split(r rune) bool {
	return r == ':' || r == ' '
}

func (l *Log_type_AA) setLogTypeA(logA string) {
	strings.TrimPrefix(logA, LOG_TYPES[0])
	res := strings.FieldsFunc(logA, Split)
	fmt.Println(res)
	for i, val := range res {
		fmt.Printf("    val2[%d]: %s\n", i, val)
		switch i {
		case 2:
			l.first, _ = strconv.Atoi(val)
			break
		case 4:
			l.second, _ = strconv.Atoi(val)
			break
		case 6:
			l.third, _ = strconv.Atoi(strings.TrimSuffix(val, "]"))
			break
		}
	}
	fmt.Println(l)
}

func (l *First_case_log) SetVal(LogF string) {
	res := strings.SplitAfter(LogF, "]")
	for i, value := range res {
		fmt.Printf("val1[%d]: %s\n", i, value)
		if i == 0 {
			continue
		}
		res1 := strings.FieldsFunc(value, Split)
		for j, val2 := range res1 {
			fmt.Printf("  val2[%d]: %s\n", j, val2)
		}
		switch i {
		case 0:
			l.LogA.setLogTypeA(value)
			break
		case 1:
			break
		}
		fmt.Println("  val2:", l.LogA)
	}
}

func (a Log_type_AA) PrintVal(logA string) {
	fmt.Println(logA)
}
