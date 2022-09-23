package main

import (
	"container/list"
	"log"
	"os"
	"strings"
)

func parse(contents string)  *list.List {
	entries := list.New()
	lines := strings.Split(strings.TrimSpace(contents), "\n")
	for _, line := range lines {
		entries.PushBack(strings.TrimSpace(line))
	}
	return entries
}

func printEntries(entries *list.List) {
	for e := entries.Front(); e != nil; e = e.Next() {
		log.Println(string(e.Value.(string)))
	}
}

func main() {
	ignore, err := os.Open(".gitignore")
	if err != nil {
		log.Fatal("not found.")
	}
	defer ignore.Close()
	bytes := make([]byte, 1024)
	_, e := ignore.Read(bytes)
	if e != nil {
		log.Fatal("read fail.")
	}
	printEntries(parse(string(bytes)))
}
