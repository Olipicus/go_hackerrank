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

	//list := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 1}

	InsertionSort(list)

}

func InsertionSort(unsorted []int) []int {

	list := make([]int, len(unsorted))
	for i := len(unsorted) - 1; i > 0; i-- {
		j := i

		tmpstr := fmt.Sprintln(list)
		copy(list, unsorted)

		for j < len(unsorted) && unsorted[j-1] > unsorted[j] {
			unsorted[j-1], unsorted[j], list[j] = unsorted[j], unsorted[j-1], unsorted[j-1]
			j++

		}

		if str := fmt.Sprintln(list); str != tmpstr {
			tmpstr = str
			fmt.Println(str[1 : len(str)-2])
		}
	}

	strsort := fmt.Sprintln(unsorted)
	strlist := fmt.Sprintln(list)

	if strsort != strlist {
		fmt.Println(strsort[1 : len(strsort)-2])
	}

	return unsorted

}
