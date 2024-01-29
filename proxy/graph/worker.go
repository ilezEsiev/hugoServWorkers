package graph

import (
	"log"
	"os"
	"time"
)

const (
	debugPathToGraphFile = "/home/ilez/hugoproxy/hugo/content/tasks/graph.md"
	appPathToGraphFile   = "/app/static/tasks/graph.md"
)

func GraphWorker() {
	t := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-t.C:
			err := os.WriteFile(appPathToGraphFile, []byte(generateMarkdownText()), 0644)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
