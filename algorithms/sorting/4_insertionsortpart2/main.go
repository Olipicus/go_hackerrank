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

	InsertionSort(list)

}

func InsertionSort(unsorted []int) []int {

	list := make([]int, len(unsorted))
	for i := 1; i < len(unsorted); i++ {
		j := i

		copy(list, unsorted)

		for j > 0 && unsorted[j-1] > unsorted[j] {
			unsorted[j-1], unsorted[j], list[j] = unsorted[j], unsorted[j-1], unsorted[j-1]
			j--
		}

		str := fmt.Sprintln(unsorted)
		fmt.Println(str[1 : len(str)-2])

	}

	return unsorted

}
