package asciiArt

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func removeString(wordsList []string, i int) []string {
	copy(wordsList[i:], wordsList[i+1:])
	return wordsList[:len(wordsList)-1]
}

func removeRune(wordsList []rune, i int) []rune {
	copy(wordsList[i:], wordsList[i+1:])
	return wordsList[:len(wordsList)-1]
}

func OutputFlag(args []string) (int, []string, string) {
	outputtxt := ""
	for i := 0; i < len(args); i++ {
		if i > 1 { // if index more than 1 it;s mean that we don't have flag but have slot for it. It is an error
			return -1, args, ""
		}
		if len(args[i]) > 9 && args[i][:9] == "--output=" {
			if args[i][len(args[i])-4:] != ".txt" {
				fmt.Println("Error: file should be in txt format")
				return -2, args, ""
			} else {
				outputtxt = args[i][9:]
				args := removeString(args, i)
				return 1, args, outputtxt
			}
		}
	}
	return 0, args, ""
}

func JustifyFlag(args []string) (int, string, []string) {
	for i := 0; i < len(args); i++ {
		if i > 1 { // if index more than 2(because we deleted outputflag and colorflag if they were here) its mean that we don't have flag but have slot for it. It is an error
			return -1, "", args
		}
		if len(args[i]) > 8 && args[i][:8] == "--align=" {
			if args[i][8:] == "center" || args[i][8:] == "left" || args[i][8:] == "right" || args[i][8:] == "justify" {
				align := args[i][8:]
				args = removeString(args, i)
				return 1, align, args
			} else {
				fmt.Println("Error: Not correct align format") // correct la g but incorrect file
				return -2, "", args
			}
		}
	}
	return 0, "", args
}

func Banner(args []string) ([]byte, int, []string) {
	file, err := ioutil.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("error: can not read file")
	}
	// we need this condition because we should find out is last args banner or not
	if args[len(args)-1] == "shadow" {
		args = removeString(args, len(args)-1)
		file, err := ioutil.ReadFile("shadow.txt")
		if err != nil {
			fmt.Println("error: can not read file")
		}
		return file, 1, args
	} else if args[len(args)-1] == "thinkertoy" {
		args = removeString(args, len(args)-1)
		file, err := ioutil.ReadFile("thinkertoy.txt")
		if err != nil {
			fmt.Println("error: can not read file")
		}
		return file, 1, args
	} else if args[len(args)-1] == "standard" {
		args = removeString(args, len(args)-1)
		return file, 1, args
	}
	return file, 0, args
}

func ColorFlag(args []string) (int, []string, string, []rune) {
	reg := regexp.MustCompile("^(black|red|green|yellow|blue|magenta|cyan|white)$")
	colorMap := map[string]string{
		"black":   "\033[30m",
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
	}
	colorcode := ""
	for i := 0; i < len(args); i++ {
		if i > 2 { // if index more than 1 it;s mean that we don't have flag but have slot for it. It is an error
			return -1, args, colorcode, nil
		}
		if len(args[i]) > 8 && args[i][:8] == "--color=" && reg.MatchString(args[i][8:]) {
			colorcode = colorMap[args[i][8:]]
			if (i+2 > len(args)-1) || (len(args[i+1]) > 9 && args[i+1][:9] == "--output=") || (len(args[i+1]) > 8 && args[i+1][:8] == "--align=") {
				args = removeString(args, i)
				return 1, args, colorcode, nil
			} else {
				if IsValid(args[i+1]) {
					lettersToColor := []rune(args[i+1])
					args = removeString(args, i)
					args = removeString(args, i)
					return 2, args, colorcode, lettersToColor
				} else {
					fmt.Println("Error:not allowed symbol to color")
					return -2, args, colorcode, nil
				}
			}
		} else if len(args[i]) > 8 && args[i][:8] == "--color=" && !reg.MatchString(args[i][8:]) {
			fmt.Println("Error: Not alloweded color")
			return -3, args, colorcode, nil
		}
	}
	return 0, args, colorcode, nil
}
