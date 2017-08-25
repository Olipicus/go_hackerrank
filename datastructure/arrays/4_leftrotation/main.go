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
	strinput, _ := reader.ReadString('\n')

	listinput := strings.Split(strinput, " ")

	strnumlist, _ := reader.ReadString('\n')
	liststrnum := strings.Split(strnumlist, " ")
	list := []int{}

	for _, strnum := range liststrnum {
		num, _ := strconv.Atoi(strings.TrimSpace(strnum))
		list = append(list, num)
	}

	round, _ := strconv.Atoi(strings.TrimSpace(listinput[1]))

	result := rotate(list, round)

	str := fmt.Sprint(result)
	fmt.Print(str[1 : len(str)-1])

}

func rotate(numlist []int, round int) []int {
	if round < len(numlist) {
		numlist = append(numlist[round:], numlist[:round]...)
	}

	return numlist
}
