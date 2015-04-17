package common

import (
	"log"
	"os"
)

func CreateDirIfNotExist(path string) {
	if b, _ := exists(path); b != true {
		os.Mkdir(path, 0777)
	} else {
		log.Println("dir exists")
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
