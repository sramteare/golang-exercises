package simplefiles

import (
	"io"
	"log"
	"os"
)

const DIR_PATH = "../data"
const FILE_PATH = DIR_PATH + "/simple-file.txt"

func simpleOpenWriteRead() (string, error) {
	if err := os.MkdirAll(DIR_PATH, 0777); err != nil {
		log.Fatalf("Couldnt create Dir %s; %v", DIR_PATH, err)
		return "", err
	}
	if err := os.Chmod(DIR_PATH, 0777); err != nil {
		log.Fatalf("Couldnt chmod Dir %s; %v", DIR_PATH, err)
		return "", err
	}
	file, err := os.Create(FILE_PATH)
	if err != nil {
		log.Fatalf("Couldnt Open/Create file %s, %v", FILE_PATH, err)
		return "", err
	}
	defer file.Close()

	if l, err := file.WriteString(`Hello world!!`); err != nil {
		log.Fatalf("Couldnt write to file %s, %v+", FILE_PATH, err)
		return "", err
	} else {
		log.Printf("Wrote %d bytes to file", l)
	}
	file.Seek(0, 0)
	byts := []byte{}
	buff := make([]byte, 3)
	var n int
	for err == nil {
		n, err = file.Read(buff)
		if err != nil && err != io.EOF {
			log.Fatalf("Failed to read file; %+v", err)
			return "", err
		}
		//log.Println(n)
		byts = append(byts, buff[:n]...)
	}

	return string(byts), nil
}
