package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	strnum, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strnum)

	fmt.Print(cal(num))
}

func cal(num int) string {
	result := ""
	for i := num; i > 0; i-- {
		for j := 1; j <= num; j++ {
			switch {
			case j >= i:
				result += "#"
			default:
				result += " "
			}
		}

		result += "\n"
	}

	return result
}
