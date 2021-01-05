package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func xor(input, key []byte) []byte {
	ret := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		ret[i] = input[i] ^ key[i%len(key)]
	}
	return ret
}

func overWrite(file string) {
	master := []byte{82, 51, 100, 95, 72, 51, 114, 114, 49, 110, 103}
	encKey := []byte{0, 122, 48, 12, 13, 112, 9, 59, 92, 30, 11, 18, 93, 16, 106, 23, 115, 0, 65, 110, 8, 18, 60, 108, 39, 23, 120, 3, 63, 15}
	unEncKey := xor(encKey, master)
	fileContents, _ := ioutil.ReadFile(file)
	temp := xor(fileContents, unEncKey)
	ioutil.WriteFile(file, temp, 0777)
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	pathSlice := strings.Split(path, "\\")
	if pathSlice[len(pathSlice)-1] != "Shifuku Halo 2020" {
		os.Exit(2)
	}
	if pathSlice[len(pathSlice)-2] != "Implants" {
		os.Exit(2)
	}
	if pathSlice[len(pathSlice)-3] != "Core" {
		os.Exit(2)
	}

	newPath := strings.Join(pathSlice[:(len(pathSlice)-2)], "\\")

	var files []string = make([]string, 0)
	temp := filepath.Walk(newPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Println(temp)
	}
	for _, v := range files {
		overWrite(v)
	}
}
