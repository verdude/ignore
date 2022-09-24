package main

import (
	"golang.org/x/exp/slices"
	"log"
	"os"
	"strings"
)

func parse(contents string) []string {
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

func getLines(fname string) []string {
	ignore, err := os.Open(fname)
	if err != nil {
		log.Fatal(fname, " not found.")
	}
	defer ignore.Close()

	b := make([]byte, 1024)
	nread, err := ignore.Read(b)
	if err != nil {
		log.Fatal("read fail.")
	}
	return parse(string(b[:nread]))
}

func isMissing(patterns []string, str string) bool {
	_, found := slices.BinarySearch(patterns, str)
	return !found
}

func main() {
	presets := getLines("presets.txt")
	patterns := getLines(".gitignore")
	missing := make([]string, 0)

	slices.Sort(patterns)

	for _, preset := range presets {
		if isMissing(patterns, preset) {
			missing = append(missing, preset)
		}
	}
	patterns = append(patterns, missing...)
	log.Println(patterns)
}
