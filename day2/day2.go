package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readReportSlice(filename string) ([][]int32, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scn := bufio.NewScanner(file)

	var reports [][]int32
	for scn.Scan() {
		line := scn.Text()

		if len(line) == 0 {
			break
		}

		numbers, err := strSliceToIntSlice(line)

		if err != nil {
			return nil, err
		}

		reports = append(reports, numbers)
	}
	return reports, nil
}

func strSliceToIntSlice(numbers string) ([]int32, error) {
	fields := strings.Fields(numbers)

	var numList []int32
	for _, numStr := range fields {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, err
		}
		numList = append(numList, int32(num))
	}

	return numList, nil
}

func isIncresing(n1, n2 int32) bool {
	return n1 < n2
}

func isDecresing(n1, n2 int32) bool {
	return n1 > n2
}

func isSafeReport(report []int32) bool {
	switch {
	case checkReport(report, isIncresing, 3):
		return true
	case checkReport(report, isDecresing, 3):
		return true
	default:
		return false
	}
}

func isSafeReportWithSingleBad(report []int32) bool {
	switch {
	case checkReportWithSingleBad(report, isIncresing, 3, 0, 1):
		return true
	case checkReportWithSingleBad(report, isDecresing, 3, 0, 1):
		return true
	case checkReportWithSingleBad(report, isIncresing, 3, 0, 0):
		return true
	case checkReportWithSingleBad(report, isDecresing, 3, 0, 0):
		return true
	default:
		return false
	}
}

func checkReport(report []int32, check func(int32, int32) bool, diference int32) bool {

	for i := 0; i < len(report)-1; i++ {
		if !check(report[i], report[i+1]) {
			return false
		}
		if int32(math.Abs(float64(report[i])-float64(report[i+1]))) > diference {
			return false
		}
	}

	return true
}

func checkReportWithSingleBad(report []int32, check func(int32, int32) bool, diference, attemp int32, rmIndex int) bool {
	if attemp > 1 {
		return false
	}

	tmp := make([]int32, len(report))
	copy(tmp, report)

	for i := 0; i < len(tmp)-1; i++ {

		if !check(tmp[i], tmp[i+1]) {
			tmp = remove(tmp, i+rmIndex)
			return checkReportWithSingleBad(tmp, check, diference, attemp+1, rmIndex)
		}
		if int32(math.Abs(float64(tmp[i])-float64(tmp[i+1]))) > diference {
			tmp = remove(tmp, i+rmIndex)
			return checkReportWithSingleBad(tmp, check, diference, attemp+1, rmIndex)
		}
	}

	return true
}

func safeAndUnsafeReports(reports [][]int32) (safe, unsafe int32) {
	for _, report := range reports {
		if isSafeReport(report) {
			safe++
			continue
		}
		unsafe++
	}
	return
}

func safeAndUnsafeReportsWithSingleBad(reports [][]int32) (safe, unsafe int32) {

	for _, report := range reports {
		if isSafeReportWithSingleBad(report) {
			safe++
			continue
		}
		unsafe++
	}
	return
}

func remove(slice []int32, s int) []int32 {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	reports, err := readReportSlice("input.txt")

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
		return
	}

	// day 2 part 1
	safe, _ := safeAndUnsafeReports(reports)
	fmt.Println("safe reports are:", safe)

	//day 2 part 2
	safeTolerant, _ := safeAndUnsafeReportsWithSingleBad(reports)
	fmt.Println("safe report with 1 tolerance are:", safeTolerant)
}
