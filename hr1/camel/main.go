package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var upperCase = regexp.MustCompile(`[A-Z]`)

// Complete the camelcase function below.
func camelcase(s string) int32 {
	if len(strings.Trim(s, " ")) == 0 {
		return 0
	}
	var result int32
	result = 1
	for pos, c := range s {
		if isNewWord(c) && pos != 0 {
			result++
		}
	}
	return result
}

func isNewWord(c int32) bool {
	return upperCase.Match([]byte(fmt.Sprintf("%c", c)))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	s := readLine(reader)
	result := camelcase(s)
	fmt.Println(result)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
