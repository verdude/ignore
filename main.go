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

func getLines(fname string, create bool) []string {
	f, err := os.Open(fname)
	if err != nil {
		if create {
			f, err = os.Create(fname)
			if err != nil {
				log.Println("Could not create", fname)
				log.Fatal("ERROR: ", err)
			} else {
				log.Println("Created", fname)
				f.Close()
				return make([]string, 0)
			}
		}
		log.Fatal(fname, " not found.")
	}
	defer f.Close()

	b := make([]byte, 1024)
	nread, err := f.Read(b)
	if err != nil {
		log.Fatal("read fail.")
	}
	return parse(string(b[:nread]))
}

func isMissing(patterns []string, str string) bool {
	_, found := slices.BinarySearch(patterns, str)
	return !found
}

func writeIgnore(patterns []string) {
	f, err := os.Open(".gitignore")
	if err != nil {
		log.Fatal("Failed to write to file")
	}
	defer f.Close()

	for _, pattern := range patterns {
		f.Write([]byte(pattern))
		f.Write([]byte("\n"))
	}
}

func collectPatterns() []string {
	presets := getLines("/etc/ignoregit/presets.txt", true)
	patterns := getLines(".gitignore", true)
	missing := make([]string, 0)

	slices.Sort(patterns)

	for _, preset := range presets {
		if isMissing(patterns, preset) {
			missing = append(missing, preset)
		}
	}
	patterns = append(patterns, missing...)
	log.Println(patterns)
	return patterns
}

func main() {
	patterns := collectPatterns()
	writeIgnore(patterns)
}
