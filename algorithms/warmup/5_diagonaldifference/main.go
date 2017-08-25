package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	strnum, _ := reader.ReadString('\n')

	numline, _ := strconv.Atoi(strings.TrimSpace(strnum))

	strline := make([]string, numline)

	for i := 0; i < numline; i++ {
		str, _ := reader.ReadString('\n')
		strline[i] = strings.TrimSpace(str)
	}

	fmt.Println(cal(strline))

}

func cal(strline []string) int {
	a, b := 0, 0
	for i, str := range strline {
		listnum := strings.Split(str, " ")
		numleft, _ := strconv.Atoi(listnum[i])
		numright, _ := strconv.Atoi(listnum[len(strline)-i-1])

		log.Printf("%d %d", numleft, numright)
		a += numleft
		b += numright
	}
	return int(math.Abs(float64(a - b)))
}
