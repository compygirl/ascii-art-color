package asciiArt

func FillMap(str map[rune]int) map[rune]int {
	arr := []rune{}
	arr2 := []int{}

	for i := 32; i <= 126; i++ {
		symbol := (i-32)*9 + 1
		arr = append(arr, rune(i))
		arr2 = append(arr2, symbol)
	}
	for i := 0; i < len(arr); i++ {
		str[arr[i]] = arr2[i]
	}
	return str
}
