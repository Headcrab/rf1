package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	inFile := flag.String("src", "mobydick.txt", "Input file")
	outFile := flag.String("dst", "con", "Output file")
	flag.Parse()

	file, err := os.Open(*inFile)
	check(err)
	defer file.Close()
	reader := bufio.NewReader(file)

	wordCounts := make(map[string]uint64)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		check(err)

		for _, word := range bytes.FieldsFunc(line, func(r rune) bool {
			return !('A' <= r && r <= 'Z' || 'a' <= r && r <= 'z' || '0' <= r && r <= '9')
		}) {
			wordCounts[strings.ToLower(string(word))]++
		}
	}

	type wordCount struct {
		word  string
		count uint64
	}
	var sorted []wordCount
	for word, count := range wordCounts {
		sorted = append(sorted, wordCount{word: word, count: count})
	}
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].count > sorted[j].count })

	outputToConsole := *outFile == "con"
	var writer *bufio.Writer
	if !outputToConsole {
		fileOut, err := os.Create(*outFile)
		check(err)
		defer fileOut.Close()
		writer = bufio.NewWriter(fileOut)
		defer writer.Flush()
	}

	for i := 0; i < len(sorted) && i < 20; i++ {
		if outputToConsole {
			fmt.Printf("%d %q\n", sorted[i].count, sorted[i].word)
		} else if writer != nil {
			fmt.Fprintf(writer, "%d %q\n", sorted[i].count, sorted[i].word)
		}
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
