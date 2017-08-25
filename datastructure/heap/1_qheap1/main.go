package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Heap struct {
	data []int64
}

func (h *Heap) Peek() int64 {
	if len(h.data) <= 0 {
		return -1
	}
	return h.data[0]
}

func (h *Heap) Add(num int64) {
	h.data = append(h.data, num)
	h.heapifyUp()
}

func (h *Heap) Delete(num int64) {

	for index, _ := range h.data {
		if num == h.data[index] {

			if len(h.data) > 1 {
				h.data[index] = h.data[len(h.data)-1]
				h.data = h.data[:len(h.data)-1]
				break
			}

			h.data = []int64{}
			break
		}
	}

	h.heapifyDown()
}

func (h *Heap) Poll() int64 {

	if len(h.data) <= 0 {
		return -1
	}

	val := h.data[0]
	h.data = h.data[1:]

	h.heapifyDown()

	return val
}

func (h *Heap) heapifyUp() {
	index := len(h.data) - 1
	parent := (index - 1) / 2

	for index > 0 && h.data[index] < h.data[parent] {

		h.data[parent], h.data[index] = h.data[index], h.data[parent]
		index = parent
		parent = (index - 1) / 2
	}
}

func (h *Heap) heapifyDown() {
	index := 0
	left := index*2 + 1
	right := index*2 + 2

	for len(h.data) > left {
		small := left
		if len(h.data) > right && h.data[right] < h.data[left] {
			small = right
		}

		if h.data[index] > h.data[small] {
			h.data[index], h.data[small] = h.data[small], h.data[index]
		}
		index = small
		left = index*2 + 1
		right = index*2 + 2
	}

}

func NewHeap() *Heap {
	return &Heap{data: []int64{}}
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	strnum, _ := reader.ReadString('\n')

	num, _ := strconv.Atoi(strings.TrimSpace(strnum))

	h := NewHeap()

	listtop := []int64{}

	for i := 0; i < num; i++ {
		cmd, _ := reader.ReadString('\n')

		listcmd := strings.Split(cmd, " ")

		var num int64
		if len(listcmd) > 1 {
			num, _ = strconv.ParseInt(strings.TrimSpace(listcmd[1]), 10, 64)
		}

		switch strings.TrimSpace(listcmd[0]) {
		case "1":
			h.Add(num)
		case "2":
			h.Delete(num)
		case "3":
			listtop = append(listtop, h.Peek())
		}

	}

	for _, top := range listtop {
		fmt.Println(top)
	}

}
