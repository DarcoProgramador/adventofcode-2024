package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func readFile(filename string) (string, error) {
	file, err := os.ReadFile(filename)

	return string(file), err
}

func proceesMul(strMul string) int {
	r, _ := regexp.Compile("\\d+")

	strNums := r.FindAllString(strMul, -1)

	num1, _ := strconv.Atoi(strNums[0])
	num2, _ := strconv.Atoi(strNums[1])

	return num1 * num2
}

func obtainMulStr(data string) []string {
	r := regexp.MustCompile(`(do\(\))|(don't\(\))|(mul\(\d+,\d+\))`)

	mulStrSlice := r.FindAllString(data, -1)

	do := true

	var mulDoSlice []string

	for _, doStr := range mulStrSlice {
		switch doStr {
		case "do()":
			do = true
		case "don't()":
			do = false

		default:
			if do {
				mulDoSlice = append(mulDoSlice, doStr)
			}
		}
	}
	return mulDoSlice
}

func main() {
	data, err := readFile("input2.txt")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
		return
	}

	mulStrSlice := obtainMulStr(data)

	var sum int
	for _, strMul := range mulStrSlice {
		sum += proceesMul(strMul)
	}

	fmt.Println("The sum of all multiplications are:", sum)
}
