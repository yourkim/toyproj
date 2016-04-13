package main

import (
	"os"
	"fmt"
	"strconv"
	"path/filepath"
	"io/ioutil"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("Usage: %s <text file name> <current line width> <desired line width>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	inputFileName := os.Args[1]
	oldWidth, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("<current line width> shoulud be a number.")
		os.Exit(1)
	}
	newWidth, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("<desired line width> shoulud be a number.")
		os.Exit(1)
	}


	fmt.Printf("Convert %q from line width (%d) to line width (%d)\n", inputFileName, oldWidth, newWidth)


	buf, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		fmt.Printf("Error in opening file %q: %s", inputFileName, err.Error())
		os.Exit(1)
	}

	// convert buf into a slice of string, so that multi-byte character is correctly counted as a single character.
	contents := []rune(string(buf))
	newContents := []rune{}

	lineWidth := 0
	for i := 0; i < len(contents); i++ {
		fmt.Printf("Char at [%d]: %c, lineWidth: %d\n", i, contents[i], lineWidth)
		if lineWidth == oldWidth && contents[i] == '\n' {
		} else {
			newContents = append(newContents, contents[i])
		}
		lineWidth++
		if lineWidth > oldWidth || contents[i] == '\n' {
			lineWidth = 0
		}
	}

	fmt.Print(string(newContents))
}
