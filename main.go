// shape-of-code - convert source code to images
// Copyright (c) 2023 Luka Ivanovic
// This code is licensed under MIT licence (see LICENCE for details)

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/fogleman/gg"
)

const (
	VERSION = "0.2.0"
)

const (
	primary   = "000000FF"
	secondary = "FFFFFFFF"
	tabsize   = 4
	blocksize = 10
)

type Line struct {
	indent int
	length int
}

type File struct {
	name  string
	lines []Line
	max   int
}

func printUsage(fail bool) {
	fmt.Println(`shape-of-code
	
	shape-of-code input.txt
	shape-of-code input1.txt input2.txt
	shape-of-code help (or --help or -h)
	shape-of-code version (or --version or -v)`)
	if fail {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func printVersion() {
	fmt.Println("shape-of-code v" + VERSION)
	os.Exit(0)
}

func main() {
	args := os.Args[1:] // skip program name
	if len(args) == 0 {
		printUsage(true)
	} else if args[0] == "help" || args[0] == "--help" || args[0] == "-h" {
		printUsage(false)
	} else if args[0] == "version" || args[0] == "--version" || args[0] == "-v" {
		printVersion()
	}
	for _, arg := range args {
		file, err := os.Open(arg)
		if err != nil {
			log.Fatalln(err)
		}
		scanner := bufio.NewScanner(file)
		t := File{arg, []Line{}, 0}
		for scanner.Scan() {
			line := scanner.Bytes()
			length := len(line)
			indent := 0
			for i := range line {
				if line[i] == ' ' {
					indent += 1
				} else if line[i] == '\t' {
					indent += tabsize
					length += tabsize - 1
				} else {
					break
				}
			}
			t.lines = append(t.lines, Line{
				indent: indent,
				length: length - indent,
			})
			if length > t.max {
				t.max = length
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatalln(err)
		}
		width := t.max*blocksize + 2*blocksize
		height := len(t.lines)*blocksize + 2*blocksize
		img := gg.NewContext(width, height)
		img.DrawRectangle(0.0, 0.0, float64(width), float64(height))
		img.SetHexColor(secondary)
		img.Fill()
		for i := range t.lines {
			x := float64(t.lines[i].indent)*blocksize + blocksize
			y := float64(i)*blocksize + blocksize
			w := float64(t.lines[i].length) * blocksize
			h := float64(blocksize)
			img.DrawRectangle(x, y, w, h)
			img.SetHexColor(primary)
			img.Fill()
		}
		img.SavePNG(t.name + ".png")
	}
}
