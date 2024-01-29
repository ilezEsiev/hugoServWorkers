package graph

import (
	"testing"
)

func TestGenerateGraph(t *testing.T) {
	// Тестируем генерацию графа
	graph := GenerateGraph()

	// Проверяем, что количество узлов в графе соответствует ожиданиям
	expectedNumNodes := len(graph)
	if expectedNumNodes < 5 || expectedNumNodes > 30 {
		t.Errorf("Unexpected number of nodes in the generated graph: %v", expectedNumNodes)
	}

	// Проверяем, что у каждого узла есть уникальный ID
	idSet := make(map[int]bool)
	for _, node := range graph {
		if idSet[node.ID] {
			t.Errorf("Duplicate node ID found: %v", node.ID)
		}
		idSet[node.ID] = true
	}

	// Проверяем, что у каждого узла есть форма
	for _, node := range graph {
		if node.Form == "" {
			t.Errorf("Node %v has no form specified", node.ID)
		}
	}
}

func TestGraphLinkage(t *testing.T) {
	// Создаем тестовый граф
	node1 := &Node{ID: 1, Name: "Node 1", Form: "circle", Links: []*Node{}}
	node2 := &Node{ID: 2, Name: "Node 2", Form: "rect", Links: []*Node{}}
	node3 := &Node{ID: 3, Name: "Node 3", Form: "ellipse", Links: []*Node{}}

	// Соединяем узлы
	node1.Links = append(node1.Links, node2)
	node1.Links = append(node1.Links, node3)

	// Проверяем, что узлы связаны правильно
	if !checkLinkage(node1.Links, node2) || !checkLinkage(node1.Links, node3) {
		t.Error("Nodes are not linked correctly")
	}

	// Проверяем, что узлы не связаны неправильно
	if checkLinkage(node2.Links, node3) {
		t.Error("Nodes are linked incorrectly")
	}
}
