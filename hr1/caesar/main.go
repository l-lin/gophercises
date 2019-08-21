package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	alphabet    = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	alphabetMap map[string]int
	upperCase   = regexp.MustCompile(`[A-Z]`)
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("Please provide the string to encrypt and the number of letters to rotate")
	}
	s := os.Args[1]
	k, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(caesarCipher(s, k))
}

func caesarCipher(s string, k int) string {
	alphabetMap = make(map[string]int, len(alphabet))
	for pos, a := range alphabet {
		alphabetMap[a] = pos
	}
	var text strings.Builder
	for _, c := range s {
		up := isUpper(c)
		var cipheredC string
		if up {
			cipheredC = strings.ToLower(string(c))
		} else {
			cipheredC = string(c)
		}
		if pos, ok := alphabetMap[cipheredC]; ok {
			newPos := pos + k
			if newPos >= len(alphabet) {
				newPos = newPos - len(alphabet)
			}
			cipheredC = alphabet[newPos]
		}
		if up {
			cipheredC = strings.ToUpper(cipheredC)
		}
		//fmt.Printf("%s -> %s\n", string(c), cipheredC)
		text.WriteString(cipheredC)
	}

	return text.String()
}

func isUpper(c int32) bool {
	return upperCase.Match([]byte(string(c)))
}
