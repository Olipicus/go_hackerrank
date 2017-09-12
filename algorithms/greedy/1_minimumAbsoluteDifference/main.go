package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var minVal float64

func main() {

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	strnumlist, _ := reader.ReadString('\n')
	liststrnum := strings.Split(strnumlist, " ")
	list := []float64{}

	for _, strnum := range liststrnum {
		num, _ := strconv.ParseFloat(strings.TrimSpace(strnum), 64)
		list = append(list, num)
	}

	fmt.Println(Solution(list))

}

func Solution(numlist []float64) float64 {

	sort.Float64s(numlist)

	var minVal float64
	minVal = 1000000000

	for i := 1; i < len(numlist); i++ {

		result := math.Abs(numlist[i] - numlist[i-1])
		//log.Printf("%v - %v = %v\n", numlist[i], numlist[i-1], result)
		minVal = math.Min(minVal, result)
	}

	return minVal
}
