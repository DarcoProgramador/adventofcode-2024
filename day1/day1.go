package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"sync"
)

func readLists(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var list1, list2 []int
	for {
		var n1, n2 int
		_, err := fmt.Fscanf(file, "%d   %d\n", &n1, &n2)
		if err != nil {
			break
		}

		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}
	return list1, list2, nil
}

func sortAsync(list []int, wg *sync.WaitGroup) {
	sort.Ints(list)
	wg.Done()
}

func Part1() {
	// Reading the lists
	flag.Parse()
	list, originalList, err := readLists("list.txt")

	if err != nil {
		log.Fatal(err)
		fmt.Println("Error reading the lists")
		os.Exit(1)
		return
	}
	// Sorting the lists
	var wg sync.WaitGroup
	wg.Add(2)
	//Original list
	go sortAsync(list, &wg)
	// Other list
	go sortAsync(originalList, &wg)
	wg.Wait()
	// compare the distances and sum them
	sum := 0
	for i := 0; i < len(originalList); i++ {
		sum += int(math.Abs(float64(originalList[i]) - float64(list[i])))
	}

	fmt.Println("The total distances between your lists is:", sum)
}

func Part2() {
	list, originalList, err := readLists("list.txt")

	if err != nil {
		log.Fatal(err)
		fmt.Println("Error reading the lists")
		os.Exit(1)
		return
	}

	starMap := make(map[int]int)

	for _, star := range originalList {
		starMap[star]++
	}

	sum := 0
	for _, similarityNum := range list {
		sum += similarityNum * starMap[similarityNum]
	}

	fmt.Println("The the similarity score is:", sum)
}

var (
	partFlag = flag.Int("part", 1, "The part of the challenge code")
)

func main() {
	flag.Parse()

	switch *partFlag {
	case 1:
		Part1()
	case 2:
		Part2()
	default:
		fmt.Println("No existe ese numero de parte del reto")
	}
}
