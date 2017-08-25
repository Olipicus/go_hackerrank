package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	strnum, _ := reader.ReadString('\n')

	numline, _ := strconv.Atoi(strings.TrimSpace(strnum))

	numlist := []big.Int{}

	for i := 0; i < numline; i++ {
		strnum, _ := reader.ReadString('\n')
		bnum := big.NewInt(0)
		bnum, _ = bnum.SetString(strings.TrimSpace(strnum), 10)
		numlist = append(numlist, *bnum)
	}

	sort.Sort(BigInt(numlist))

	for _, num := range numlist {
		fmt.Println(num.String())
	}

}
