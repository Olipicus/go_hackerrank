package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	inputnum, _ := reader.ReadString('\n')

	strnumlist := strings.Split(inputnum, " ")

	numlist := []int{}

	for _, strnum := range strnumlist {
		num, _ := strconv.Atoi(strings.TrimSpace(strnum))

		numlist = append(numlist, num)
	}

	fmt.Println(solution(numlist))
}

func solution(numlist []int) int {
	sort.Ints(numlist)
	max, count := 0, 0

	if len(numlist) > 0 {
		max = numlist[len(numlist)-1]
	}

	for _, num := range numlist {
		if num == max {
			count++
		}
	}

	return count

}
