package main

import (
	"testing"
)

func TestCircleGraph(t *testing.T) {

	graph := Graph{}

	A := Node{Value: "A"}
	B := Node{Value: "B"}
	C := Node{Value: "C"}
	D := Node{Value: "D"}
	E := Node{Value: "E"}

	A.ConnectNode = append(A.ConnectNode, &B)
	B.ConnectNode = append(B.ConnectNode, &C)
	C.ConnectNode = append(C.ConnectNode, &A)
	C.ConnectNode = append(C.ConnectNode, &D)
	C.ConnectNode = append(C.ConnectNode, &E)

	graph.Edge = append(graph.Edge, &A)
	graph.Edge = append(graph.Edge, &B)
	graph.Edge = append(graph.Edge, &C)
	graph.Edge = append(graph.Edge, &D)
	graph.Edge = append(graph.Edge, &E)

	visited := make(map[string]bool)
	resultNode := graph.DFS(&A, "C", &visited)

	if resultNode == nil {
		t.Fatalf("It should found Node C")
	}

	r := *resultNode
	if r.Value != "C" {
		t.Errorf("Value of result should C but got %s", resultNode.Value)
	}

	visited = make(map[string]bool)

	resultNode = graph.DFS(&A, "E", &visited)

	if resultNode == nil {
		t.Fatalf("It should found Node E")
	}

	r = *resultNode
	if r.Value != "E" {
		t.Errorf("Value of result should E but got %s", resultNode.Value)
	}

}

func TestTreeGraph(t *testing.T) {

	graph := Graph{}

	A := Node{Value: "A"}
	B := Node{Value: "B"}
	C := Node{Value: "C"}
	D := Node{Value: "D"}
	E := Node{Value: "E"}

	A.ConnectNode = append(A.ConnectNode, &B)
	A.ConnectNode = append(A.ConnectNode, &C)
	B.ConnectNode = append(B.ConnectNode, &D)
	D.ConnectNode = append(D.ConnectNode, &E)

	graph.Edge = append(graph.Edge, &A)
	graph.Edge = append(graph.Edge, &B)
	graph.Edge = append(graph.Edge, &C)
	graph.Edge = append(graph.Edge, &D)
	graph.Edge = append(graph.Edge, &E)

	visited := make(map[string]bool)
	resultNode := graph.DFS(&A, "C", &visited)

	if resultNode == nil {
		t.Fatalf("It should found Node C")
	}

	r := *resultNode
	if r.Value != "C" {
		t.Errorf("Value of result should C but got %s", resultNode.Value)
	}

	visited = make(map[string]bool)

	resultNode = graph.DFS(&A, "E", &visited)

	if resultNode == nil {
		t.Fatalf("It should found Node E")
	}

	r = *resultNode
	if r.Value != "E" {
		t.Errorf("Value of result should E but got %s", resultNode.Value)
	}

}

func TestUnlinkGraph(t *testing.T) {
	graph := Graph{}

	A := Node{Value: "A"}
	B := Node{Value: "B"}
	C := Node{Value: "C"}
	D := Node{Value: "D"}
	E := Node{Value: "E"}

	//Make Circle Graph
	A.ConnectNode = append(A.ConnectNode, &B)
	B.ConnectNode = append(B.ConnectNode, &C)
	C.ConnectNode = append(C.ConnectNode, &A)

	D.ConnectNode = append(D.ConnectNode, &E)

	graph.Edge = append(graph.Edge, &A)
	graph.Edge = append(graph.Edge, &B)
	graph.Edge = append(graph.Edge, &C)
	graph.Edge = append(graph.Edge, &D)
	graph.Edge = append(graph.Edge, &E)

	visited := make(map[string]bool)
	resultNode := graph.DFS(&A, "D", &visited)

	if resultNode != nil {
		t.Fatalf("It should not found Node %s", resultNode.Value)
	}

	visited = make(map[string]bool)

	resultNode = graph.DFS(&E, "A", &visited)

	if resultNode != nil {
		t.Fatalf("It should not found Node %s", resultNode.Value)
	}

}
