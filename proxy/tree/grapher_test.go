package tree

import (
	"strings"
	"testing"
)

func TestGenerateMarkdownText(t *testing.T) {
	expectedPrefix := `---
        menu:
            before:
                name: binary
                weight: 5
        title: Построение бинарного дерева
        ---
        # Построение бинарного дерева
        
        {{< mermaid >}}
		graph TD`
	expectedSuffix := `{{< /mermaid >}}`
	tree := NewAVLTree()
	output := generateMarkdownText(tree)

	// Проверяем, что сгенерированный текст соответствует ожиданиям
	if strings.HasPrefix(output, expectedPrefix) && strings.HasSuffix(output, expectedSuffix) {
		t.Errorf("Итоговая строка не соответствует общему виду")
	}
}
