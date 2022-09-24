package main

import (
	"container/list"
	"log"
	"os"
	"strings"
	//	"bytes"
)

func parse(contents string) *list.List {
	entries := list.New()
	lines := strings.Split(strings.TrimSpace(contents), "\n")
	for _, line := range lines {
		str := strings.TrimSpace(line)
		if len(str) > 0 {
			entries.PushBack(str)
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
	words := make([]string, 0)
	for e := entries.Front(); e != nil; e = e.Next() {
		if e.Value != "" {
			word := e.Value.(string)
			words = append(words, word)
			log.Println(word)
		}
	}

}
