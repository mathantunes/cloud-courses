package graphs

import (
	"testing"
)

func TestAdjacencyListGraph_AddVertex(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		a    *AdjacencyListGraph
		args args
	}{
		{
			name: "first",
			a:    NewAdjacencyListGraph(),
			args: args{k: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.AddVertex(tt.args.k)
			list := tt.a.GetConnections(tt.args.k)
			if list == nil || len(list) != 0 {
				t.Fatal()
			}
		})
	}
}

func TestAdjacencyListGraph_Print(t *testing.T) {
	tests := []struct {
		name string
		a    *AdjacencyListGraph
	}{
		{
			name: "first",
			a: &AdjacencyListGraph{
				vertices: map[int][]int{
					1: []int{2, 3},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.Print()
		})
	}
}

func TestAdjacencyListGraph_AddEdge(t *testing.T) {
	type args struct {
		from int
		to   int
	}
	tests := []struct {
		name string
		a    *AdjacencyListGraph
		args args
	}{
		{
			name: "first",
			a: &AdjacencyListGraph{
				vertices: map[int][]int{
					1: []int{},
				},
			},
			args: args{
				from: 1,
				to:   2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.AddEdge(tt.args.from, tt.args.to)
			currentList := tt.a.GetConnections(tt.args.from)
			ok := false
			for _, v := range currentList {
				if v == tt.args.to {
					ok = true
				}
			}
			if !ok {
				t.Fatal()
			}
		})
	}
}

func TestMock(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "MOCK",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Mock()
		})
	}
}
