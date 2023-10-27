package asciiArt

func Split(file string) []string {
	splitedstring := make([]string, 0)
	str := ""
	for i := 0; i < len(file); i++ {
		if (i+1 < len(file) && file[i] == '\\' && file[i+1] == 'n') || file[i] == 10 {
			if str != "" {
				splitedstring = append(splitedstring, str)
				str = ""
			}
			splitedstring = append(splitedstring, "\n")
			if file[i] != 10 {
				i++
			}
		} else {
			str += string(file[i])
		}
	}
	if str != "" {
		splitedstring = append(splitedstring, str)
	}
	// fmt.Println(splitedstring)
	return splitedstring
}
