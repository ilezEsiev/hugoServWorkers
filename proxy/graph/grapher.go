package graph

import (
	"fmt"
	"strings"
)

func generateMarkdownText() string {
	foos := getFuncMap()
	meta := getMetaMap()
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%v\n%v\n", meta["in"], meta["head"]))
	graph := GenerateGraph()
	for _, node := range graph {
		for _, link := range node.Links {
			builder.WriteString(fmt.Sprintf("%v%v --> %v%v\n", node.ID, foos[node.Form](node.Name), link.ID, foos[link.Form](link.Name)))
			deleteLink(node, link)

		}
	}
	builder.WriteString(meta["out"])
	return builder.String()
}

// Удаляет обратную связь на узел
func deleteLink(node, link *Node) {
	for i, corr := range link.Links {
		if node == corr {
			link.Links = append(link.Links[:i], link.Links[i+1:]...)
		}
	}

}

func getIn() string {
	return `---
menu:
    before:
        name: graph
        weight: 1
title: Построение графа
---
# Построение графа

{{< mermaid >}}`
}

func getOut() string {
	return `{{< /mermaid >}}`
}

func circle(label string) string {
	return fmt.Sprintf("((%v))", label)
}

func rect(label string) string {
	return fmt.Sprintf("[%v]", label)
}

func ellipse(label string) string {
	return fmt.Sprintf("([%v])", label)
}

func roundRect(label string) string {
	return fmt.Sprintf("(%v)", label)
}

func rhombus(label string) string {
	return fmt.Sprintf("{%v}", label)
}

func getFuncMap() map[string]func(string) string {
	outMap := map[string]func(string) string{
		"circle":     circle,
		"rect":       rect,
		"ellipse":    ellipse,
		"round-rect": roundRect,
		"rhombus":    rhombus,
	}
	return outMap
}
func getMetaMap() map[string]string {
	outMap := map[string]string{
		"in":   getIn(),
		"out":  getOut(),
		"head": `graph LR`,
	}
	return outMap
}
