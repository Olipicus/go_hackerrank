package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	strnumlist, _ := reader.ReadString('\n')
	liststrnum := strings.Split(strnumlist, " ")
	list := []int{}

	for _, strnum := range liststrnum {
		num, _ := strconv.Atoi(strings.TrimSpace(strnum))
		list = append(list, num)
	}

	QuickSort(list)

	str := fmt.Sprint(list)
	fmt.Print(str[1 : len(str)-1])

}

func QuickSort(unsorted []int) {

	if len(unsorted) <= 1 {
		return
	}

	index := hoarePartition(unsorted)

	QuickSort(unsorted[:index+1])
	QuickSort(unsorted[index+1:])

}

func hoarePartition(unsorted []int) int {
	if len(unsorted) <= 1 {
		return 0
	}

	pivot := unsorted[len(unsorted)/2]
	left := 0
	right := len(unsorted) - 1

	for left <= right {

		for unsorted[left] < pivot {
			left++
		}

		for unsorted[right] > pivot {

			right--
		}

		if left <= right {
			unsorted[left], unsorted[right] = unsorted[right], unsorted[left]
			left++
			right--
		}
	}

	return right

}
