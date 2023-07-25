package simplefiles

import (
	"io"
	"log"
	"os"
)

const DIR_PATH = "../data"
const FILE_PATH = DIR_PATH + "/simple-file.txt"

func createDir() error {
	if err := os.MkdirAll(DIR_PATH, 0777); err != nil {
		log.Fatalf("Couldnt create Dir %s; %v", DIR_PATH, err)
		return err
	}
	if err := os.Chmod(DIR_PATH, 0777); err != nil {
		log.Fatalf("Couldnt chmod Dir %s; %v", DIR_PATH, err)
		return err
	}
	return nil
}

func createDefaultFile() error {
	file, err := os.Create(FILE_PATH)
	if err != nil {
		log.Fatalf("Couldnt Open/Create file %s, %v", FILE_PATH, err)
		return err
	}
	defer file.Close()

	if l, err := file.WriteString(`Hello world!!`); err != nil {
		log.Fatalf("Couldnt write to file %s, %v+", FILE_PATH, err)
		return err
	} else {
		log.Printf("Wrote %d bytes to file", l)
	}
	return nil
}

func simpleOpenWriteRead() (string, error) {
	if err := createDir(); err != nil {
		return "", err
	}
	err := createDefaultFile()

	file, err := os.Open(FILE_PATH)
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

func simpleFileIOReadAll() (string, error) {
	if err := createDir(); err != nil {
		return "", err
	}

	if err := createDefaultFile(); err != nil {
		return "", err
	}
	f, err := os.Open(FILE_PATH)
	if err != nil {
		log.Fatalf("Failed to open file %v; %+v", FILE_PATH, err)
		return "", err
	}
	bContent, err := io.ReadAll(f)
	if err != nil {
		if err != io.EOF {
			log.Fatalf("Failed to read file %v into buffer; %+v", FILE_PATH, err)
			return "", err
		}
	}
	return string(bContent), nil
}
