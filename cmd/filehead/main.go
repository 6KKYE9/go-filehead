package main

import (
	"fmt"
	"os"
	"strconv"

	"filehead"
)

func main() {
	mode := "lines"
	n := 10
	args := os.Args[1:]
	var files []string
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-bytes":
			mode = "bytes"
		case "-n":
			if i+1 < len(args) {
				if v, err := strconv.Atoi(args[i+1]); err == nil {
					n = v
				}
				i++
			}
		default:
			files = append(files, args[i])
		}
	}
	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "usage: filehead [-bytes] [-n N] <file>...")
		os.Exit(1)
	}
	for _, f := range files {
		if mode == "bytes" {
			b, err := filehead.HeadBytes(f, n)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %v\n", f, err)
				continue
			}
			os.Stdout.Write(b)
			fmt.Println()
		} else {
			lines, err := filehead.HeadLines(f, n)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %v\n", f, err)
				continue
			}
			fmt.Print(filehead.Preview(lines))
		}
	}
}
