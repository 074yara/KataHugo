package main

import (
	"fmt"
	"math/rand"
)

type NodeGraph struct {
	ID    int
	Name  string
	Form  string // "circle", "rect", "square", "ellipse", "round-rect", "rhombus"
	Links []*NodeGraph
}

func GenerateRandomGraph(nodeCount int) *NodeGraph {
	var graph []*NodeGraph
	//Generate nodes
	for i := 1; i <= nodeCount; i++ {
		node := &NodeGraph{
			ID:    i,
			Name:  fmt.Sprintf("Node_%v", i),
			Form:  forms[rand.Intn(len(forms))],
			Links: []*NodeGraph{},
		}
		graph = append(graph, node)
	}
	//Connecting nodes
	for _, node := range graph {
		linksCount := rand.Intn(nodeCount)
		for j := 0; j < linksCount; j++ {
			linkedNode := getRandomNode(graph, node)
			node.Links = append(node.Links, linkedNode)
		}

	}
	return graph[rand.Intn(len(graph))]
}

func getRandomNode(graph []*NodeGraph, excludeNode *NodeGraph) *NodeGraph {
	for {
		node := graph[rand.Intn(len(graph))]
		if node != excludeNode {
			return node
		}
	}
}

func printGraph(start *NodeGraph) {
	fmt.Println("Generated graph:")
	printNode(start, make(map[int]struct{}))
}

func printNode(node *NodeGraph, visited map[int]struct{}) {
	if _, exists := visited[node.ID]; exists {
		return
	}
	fmt.Printf("Node %v (Name: %v Form: %v)\n", node.ID, node.Name, node.Form)
	visited[node.ID] = struct{}{}
	for _, linkedNode := range node.Links {
		printNode(linkedNode, visited)
	}
}

func GenerateMermaidCode(graph *NodeGraph) string {
	mermaidCode := "{{< mermaid >}}\ngraph LR\n"
	visited := make(map[int]struct{})
	generateMermaidRec(graph, visited, &mermaidCode)
	mermaidCode += "{{< /mermaid >}}"
	return mermaidCode
}

func generateMermaidRec(node *NodeGraph, visited map[int]struct{}, mermaidCode *string) {
	if _, exists := visited[node.ID]; exists {
		return
	}
	visited[node.ID] = struct{}{}
	leftBR, rightBR := bracketChoice(node)
	*mermaidCode += fmt.Sprintf("  %v%v%v, %v%v\n", node.ID, leftBR, node.Name, node.Form, rightBR)
	for _, linkedNode := range node.Links {
		*mermaidCode += fmt.Sprintf("  %d --> %d\n", node.ID, linkedNode.ID)
		generateMermaidRec(linkedNode, visited, mermaidCode)
	}
}

func bracketChoice(node *NodeGraph) (string, string) {
	switch node.Form {
	case "circle":
		return "((", "))"
	case "rect":
		return "[", "]"
	case "square":
		return "[", "]"
	case "round-rect":
		return "(", ")"
	default:
		return "{", "}"
	}
}
