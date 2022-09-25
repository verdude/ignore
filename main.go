package main

import (
	"flag"
	"fmt"
	"golang.org/x/exp/slices"
	"io"
	"io/ioutil"
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

	b, err := ioutil.ReadFile(fname)
	if err != nil {
		if err != io.EOF {
			log.Fatal("read fail: ", err)
		}
	}
	return parse(string(b))
}

func isMissing(patterns []string, str string) bool {
	_, found := slices.BinarySearch(patterns, str)
	return !found
}

func writeIgnore(patterns []string, stdout bool) {
	var f *os.File
	if stdout {
		f = os.Stdout
	} else {
		f, err := os.OpenFile(".gitignore", os.O_WRONLY, 0444)
		if err != nil {
			log.Fatal("Failed to write to file")
		}
		defer f.Close()
	}

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
	return patterns
}

func main() {
	version := "0.1.4"

	var stdout = flag.Bool("1", false, "Write to stdout instead of to a file")
	var print_version = flag.Bool("v", false, "Print version and exit")
	flag.Parse()

	if *print_version {
		fmt.Println(version)
		return
	}

	patterns := collectPatterns()
	writeIgnore(patterns, *stdout)
}
