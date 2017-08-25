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
	strnum, _ := reader.ReadString('\n')

	liststrnum := strings.Split(strnum, " ")
	listnum := []int{}

	for _, num := range liststrnum {
		if n, err := strconv.Atoi(strings.TrimSpace(num)); err == nil {
			listnum = append(listnum, n)
		}
	}

	min, max := solution(listnum)
	fmt.Printf("%d %d", min, max)

}

func solution(numlist []int) (int, int) {
	min, max := 9223372036854775807, 0
	for i, _ := range numlist {
		l := []int{}
		left := numlist[0:i]
		right := numlist[i+1:]

		l = append(l, left...)
		l = append(l, right...)

		sum := sum(l)

		if sum < min {
			min = sum
		}

		if sum > max {
			max = sum
		}

	}

	return min, max
}

func sum(s []int) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return s[0]
	}

	return s[0] + sum(s[1:])
}
