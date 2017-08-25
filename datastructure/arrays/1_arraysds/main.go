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

	reverse(list)

	str := fmt.Sprint(list)
	fmt.Print(str[1 : len(str)-1])
}

func reverse(numlist []int) {
	i := 0
	j := len(numlist) - 1
	for i < j {
		numlist[i], numlist[j] = numlist[j], numlist[i]
		i++
		j--
	}

}
