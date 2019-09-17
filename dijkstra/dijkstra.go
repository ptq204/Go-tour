package main

import (
	"fmt"
	"math"
)

type edge struct {
	node   string
	weight int
}

type graph struct {
	V        int
	vertices map[string][]edge
}

func createGraph() *graph {
	return &graph{vertices: make(map[string][]edge)}
}

func (g *graph) addEdge(src, des string, w int) {
	g.vertices[src] = append(g.vertices[src], edge{node: des, weight: w})
	g.vertices[des] = append(g.vertices[des], edge{node: src, weight: w})
}

func (g *graph) getEdges(node string) []edge {
	return g.vertices[node]
}

func findMinNode(dist map[string]int, visited map[string]bool) string {
	var minNode = ""
	var minValue = math.MaxInt64
	for k, v := range dist {
		//fmt.Println("value at %s is %d", k, v)
		if minValue > v && visited[k] == false {
			//fmt.Println("OOOOOOOOOOOO")
			minNode = k
			minValue = v
		}
	}
	//fmt.Println("min node is: " + minNode)
	return minNode
}

func (g *graph) findPath(src, des string) int {

	dist := make(map[string]int, g.V)
	visited := make(map[string]bool)

	for k := range g.vertices {
		visited[k] = false
		dist[k] = math.MaxInt64
	}

	dist[src] = 0

	var i = 0
	for i = 0; i < len(g.vertices); i++ {

		//fmt.Println("Processing at node " + k)
		u := findMinNode(dist, visited)

		visited[u] = true

		if u == des {
			return dist[u]
		}

		edgeList := g.getEdges(u)
		//fmt.Println(edgeList)

		for _, e := range edgeList {
			// fmt.Println(u)
			if visited[e.node] == false && dist[u]+e.weight < dist[e.node] {
				dist[e.node] = e.weight + dist[u]
			}
		}
	}

	fmt.Println(dist)

	return 0
}
