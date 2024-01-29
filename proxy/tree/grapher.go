package tree

import (
	"fmt"
	"strings"
)

func generateMarkdownText(tree *AVLTree) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%v\n%v\n", getIn(), "graph TD"))
	dfs(tree.Root, &builder)
	builder.WriteString(getOut())
	return builder.String()
}

func dfs(node *Node, builder *strings.Builder) {
	if node != nil {
		if node.Left != nil {
			builder.WriteString(fmt.Sprintf("%v --> %v\n", node.Key, node.Left.Key))
			dfs(node.Left, builder)
		}
		if node.Right != nil {
			builder.WriteString(fmt.Sprintf("%v --> %v\n", node.Key, node.Right.Key))
			dfs(node.Right, builder)
		}
	}
}

func getIn() string {
	return `---
menu:
    before:
        name: binary
        weight: 5
title: Построение бинарного дерева
---
# Построение бинарного дерева

{{< mermaid >}}`
}

func getOut() string {
	return `{{< /mermaid >}}`
}
