package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {

	in_file := flag.String("src", "mobydick.txt", "Input file")
	out_file := flag.String("dst", "con", "Output file")
	flag.Parse()
	fil, err := os.Open(*in_file)
	//// or just
	// fil, err := os.Open("mobydick.txt")
	//// if i can`t use string
	check(err)
	defer fil.Close()
	rd := bufio.NewReader(fil)

	var out MyMap
	for {
		var str []byte
		str, err := rd.ReadSlice('\n')
		str_div := bytes.FieldsFunc(str, DivByWords)
		if len(str_div) > 0 {
			for _, wrd := range str_div {
				ToLower(&wrd)
				i := out.Find(&wrd)
				if i == -1 {
					out = append(out, MyMapItem{1, append([]byte{}, wrd...)})
				} else {
					out[i].count++
				}
			}
		}
		if err == io.EOF {
			break
		}
		check(err)
	}

	sort.Sort(out)

	var wr *bufio.Writer
	if *out_file != "con" {
		fil_out, err := os.Create(*out_file)
		check(err)
		defer fil_out.Close()
		wr = bufio.NewWriter(fil_out)
		defer wr.Flush()
	}
	for i, it := range out {
		if i > 19 {
			break
		}
		if wr == nil {
			fmt.Printf("%d %q\n", it.count, it.text)
		} else {
			fmt.Fprintf(wr, "%d %q\n", it.count, it.text)
		}
	}
}

type MyMapItem struct {
	count uint64
	text  []byte
}

type MyMap []MyMapItem

func (s MyMap) Len() int {
	return len(s)
}

func (s MyMap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s MyMap) Less(i, j int) bool {
	return s[i].count > s[j].count
}

func (slice MyMap) Find(value *[]byte) int {
	for p, v := range slice {
		if bytes.Equal(v.text, *value) {
			return p
		}
	}
	return -1
}

func ToLower(wrd *[]byte) {
	for i, ltr := range *wrd {
		if (ltr >= 65) && (ltr <= 90) {
			(*wrd)[i] += 32
		}
	}
}

func DivByWords(r rune) bool {
	return r == ' ' || r == '.' || r == ',' || r == ':' || r == ';' || r == '!' ||
		r == '?' || r == '*' || r == '\n' || r == '(' || r == ')' || r == '-' ||
		r == '"' || r == '\r' || r == '\\' || r == '[' || r == ']' || r == '@' || r == '/' ||
		r == '\t' || r == '{' || r == '}' || r == '\''
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
