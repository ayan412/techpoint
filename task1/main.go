package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"
)

const (
	fileFolder string = "./55_1/"
)

func diff(a, b int) int {
	return a - b
}

func duration(msg string, start time.Time) {
	log.Printf("%v:%v\n", msg, time.Since(start))
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

// reading set of tests and writing output to my_output.txt file
func readInput() {

	defer duration(track("readInput"))

	outFile, err := os.OpenFile("my_output.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("File does not exists or cannot be created")
		os.Exit(1)
	}
	defer outFile.Close()

	var fileFolderFull string
	for i := 1; i <= 25; i++ {
		iStr := strconv.Itoa(i)
		fileFolderFull = fileFolder + iStr

		file, err := os.Open(fileFolderFull)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		//fmt.Println(fileFolderFull)

		reader := bufio.NewReader(file)

		numAstr, _ := reader.ReadString(' ')
		numBstr, _ := reader.ReadString('\n')

		numAstr = strings.TrimSpace(numAstr)
		numBstr = strings.TrimSpace(numBstr)

		numA, _ := strconv.Atoi(numAstr)
		numB, _ := strconv.Atoi(numBstr)
		//fmt.Println(numA, numB)
		res := diff(numA, numB)
		outWr := bufio.NewWriter(outFile)
		fmt.Fprintf(outWr, "%v\n", res)
		outWr.Flush()
	}
}

// reading set of answers and writing output to answers.txt file
func readAnsw() {

	outFileAns, err := os.OpenFile("answers.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("File does not exists or cannot be created")
		os.Exit(1)
	}
	defer outFileAns.Close()

	var fileAnswFull string
	for i := 1; i <= 25; i++ {
		iStr := strconv.Itoa(i) + ".a"
		fileAnswFull = fileFolder + iStr

		file, err := os.Open(fileAnswFull)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		reader := bufio.NewReader(file)

		numAns, _ := reader.ReadString(' ')
		numAns = strings.TrimSpace(numAns)

		outAnsWr := bufio.NewWriter(outFileAns)
		fmt.Fprintf(outAnsWr, "%v\n", numAns)
		outAnsWr.Flush()
	}

}

func main() {
	defer duration(track("main"))
	
	readAnsw()

	f, err := os.Create("memprofile.prof")
	if err != nil {
		fmt.Println("Error while execution:", err)
		return
	}
	defer f.Close()

	pprof.WriteHeapProfile(f)
	readInput()
}
