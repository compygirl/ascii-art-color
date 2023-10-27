package main

import (
	asciiArt "asciiArt/functions"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) > 6 || len(args) < 1 {
		asciiArt.IncorrectInputError()
		return
	}
	// defining flags from input
	styleFile, bannerFlag, input := asciiArt.Banner(args)
	if !asciiArt.HashChecker(styleFile) { // checking hash
		return
	}
	banner := strings.Split(strings.ReplaceAll(string(styleFile), "\r", ""), "\n") // creating array  of strings of txt file
	colorFlag, input, color, lettersToColor := asciiArt.ColorFlag(input)           // checking color flag
	outputFlag, input, outputtxt := asciiArt.OutputFlag(input)                     // checking output flag
	align := "left"                                                                // default align
	justifyFlag, align, input := asciiArt.JustifyFlag(input)                       // checking align flag
	if justifyFlag == -2 || outputFlag == -2 || colorFlag == -3 {
		return
	}
	flagsum := outputFlag + justifyFlag + bannerFlag + colorFlag // checkibg is input flags correct
	if len(input) == 0 || !asciiArt.IsValid(input[0]) || flagsum+1 != len(args) {
		asciiArt.IncorrectInputError()
		return
	}
	asciinumber := asciiArt.FillMap(make(map[rune]int)) // map for functions
	if len(input[0]) == 0 {
		fmt.Println()
		return
	}
	if !asciiArt.Write(input[0], banner, asciinumber, outputFlag, outputtxt, align, color, lettersToColor) { // running functions function
		return
	}
}
