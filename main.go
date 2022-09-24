package main

import (
	"container/list"
	"log"
	"os"
	"strings"
	"golang.org/x/exp/slices"
)

func parse(contents string) *list.List {
	entries := make([]string, 0)
	lines := strings.Split(strings.TrimSpace(contents), "\n")
	for _, line := range lines {
		str := strings.TrimSpace(line)
		if len(str) > 0 {
			entries = append(entries, str)
		}
	}
	return entries
}

func main() {
	ignore, err := os.Open(".gitignore")
	if err != nil {
		log.Fatal("not found.")
	}
	defer ignore.Close()

	b := make([]byte, 1024)
	nread, err := ignore.Read(b)
	if err != nil {
		log.Fatal("read fail.")
	}

	entries := parse(string(b[:nread]))
	slices.Sort(words)
	log.Println(words)
}
