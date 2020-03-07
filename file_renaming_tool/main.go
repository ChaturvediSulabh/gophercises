package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fileName := "ngi.nx .Log\\/?"
	stdFileName(fileName)
}

func stdFileName(fileName string) {
	newFileName := strings.ToLower(fileName)
	slOfNewFileName := strings.Split(newFileName, ".")
	fileExt := slOfNewFileName[len(slOfNewFileName)-1]
	newFileName = strings.Join(slOfNewFileName[0:len(slOfNewFileName)-1], ``)

	re := regexp.MustCompile(`\W`)
	bsOfFileExt := re.ReplaceAll([]byte(fileExt), []byte(``))
	fileExt = string(bsOfFileExt)
	if fileExt == "" || len(slOfNewFileName) == 1 {
		fileExt = "txt"
	}

	bsOfNewFileName := re.ReplaceAll([]byte(newFileName), []byte(``))
	newFileName = string(bsOfNewFileName)
	newFileName = strings.Title(newFileName) + "." + fileExt
	fmt.Println(newFileName)
}
