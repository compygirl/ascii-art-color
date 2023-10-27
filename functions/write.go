package asciiArt

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func removeEmpty(array []string) []string {
	for i := 0; i < len(array); i++ {
		if array[i] == "" {
			copy(array[i:], array[i+1:])
			array = array[:len(array)-1]
			i--
		}
	}
	return array
}

func displaySize() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	size_str := out[3 : len(out)-1]
	size_num, err := strconv.Atoi(string(size_str))
	if err != nil {
		log.Fatal(err)
		return -1
	}
	return size_num
}

func isColorSymbol(colorArray []rune, symbol rune) bool {
	if colorArray == nil {
		return false
	}
	for i := 0; i < len(colorArray); i++ {
		if symbol == colorArray[i] {
			return true
		}
	}
	return false
}

func alignAndColorLetter(str string, lettersToColor []rune, banner []string, asciiMap map[rune]int, color, align string, size_num int) string {
	allColorCmdLen := 0
	colorCmdLen := len(color) + len("\033[0m")
	asciiartlen := 0
	spaceRangeNum := 1
	spaceNum := 0
	strSplit := []string{str}
	if align == "justify" {
		strSplit = strings.Split(str, " ") // SPLITING LINE TO WORDS FOR JUSTIFY ALIGN
	}
	strSplit = removeEmpty(strSplit) // REMOVEING NOT TRASH EMPTYIES AFTER SPLIT
	if len(strSplit) != 1 {
		spaceRangeNum = len(strSplit) - 1
	}
	asciiArtWord := make([][]string, 0)
	for i := 0; i < len(strSplit); i++ {
		asciiArtWord = append(asciiArtWord, make([]string, 8)) // CREATING 2DIM ARRAY OF STRINGS WHERE EVERY ARRAY OF FIRST LEVEL CONTAIN 8 STRINGS OF AN WORD/LINE(DEPENDS FROM FLAG)
	}

	for strWordInd := 0; strWordInd < len(strSplit); strWordInd++ { // LOOP FOR EVERY WORD/LINE
		for i := 0; i < len(strSplit[strWordInd]); i++ { //
			asciiArtSymbol := make([]string, 8)
			for j := 0; j < 8; j++ { // writing SYMBOL
				asciiArtSymbol[j] += banner[asciiMap[rune(strSplit[strWordInd][i])]+j]
			}
			if isColorSymbol(lettersToColor, rune(strSplit[strWordInd][i])) || (lettersToColor == nil && color != "") { // color letter if it is needed
				allColorCmdLen += colorCmdLen
				for j := 0; j < 8; j++ {
					asciiArtSymbol[j] = color + asciiArtSymbol[j] + "\033[0m"
				}
			}
			for j := 0; j < 8; j++ { // ADDING SYMBOL TO ITS LINE/WORD
				asciiArtWord[strWordInd][j] += asciiArtSymbol[j]
			}
		}
		asciiartlen += len(asciiArtWord[strWordInd][0]) // COUNTINF LEN OF ASCIIART FOR RESIZING
	}
	temp := ""
	if align == "center" { // adding space for align
		spaceNum = (size_num - asciiartlen + allColorCmdLen) / 2
		for i := 0; i < spaceNum; i++ {
			temp += " "
		}
		for i := 0; i < len(asciiArtWord); i++ {
			for j := 0; j < 8; j++ {
				asciiArtWord[i][j] = temp + asciiArtWord[i][j] + temp
			}
		}
	} else if align == "right" {
		spaceNum = (size_num - asciiartlen + allColorCmdLen)
		for i := 0; i < spaceNum; i++ {
			temp += " "
		}
		for i := 0; i < len(asciiArtWord); i++ {
			for j := 0; j < 8; j++ {
				asciiArtWord[i][j] = temp + asciiArtWord[i][j]
			}
		}
	} else if align == "justify" {
		spaceNum = (size_num - asciiartlen + allColorCmdLen) / spaceRangeNum
		for i := 0; i < spaceNum; i++ {
			temp += " "
		}
		for i := 0; i < len(asciiArtWord)-1; i++ {
			for j := 0; j < 8; j++ {
				asciiArtWord[i][j] = asciiArtWord[i][j] + temp
			}
		}

	}
	resutString := ""
	for j := 0; j < 8; j++ { // rewriting an array to return string
		for i := 0; i < len(asciiArtWord); i++ {
			resutString += asciiArtWord[i][j]
		}
		if j != 7 {
			resutString += "\n"
		}
	}
	return resutString
}

func Write(sentence string, banner []string, asciiMap map[rune]int, outputFlagIndex int, outputtxt string, align string, color string, lettersToColor []rune) bool {
	size_num := displaySize() // terminal size
	if size_num == -1 {
		fmt.Println("Error: Impossible to get display size")
		return false
	}
	splitedString := Split(sentence) // split by newlines
	resultString := make([]string, 0)
	for i := 0; i < len(splitedString); i++ {
		if splitedString[i] == "\n" || splitedString[i] == "" {
			resultString = append(resultString, "\n")
		} else {
			resultString = append(resultString, alignAndColorLetter(splitedString[i], lettersToColor, banner, asciiMap, color, align, size_num))
		}
	}
	if outputFlagIndex > 0 { // write to a file if output flag is true
		write_file, err := os.Create(outputtxt)
		if err != nil {
			fmt.Println("Error: cannot create writing  file")
			return false
		}
		writer := bufio.NewWriter(write_file)
		for i := 0; i < len(resultString); i++ {
			_, err := writer.WriteString(resultString[i])
			if err != nil {
				fmt.Println("Error: cannot write to file")
				return false
			}
		}
		defer write_file.Close()
		writer.Flush()
	} else { // print to terminal
		for i := 0; i < len(resultString); i++ {
			fmt.Print(resultString[i])
		}
	}
	return true
}
