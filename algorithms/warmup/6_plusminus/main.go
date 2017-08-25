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
	strnumline, _ := reader.ReadString('\n')
	strnumlist, _ := reader.ReadString('\n')

	fmt.Print(cal(strnumline, strnumlist))

}

func cal(strnumline, strnumlist string) string {

	countPos, countNeg, countZero := 0, 0, 0
	numline, _ := strconv.Atoi(strings.TrimSpace(strnumline))

	lnum := strings.Split(strnumlist, " ")

	for _, strnum := range lnum {
		num, _ := strconv.Atoi(strings.TrimSpace(strnum))

		switch {
		case num == 0:
			countZero++
		case num < 0:
			countNeg++
		case num > 0:
			countPos++
		}

	}

	return fmt.Sprintf("%f\n%f\n%f\n", float64(countPos)/float64(numline), float64(countNeg)/float64(numline), float64(countZero)/float64(numline))

}
