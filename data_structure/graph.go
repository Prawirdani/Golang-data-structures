package data_structure

import (
	"fmt"
	"os"
	"sort"
)

type graphImpl struct {
	matrix_adj_list map[string][]string
}

type graph interface {
	AddVertex(val string)
	AddEdges(vertex string, toVertex string)
	Print()
}

func newGraph() graph {
	return &graphImpl{
		make(map[string][]string),
	}
}

func (g *graphImpl) AddVertex(val string) {
	g.matrix_adj_list[val] = []string{}
}

func (g *graphImpl) AddEdges(vertex string, toVertex string) {
	_, ok := g.matrix_adj_list[vertex]
	if !ok {
		fmt.Printf("Vertex \"%v\" not exist\n", vertex)
		os.Exit(1)
	}
	_, ok = g.matrix_adj_list[toVertex]
	if !ok {
		fmt.Printf("Vertex \"%v\" not exist\n", toVertex)
		os.Exit(1)
	}
	for each := range g.matrix_adj_list {
		if each == vertex {
			g.matrix_adj_list[vertex] = append(g.matrix_adj_list[vertex], toVertex)
			g.matrix_adj_list[toVertex] = append(g.matrix_adj_list[toVertex], each)
		}
	}
}

func (g *graphImpl) Print() {
	for vertex, edges := range g.matrix_adj_list {
		fmt.Printf("%v: %v\n", vertex, edges)
	}

	vertices := make([]string, 0)
	for vertex := range g.matrix_adj_list {
		vertices = append(vertices, vertex)
	}
	sort.Slice(vertices, func(i, j int) bool {
		return fmt.Sprintf("%v", vertices[i]) < fmt.Sprintf("%v", vertices[j])
	})

	fmt.Print("  ")
	for _, vertex := range vertices {
		fmt.Printf("%v ", vertex)
	}
	fmt.Println()

	for _, vertex := range vertices {
		fmt.Printf("%v ", vertex)
		for _, adjVertex := range vertices {
			adjacent := false
			for _, v := range g.matrix_adj_list[vertex] {
				if v == adjVertex {
					adjacent = true
					break
				}
			}
			if adjacent {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
}

func Graph() {
	g := newGraph()
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")

	g.AddEdges("D", "A")
	g.AddEdges("D", "B")
	g.AddEdges("D", "C")
	g.Print()
}
