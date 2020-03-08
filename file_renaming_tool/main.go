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
			if info.IsDir() {
				fmt.Printf("%s: %s\n", fileType, fileName)
			} else {
				fileType = "File"
				fileName, err = stdFileName(fileName)
				if err != nil {
					log.Fatal(err)
				}
				fileName = fileName + "_" + strings.ReplaceAll(info.ModTime().Local().String(), " ", "_")
				fmt.Printf("\t%s: %s\n", fileType, fileName)
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
}

func stdFileName(fileName string) (string, error) {
	newFileName := strings.ToLower(fileName)
	fileExt := filepath.Ext(newFileName)
	if fileExt == "" {
		fileExt = ".txt"
	}
	re := regexp.MustCompile(`\W`)
	bsOfNewFileName := re.ReplaceAll([]byte(newFileName), []byte(`_`))
	newFileName = string(bsOfNewFileName)
	newFileName = strings.Title(newFileName) + fileExt
	return fmt.Sprintf("%s", newFileName), nil
}
