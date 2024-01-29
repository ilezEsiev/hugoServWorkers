package tree

import (
	"log"
	"os"
	"time"
)

const (
	debugPathToTreeFile = "/home/ilez/hugoproxy/hugo/content/tasks/field.md"
	appPathToTreeFile   = "/app/static/tasks/binary.md"
)

func TreeWorker() {
	tree := NewAVLTree()
	count := 5
	t := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-t.C:
			err := os.WriteFile(appPathToTreeFile, []byte(generateMarkdownText(tree)), 0644)
			if err != nil {
				log.Println(err)
			}
			if count >= 100 {
				tree = NewAVLTree()
				count = 5

			} else {
				count++
				tree.Insert(count)
			}
		}
	}
}
