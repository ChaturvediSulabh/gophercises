package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	var dirPath string
	flag.StringVar(&dirPath, "pathToDir", ".", "full path to a directory")
	flag.Parse()

	dir := filepath.Dir(dirPath)
	err := filepath.Walk(dir,
		func(dir string, info os.FileInfo, err error) error {
			if err != nil {
				log.Fatal(err)
			}

			fileName := info.Name()
			fileType := "Directory"
			if !info.IsDir() {
				fileType = "File"
				fileName, err = stdFileName(fileName)
				if err != nil {
					log.Fatal(err)
				}
				fileName = fileName + "_" + strings.ReplaceAll(info.ModTime().Local().String(), " ", "_")
			}
			fmt.Println(fileType + ": " + fileName)
			return nil
		})
	if err != nil {
		log.Fatal(err)
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
		newFileName = slOfNewFileName[0]
	}

	bsOfNewFileName := re.ReplaceAll([]byte(newFileName), []byte(`_`))
	newFileName = string(bsOfNewFileName)
	newFileName = strings.Title(newFileName) + "." + fileExt
	return fmt.Sprintf("%s", newFileName), nil
}
