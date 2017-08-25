package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value       string
	ConnectNode []*Node
}

type Graph struct {
	Edge []*Node
}

func (g *Graph) DFS(node *Node, find string, visited *map[string]bool) *Node {

	if node.Value == find {
		return node
	}

	v := *visited
	for _, childNode := range node.ConnectNode {
		if _, ok := v[childNode.Value]; !ok {

			v[childNode.Value] = true
			result := g.DFS(childNode, find, visited)

			if result != nil {
				return result
			}
		}
	}

	return nil
}

var found map[string]bool

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	listinput := strings.Split(input, " ")

	totalNode, _ := strconv.Atoi(strings.TrimSpace(listinput[0]))
	totalLink, _ := strconv.Atoi(strings.TrimSpace(listinput[1]))

	nodeList := make(map[int]*Node)

	for i := 0; i < totalNode; i++ {
		nodeList[i] = &Node{Value: strconv.Itoa(i)}
	}

	for j := 0; j < totalLink; j++ {
		input, _ := reader.ReadString('\n')
		listinput := strings.Split(input, " ")

		start, _ := strconv.Atoi(strings.TrimSpace(listinput[0]))
		end, _ := strconv.Atoi(strings.TrimSpace(listinput[1]))

		startNode := nodeList[start]
		endNode := nodeList[end]

		startNode.ConnectNode = append(startNode.ConnectNode, endNode)
		endNode.ConnectNode = append(endNode.ConnectNode, startNode)
	}

	graph := Graph{}

	//Add Node to Graph
	for _, node := range nodeList {
		graph.Edge = append(graph.Edge, node)
	}

	result := 0
	found = make(map[string]bool)

	for _, node := range graph.Edge {
		result += findUnlink(node, &graph)
	}

	fmt.Println(result)
}

func findUnlink(startNode *Node, graph *Graph) int {
	result := 0

	for _, node := range graph.Edge {

		start, _ := strconv.Atoi(startNode.Value)
		end, _ := strconv.Atoi(node.Value)

		var key string

		switch {
		case end < start:
			key = node.Value + ":" + startNode.Value
		default:
			key = startNode.Value + ":" + node.Value
		}

		if _, ok := found[key]; !ok {

			visit := make(map[string]bool)
			n := graph.DFS(startNode, node.Value, &visit)

			if n == nil {
				//log.Println("Not Found ", startNode.Value, "->", node.Value)
				found[key] = true
				result++
			}
		}
	}

	return result
}
