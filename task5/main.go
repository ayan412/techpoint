package main

import (
	"bufio"
	"encoding/json"
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
	Dir     string        `json:"dir"`
	Files   []string      `json:"files"`
	Folders []FoldersPath `json:"folders"`
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
	for i := 4; i <= 4; i++ {
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

func readJson(numRows int, rdr *bufio.Reader) {
	// Иниц-ия структуры чтобы избежать пустого указ-ля
	var root FoldersPath

	// Чтение строк где содержится json

	var jsonBuilder strings.Builder
	for j := 1; j <= numRows; j++ {
		line, err := rdr.ReadString('\n')
		if err != nil {
			fmt.Println("ошибка при чтении строки JSON:", err)
		}
		jsonBuilder.WriteString(line)
	}

	if err := json.Unmarshal([]byte(jsonBuilder.String()), &root); err != nil {
		fmt.Println("ошибка при разборе JSON: %w", err)
	}

	infected := make(map[string]bool)
	findInfected(&root, "", infected)

	// count := 0
	// for path := range infected {
	// 	fmt.Println(path)
	// 	count++
	// }
	fmt.Printf("Total infected file and folders: %d\n", len(infected))
	//fmt.Println(infected)

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

func findInfected(node *FoldersPath, currentPath string, infected map[string]bool) {
	fullPath := currentPath + "/" + node.Dir

	hasInfectedFile := false

	// Проверка файлов в текущей директории
	for _, file := range node.Files {
		if strings.HasSuffix(file, ".hack") {
			infected[fullPath+"/"+file] = true
			hasInfectedFile = true
		}
	}

	// Проверка вложенных директорий
	for i := range node.Folders {
		findInfected(&node.Folders[i], fullPath, infected)
	}

	// Если текущая директория заражена, все файлы и папки внутри нее заражены
	if hasInfectedFile || isFolderInfected(node, fullPath, infected) {

		//infected[fullPath] = true
		for _, file := range node.Files {
			infected[fullPath+"/"+file] = true
		}
		for _, subfolder := range node.Folders {
			markAllInfected(&subfolder, fullPath, infected)
		}
	}
}

func isFolderInfected(node *FoldersPath, currentPath string, infected map[string]bool) bool {
	fullPath := currentPath + "/" + node.Dir
	if infected[fullPath] {
		return true
	}
	for _, subfolder := range node.Folders {
		if isFolderInfected(&subfolder, fullPath, infected) {
			return true
		}
	}
	return false
}

func markAllInfected(node *FoldersPath, currentPath string, infected map[string]bool) {
	fullPath := currentPath + "/" + node.Dir

	//infected[fullPath] = true
	for _, file := range node.Files {
		infected[fullPath+"/"+file] = true
	}
	for i := range node.Folders {
		markAllInfected(&node.Folders[i], fullPath, infected)
	}
}
