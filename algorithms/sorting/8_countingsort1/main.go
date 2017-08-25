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

	countnum := make(map[int]int)

	for _, strnum := range liststrnum {
		num, _ := strconv.Atoi(strings.TrimSpace(strnum))

		if _, ok := countnum[num]; !ok {
			countnum[num] = 0
		}

		countnum[num]++
		list = append(list, num)
	}

	for i := 0; i < 100; i++ {
		if _, ok := countnum[i]; !ok {
			countnum[i] = 0
		}

		fmt.Printf("%d ", countnum[i])
	}

}
