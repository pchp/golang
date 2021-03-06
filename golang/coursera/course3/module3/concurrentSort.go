package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func StringSliceToStringInt(sli []string) []int {

	sliInt := make([]int, len(sli))
	for i := 0; i < len(sli); i++ {

		s, err := strconv.Atoi(sli[i])

		if err == nil {
			sliInt[i] = s
		}
	}
	return sliInt
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	numberArray := 4
	slices := make([][]int, 4)
	result := []int{}

	ch := make(chan []int)

	fmt.Print("> (Enter a list of integer - minimum 7 ints - sparated by space): ")
	input, _ := reader.ReadString('\n')
	input = strings.Trim(input, "\n")
	input = strings.ToLower(input)

	sli := StringSliceToStringInt(strings.Split(input, " "))

	//fmt.Println(sli)
	chunkSize := (len(sli) + numberArray - 1) / numberArray

	for i := 0; i < len(sli); i += chunkSize {
		end := i + chunkSize
		if end > len(sli) {
			end = len(sli)
		}

		slices = append(slices, sli[i:end])
	}

	slices = slices[4:]
	//fmt.Println(slices)
	if len(slices) < 4 {
		fmt.Println("only ", len(slices), " subarray can be built from your list: ", slices)
		os.Exit(0)
	}

	for i := 0; i < len(slices); i++ {
		go func(sli []int, i int, ch chan<- []int) {
			sort.Sort(sort.IntSlice(sli))
			fmt.Println("sorted subarray #", i, sli)
			ch <- sli
		}(slices[i], i, ch)
	}

	for i := 0; i < numberArray; i++ {
		result = append(result, <-ch...)
	}

	fmt.Println("====== result (sorted array) =====")
	sort.Sort(sort.IntSlice(result))
	fmt.Println(result)
}
