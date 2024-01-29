package graph

import (
	"strings"
	"testing"
)

func TestGenerateMarkdownText(t *testing.T) {
	expectedPrefix := `---
        menu:
            before:
                name: graph
                weight: 1
        title: Построение графа
        ---
        # Построение графа
        
        {{< mermaid >}}
        graph LR`
	expectedSuffix := `{{< /mermaid >}}`
	output := generateMarkdownText()

	// Проверяем, что сгенерированный текст соответствует ожиданиям
	if strings.HasPrefix(output, expectedPrefix) && strings.HasSuffix(output, expectedSuffix) {
		t.Errorf("Итоговая строка не соответствует общему виду")
	}
}
