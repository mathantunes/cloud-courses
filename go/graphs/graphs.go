package graphs

import "fmt"

type AdjacencyListGraph struct {
	vertices map[int][]int
}

func NewAdjacencyListGraph() *AdjacencyListGraph {
	return &AdjacencyListGraph{
		vertices: make(map[int][]int),
	}
}

func Mock() {

	graph := NewAdjacencyListGraph()
	graph.AddVertex(0)
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)

	graph.AddEdge(0, 1) // 0 --> [1]
	graph.AddEdge(1, 0) // 1 --> [0]

	graph.AddEdge(0, 2) // 0 --> [1 , 2]
	graph.AddEdge(2, 3) // 2 --> [3]
	graph.AddEdge(3, 1) // 3 --> [1]
	graph.AddEdge(3, 2) // 3 --> [1, 2]
	graph.AddEdge(3, 3) // 3 --> [1, 2, 3]

	graph.Print()
}

func (a *AdjacencyListGraph) AddVertex(k int) {
	a.vertices[k] = []int{}
}

func (a *AdjacencyListGraph) Print() {
	for k, v := range a.vertices {
		fmt.Printf("Vertex %v has edges with %v\n", k, v)
	}
}

func (a *AdjacencyListGraph) GetConnections(k int) []int {
	return a.vertices[k]
}

func (a *AdjacencyListGraph) AddEdge(from, to int) {
	currentList := a.vertices[from]
	for _, v := range currentList {
		if v == to {
			return
		}
	}
	currentList = append(currentList, to)
	a.vertices[from] = currentList
}
