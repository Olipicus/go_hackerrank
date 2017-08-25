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

	left, equal, right := partition(list, list[0])

	print := append(left, equal...)
	print = append(print, right...)

	strprint := fmt.Sprint(print)

	fmt.Println(strprint[1 : len(strprint)-1])

}

func partition(numlist []int, pivot int) (left []int, equal []int, right []int) {

	for _, num := range numlist {
		switch {
		case num < pivot:
			left = append(left, num)
		case num > pivot:
			right = append(right, num)
		case num == pivot:
			equal = append(equal, num)
		}
	}
	return

}

/*
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
*/
