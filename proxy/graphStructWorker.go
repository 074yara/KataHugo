package main

import (
	"math/rand"
	"os"
	"time"
)

var graphHead = `---
menu:
    after:
        name: graph
        weight: 1
title: Построение графа
---
`

func GraphStructWorker() {
	graph := GenerateRandomGraph(rand.Intn(10) + 5)
	mermaidString := GenerateMermaidCode(graph)
	code := graphHead + mermaidString
	err := os.WriteFile(graphFilePath, []byte(code), 0644)
	checkError(err)
	for {
		time.Sleep(time.Second * 5)
		graph = GenerateRandomGraph(rand.Intn(10) + 5)
		mermaidString = GenerateMermaidCode(graph)
		code = graphHead + mermaidString
		err = os.WriteFile(graphFilePath, []byte(code), 0644)
		checkError(err)
	}
}
