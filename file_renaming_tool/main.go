package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	var dirPath string
	flag.StringVar(&dirPath, "pathToDir", ".", "full path to a directory")
	flag.Parse()

	dir := filepath.Dir(dirPath)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fileName := file.Name()
		fileType := "File"
		modTime := file.ModTime()
		if file.IsDir() {
			fileType = "Directory"
		}
		fileName, err := stdFileName(fileName)
		if err != nil {
			log.Fatal(err)
		}
		fileName = fileName + "_" + strings.ReplaceAll(modTime.Local().String(), " ", "_")
		fmt.Println(fileType, ":", fileName)
	}
}

func stdFileName(fileName string) (string, error) {
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

	bsOfNewFileName := re.ReplaceAll([]byte(newFileName), []byte(`_`))
	newFileName = string(bsOfNewFileName)
	newFileName = strings.Title(newFileName) + "." + fileExt
	return fmt.Sprintf("%s", newFileName), nil
}
