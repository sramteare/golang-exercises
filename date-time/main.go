package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	queue := []string{"."}

	// While the queue is not empty, do the following:
	for len(queue) > 0 {
		// Pop the first node from the queue.
		node := queue[0]
		queue = queue[1:]

		// Visit the node.
		fmt.Println(node)
		files, err := os.ReadDir(node)
		if err != nil {
			log.Println(err)
			continue
		}

		// For each child of the node, do the following:
		for _, child := range files {
			// If the child is a directory, add it to the queue.
			if child.IsDir() {
				queue = append(queue, child.Name())
			}
		}
	}
}
