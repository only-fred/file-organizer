package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var fileExtension string
	var fileName string

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileName = file.Name()
		_fileExtension := filepath.Ext(fileName)

		runeFileExtension := []rune(_fileExtension)

		fileExtension = ""

		for i := 1; i < len(runeFileExtension); i++ {
			fileExtension = fileExtension + string(runeFileExtension[i])
		}

		fileExtension = strings.ToUpper(fileExtension)

		createNewDirectory(fileExtension)

		oldLocation := file.Name()
		newLocation := fileExtension + "/" + file.Name()
		isDirectory := file.IsDir()

		if isDirectory == false {
			err := os.Rename(oldLocation, newLocation)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func createNewDirectory(fileExtension string) {
	if _, err := os.Stat(fileExtension); os.IsNotExist(err) {
		os.Mkdir(fileExtension, os.ModePerm)
	}
}
