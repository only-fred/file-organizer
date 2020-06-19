package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var _fileExtension string

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		_fileExtension = filepath.Ext(file.Name())

		_runeFileExtension := []rune(_fileExtension)

		for i := 1; i < len(_runeFileExtension); i++ {
			_runeFileExtension[(i - 1)] = _runeFileExtension[i]

			if i == (len(_runeFileExtension) - 1) {
				_runeFileExtension[i] = ' '
			}
		}

		_fileExtension = string(_runeFileExtension)

		if _, err := os.Stat(_fileExtension); os.IsNotExist(err) {
			os.Mkdir(_fileExtension, os.ModePerm)
		}

		oldLocation := file.Name()
		newLocation := _fileExtension + "/" + file.Name()

		testDirectory := file.IsDir()

		if testDirectory == false && _fileExtension != "go " {
			err := os.Rename(oldLocation, newLocation)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
