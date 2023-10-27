package asciiArt

import "fmt"

func TooLongWordError(str string, ind, index int) {
	fmt.Println("'", str[ind:index], "'", "this word is to long to write in one line")
}

func IncorrectInputError() {
	// fmt.Println("Error: Incorrect useage")
	// fmt.Println("Possible Usage: go run . [STRING] [BANNER]	")
	fmt.Println("EX: go run . something standard	")
	fmt.Println()
	fmt.Println("Possible Usage: go run . [OPTION] [STRING] [BANNER]	")
	// fmt.Println("EX: go run .--align=center/--output=test.txt something standard	")
	// fmt.Println()
	// fmt.Println("Possible Usage: go run . [OPTION1] [OPTION2] [STRING] [BANNER]	")
	// fmt.Println("EX: go run . --align=center --output=test.txt something standard	")
}
