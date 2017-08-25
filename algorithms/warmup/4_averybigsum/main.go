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
	strnum, _ := reader.ReadString('\n')

	fmt.Println(cal(strnum))

}

func cal(strnum string) string {

	listnum := strings.Split(strings.TrimSpace(strnum), " ")

	sum := 0

	for _, num := range listnum {
		n, _ := strconv.Atoi(num)
		sum += n
	}

	return strconv.Itoa(sum)

}
