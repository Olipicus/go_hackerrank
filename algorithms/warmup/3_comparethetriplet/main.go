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
	alice, _ := reader.ReadString('\n')
	bob, _ := reader.ReadString('\n')

	fmt.Println(cal(alice, bob))

}

func cal(alice, bob string) string {
	listnuma, listnumb := strings.Split(strings.TrimSpace(alice), " "), strings.Split(strings.TrimSpace(bob), " ")
	ra, rb := 0, 0

	for i, _ := range listnuma {
		numa, _ := strconv.Atoi(listnuma[i])
		numb, _ := strconv.Atoi(listnumb[i])

		switch {
		case numa > numb:
			ra++
		case numb > numa:
			rb++
		}
	}

	return fmt.Sprintf("%d %d", ra, rb)

}
