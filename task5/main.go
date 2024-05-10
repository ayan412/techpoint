package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
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
func writeToOutputFile(iter int) *os.File {
	// Формирование пути к выходному файлу
	pathToFile := fmt.Sprintf("%s%d.txt", outputFilePath, iter)

	// Открытие или создание файла
	outFile, err := os.OpenFile(pathToFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("failed to open or create file: %w", err)
		return nil
	}
	return outFile
}

type FoldersPath struct {
	Dir     string        `json:"-"`
	Files   []string      `json:"files"`
	Folders []FoldersPath `json:"folders,omitempty"`
}

func main() {

	numSets, rdr, _ := readSets()

	for i := 1; i <= numSets; i++ {
		numRows := readRows(rdr)
		readJson(numRows, rdr)

	}

	// for i := 2; i <= numSets; i++ {
	// 	numRows := readRows(rdr)
	// 	readJson(numRows, rdr)
	// }

}

// Ф-я для чтения первого числа - кол-ва набора данных
func readSets() (int, *bufio.Reader, error) {

	//defer duration(track("readiInput"))

	var numSets int

	var rdr *bufio.Reader

	// Цикл для обработки каждого файла в каталоге "71_5"
	for i := 1; i <= 1; i++ {
		iStr := strconv.Itoa(i)
		// Путь до конкретного файла с наборами данных
		filePathFull := fmt.Sprintf("%s%s", filePath, iStr)

		// Открытие файла
		file, err := os.Open(filePathFull)
		if err != nil {
			return 0, nil, fmt.Errorf("error opening the file: %w", err)
		}
		//defer file.Close() // Закрытие файла по окончании работы с ним

		// Создание Reader для чтения файла
		rdr = bufio.NewReader(file)

		// Чтение количества наборов входных данных в файле
		numOfSetsStr, err := rdr.ReadString('\n')
		if err != nil {
			return 0, nil, fmt.Errorf("Error reading amount of sets: %w", err)
		}

		// Преобразование количества наборов из строки в число
		numOfSetsStr = strings.TrimSuffix(numOfSetsStr, "\n")
		numOfSetsInt, err := strconv.Atoi(numOfSetsStr)
		if err != nil {
			return 0, nil, fmt.Errorf("failed to convert sets count to int: %w", err)
		}

		// Создание Writer для записи в соответ-й выходной файл
		outWrite := bufio.NewWriter(writeToOutputFile(i))
		defer outWrite.Flush()

		// количество наборов входных данных
		numSets = numOfSetsInt
		// кол-во строк
		// numRows = readRows(rdr)
	}
	return numSets, rdr, nil
}

var ErrNumberInStream = errors.New("number found in JSON stream")

func readJson(numRows int, rdr *bufio.Reader) {
	// Иниц-ия структуры чтобы избежать пустого указ-ля
	var r FoldersPath

	// Чтение строк где содержится json

	var jsonBuilder strings.Builder
	for j := 1; j <= numRows; j++ {
		line, err := rdr.ReadString('\n')
		if err != nil {
			fmt.Println("ошибка при чтении строки JSON:", err)
		}
		jsonBuilder.WriteString(line)
	}

	if err := json.Unmarshal([]byte(jsonBuilder.String()), &r); err != nil {
		fmt.Println("ошибка при разборе JSON: %w", err)
	}
	fmt.Println("json object:", r)
	files := findHackFiles(r, ".hack", ".exe")
	fmt.Println(len(files))
}

// Ф-я для чтения кол-ва строк с описанием директорий
func readRows(rdr *bufio.Reader) int {
	strOfRow, err := rdr.ReadString('\n')
	if err != nil {
		fmt.Println("Error in reading row with sets of rows", err)
	}
	strOfRows := strings.TrimSuffix(strOfRow, "\n")
	numOfRows, err := strconv.Atoi(strOfRows)
	if err != nil {
		fmt.Println("Error in type converting", err)
	}
	return numOfRows
}

func findHackFiles(fp FoldersPath, extensions ...string) []string {
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

// Нужно считать оставшиеся строки с json - взять из task4:263
