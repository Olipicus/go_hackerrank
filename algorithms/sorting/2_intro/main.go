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
	strfind, _ := reader.ReadString('\n')

	find, _ := strconv.Atoi(strings.TrimSpace(strfind))

	reader.ReadString('\n')

	strinput, _ := reader.ReadString('\n')

	strnumlist := strings.Split(strinput, " ")
	numlist := []int{}
	for _, strnum := range strnumlist {
		num, _ := strconv.Atoi(strings.TrimSpace(strnum))
		numlist = append(numlist, num)
	}

	fmt.Println(solution(find, numlist))
}

func solution(find int, numlist []int) int {

	for i, num := range numlist {
		if find == num {
			return i
		}
	}

	return 0
}
