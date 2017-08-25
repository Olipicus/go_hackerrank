package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	time, _ := reader.ReadString('\n')

	fmt.Println(solution(time))
}

func solution(time string) string {

	r := regexp.MustCompile("^([0-9]|0[0-9]|1[0-2]):[0-5][0-9]:[0-5][0-9](PM|AM)$")
	match := r.FindStringSubmatch(time)

	if len(match) > 2 {
		switch match[2] {
		case "PM":
			hour, _ := strconv.Atoi(match[1])
			if match[1] != "12" {
				hour += 12
			}

			return fmt.Sprintf("%02d", hour) + time[2:8]
		case "AM":
			if match[1] != "12" {
				return time[:8]
			}

			return "00" + time[2:8]
		}
	}

	return ""
}
