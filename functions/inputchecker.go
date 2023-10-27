package asciiArt

import (
	"crypto/md5"
	"fmt"
	"log"
)

func Hash(data string) string {
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}

func HashChecker(file []byte) bool {
	stdHash := "ac85e83127e49ec42487f272d9b9db8b"
	shdHash := "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	thkHash := "db448376863a4b9a6639546de113fa6f"
	if Hash(string(file)) != stdHash && Hash(string(file)) != shdHash && Hash(string(file)) != thkHash {
		log.Fatal("Error: banner file was changed")
		return false
	}
	return true
}

func IsValid(str string) bool {
	for i := 0; i < len(str); i++ {
		if (str[i] < 32 || str[i] > 126) && str[i] != 10 {
			fmt.Println("Error:This string contains symbols that can not be graphically represented")
			return false
		}
	}
	return true
}
