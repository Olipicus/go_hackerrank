package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	result := Sub(strings.TrimSpace(input))

	if result == "" {
		result = "Empty String"
	}

	fmt.Println(result)
}

func Sub(s string) string {
	if len(s) == 1 {
		return s
	}

	if len(s) == 0 {
		return ""
	}

	if s[0] == s[1] {
		return Sub(s[2:])
	}

	r := Sub(s[0:1]) + Sub(s[1:])

	if s == r {
		return s
	}

	return Sub(r)
}
