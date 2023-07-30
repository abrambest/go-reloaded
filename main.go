package main

import (
	"errors"
	"fmt"
	"goreloaded/progs"
	"log"
	"os"
	"time"
)

func start() error {
	strError := "to run use the following format: go run . sample.txt result.txt"
	args := os.Args[1:]
	if len(args) != 2 {
		return errors.New(strError)
	}
	if !(len(args[0]) >= 4 && args[0][len(args[0])-4:] == ".txt" && len(args[1]) >= 4 && args[1][len(args[1])-4:] == ".txt") {
		return errors.New(strError)
	}

	edit(string(args[0]), string(args[1]))
	return nil

}

func edit(fileName, SaveFileName string) {

	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err, time.Now())
		return
	}

	str := progs.Punctuation(string(file))

	str = progs.Atoan(str)
	str = progs.ChangeTo(str)
	str = progs.NumChangeTo(str)
	str = progs.Apostrophe(str)
	str = progs.Punctuation(str)
	str = progs.Apostrophe(str)
	str = progs.FormatN(str)
	saveFile(str, SaveFileName)

}
func saveFile(str, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.WriteString(str)

}

func main() {
	err := start()
	if err != nil {
		log.Fatal(err)
	}
}
