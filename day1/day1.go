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

var (
	flagList = flag.String("list", "list.txt", "The file with the lists")
)

func main() {
	// Reading the lists
	flag.Parse()
	list, originalList, err := readLists(*flagList)

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
