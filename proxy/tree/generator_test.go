package tree

import (
	"testing"
)

func TestAVLTreeInsert(t *testing.T) {
	tree := NewAVLTree()

	// Проверяем, что дерево не пустое
	if tree.Root == nil {
		t.Error("Tree is empty after insertion")
	}

	// Проверяем, что высота корня правильная
	if tree.Root.Height != 3 {
		t.Errorf("Root height is incorrect, expected 3, got %d", tree.Root.Height)
	}

	// Проверяем, что вставка элементов работает правильно
	values := []int{5, 6, 7}
	for _, value := range values {
		tree.Insert(value)
		if !isAVLTreeBalanced(tree.Root) {
			t.Error("Tree is not balanced after insertion")
		}
	}
}

// Функция для проверки сбалансированности AVL-дерева
func isAVLTreeBalanced(node *Node) bool {
	if node == nil {
		return true
	}

	balance := getBalance(node)
	if balance > 1 || balance < -1 {
		return false
	}

	return isAVLTreeBalanced(node.Left) && isAVLTreeBalanced(node.Right)
}
