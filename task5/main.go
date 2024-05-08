package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	filePath       string = "./71_5/"   // Папка, содержащая входные файлы для теста
	outputFilePath string = "./output/" // Папка в которую будут записываться выходные файлы
)

// Функция для измерения времени выполнения
func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

// Функция для отслеживания времени выполнения функции
func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

// Функция для создания выходного файла
func writeOutFile(iter int) *os.File {
	// Формирование пути к выходному файлу
	pathToFile := fmt.Sprintf("%s%d.txt", outputFilePath, iter)

	// Открытие или создание файла
	outFile, err := os.OpenFile(pathToFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("File does not exist or cannot be created:", err)
		os.Exit(1)
	}
	return outFile
}

// Ф-я для чтения первого числа - кол-ва набора данных
func getSets() (int, int) {

	//defer duration(track("readiInput"))

	var numSets, numRows int
	var rdr *bufio.Reader
	// Цикл для обработки каждого файла в каталоге "71_5"
	for i := 1; i <= 1; i++ {
		iStr := strconv.Itoa(i)
		// Путь до конкретного файла с наборами данных
		filePathFull := fmt.Sprintf("%s%s", filePath, iStr)

		// Открытие файла
		f, err := os.Open(filePathFull)
		if err != nil {
			fmt.Println("Error while opening the file", err)
			break
		}
		defer f.Close() // Закрытие файла по окончании работы с ним

		// Создание Reader для чтения файла
		
		rdr = bufio.NewReader(f)

		// Чтение количества наборов входных данных в файле
		numOfSetsStr, err := rdr.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading amount of sets:", err)
			break
		}

		//Преобразование количества наборов из строки в число
		numOfSetsStrTrim := strings.TrimSuffix(numOfSetsStr, "\n")
		numOfSetsInt, err := strconv.Atoi(numOfSetsStrTrim)
		if err != nil {
			fmt.Println("Error in casting of set if input Data", err)
		}

		//fmt.Println("amount of sets:", numOfSetsInt)

		// Создание Writer для записи в соответ-й выходной файл
		outWrite := bufio.NewWriter(writeOutFile(i))
		defer outWrite.Flush()

		// количество наборов входных данных
		numSets = numOfSetsInt
		// кол-во строк
		numRows = readRows(rdr)

	}
	readJson(numRows, rdr)
	return numSets, numRows
}

func readJson(numRows int, rdr *bufio.Reader) {

	var r FolderPath

	for j := 1; j <= numRows; j++ {
		dec := json.NewDecoder(rdr)
		if err := dec.Decode(&r); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s, %s\n", r.Files, r.Folders)

		// Находим все файлы с расширениями .hack или .exe
		files := findHackFiles(r, ".hack", "exe")
		fmt.Println(len(files))
	}

}

// Ф-я для чтения кол-ва строк с описанием директорий
func readRows(rdr *bufio.Reader) int {
	strOfRow, err := rdr.ReadString('\n')
	if err != nil {
		fmt.Println("Error in reading row with sets of rows", err)
	}
	strOfRowsTrim := strings.TrimSuffix(strOfRow, "\n")
	strOfRows, err := strconv.Atoi(strOfRowsTrim)

	if err != nil {
		fmt.Println("Error in type converting", err)
	}
	return strOfRows
}

type FolderPath struct {
	Dir     string       `json:"-"`
	Files   []string     `json:"files"`
	Folders []FolderPath `json:"folders,omitempty"`
}

func findHackFiles(fp FolderPath, extensions ...string) []string {
	var matchingFiles []string

	// Проверяем все файлы в Files
	for _, file := range fp.Files {
		for _, ext := range extensions {
			if strings.HasSuffix(file, ext) {
				matchingFiles = append(matchingFiles, file)
				break
			}
		}
	}

	// Проверяем все файлы в Folders
	for _, folder := range fp.Folders {
		matchingFiles = append(matchingFiles, findHackFiles(folder, extensions...)...)
	}
	return matchingFiles
}

func main() {

	_, numRows := getSets()

	fmt.Println(numRows)

}

// Нужно считать оставшиеся строки с json - взять из task4:263
