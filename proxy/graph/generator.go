package graph

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	ID    int
	Name  string
	Form  string // "circle", "rect", "square", "ellipse", "round-rect", "rhombus"
	Links []*Node
}

func GenerateGraph() []*Node {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	numNodes := rng.Intn(26) + 5
	graph := make([]*Node, numNodes)
	for i := 0; i < numNodes; i++ {
		graph[i] = &Node{
			ID:    i + 1,
			Name:  fmt.Sprintf("Node %v", i+1),
			Form:  getRandForm(),
			Links: []*Node{},
		}
	}

	for _, node := range graph {
		numLinks := rng.Intn(numNodes - 1)
		for j := 0; j < numLinks; j++ {
			linkedNode := graph[rng.Intn(numNodes)]
			if !checkLinkage(node.Links, linkedNode) && node != linkedNode {
				node.Links = append(node.Links, linkedNode)
			}
		}
	}
	return graph
}

func checkLinkage(nodes []*Node, item *Node) bool {
	for _, node := range nodes {
		if item == node {
			return true
		}
	}
	return false
}

func getRandForm() string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	forms := []string{"circle", "rect", "ellipse", "round-rect", "rhombus"} //"square"
	return forms[rng.Intn(len(forms))]
}
