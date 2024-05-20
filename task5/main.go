package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"
)

const (
	filePath       string = "./71_5/"   // Папка, содержащая входные файлы для теста
	outputFilePath string = "./output/" // Папка, в которую будут записываться выходные файлы
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
func writeToOutputFile(iter int) (*os.File, error) {
	// Формирование пути к выходному файлу
	pathToFile := fmt.Sprintf("%s%d.txt", outputFilePath, iter)

	// Открытие или создание файла
	outFile, err := os.OpenFile(pathToFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("failed to open or create file:", err)
		return nil, err
	}
	return outFile, nil
}

type FoldersPath struct {
	Dir     string        `json:"dir"`
	Files   []string      `json:"files"`
	Folders []FoldersPath `json:"folders"`
}

func main() {
	// Профилирование памяти
	f, err := os.Create("memprofile.prof")
	if err != nil {
		fmt.Println("Не удалось создать файл профиля:", err)
		return
	}
	defer f.Close()
	runtime.GC()
	pprof.WriteHeapProfile(f)

	defer duration(track("Total execution time"))

	// Обработка каждого файла в каталоге
	for i := 1; i <= 11; i++ {
		// Чтение набора данных из файла
		numSets, rdr, outwrt, err := readSets(i)
		if err != nil {
			fmt.Println("Ошибка при чтении наборов данных:", err)
			continue
		}

		// Обработка каждого набора данных
		for j := 1; j <= numSets; j++ {
			numRows := readRows(rdr)
			readJson(numRows, rdr, outwrt)
		}

		// Сброс буфера для записи всех данных в файл
		outwrt.Flush()
	}
}

// Функция для чтения первого числа - количества наборов данных
func readSets(fileIndex int) (int, *bufio.Reader, *bufio.Writer, error) {
	var numSets int
	var rdr *bufio.Reader

	// Формирование пути к конкретному файлу с наборами данных
	iStr := strconv.Itoa(fileIndex)
	filePathFull := fmt.Sprintf("%s%s", filePath, iStr)

	// Открытие файла
	file, err := os.Open(filePathFull)
	if err != nil {
		return 0, nil, nil, fmt.Errorf("error opening the file: %w", err)
	}

	// Создание Reader для чтения файла
	rdr = bufio.NewReader(file)

	// Чтение количества наборов входных данных в файле
	numOfSetsStr, err := rdr.ReadString('\n')
	if err != nil {
		file.Close()
		return 0, nil, nil, fmt.Errorf("Error reading amount of sets: %w", err)
	}

	// Преобразование количества наборов из строки в число
	numOfSetsStr = strings.TrimSuffix(numOfSetsStr, "\n")
	numOfSetsInt, err := strconv.Atoi(numOfSetsStr)
	if err != nil {
		file.Close()
		return 0, nil, nil, fmt.Errorf("failed to convert sets count to int: %w", err)
	}

	// Создание Writer для записи в соответствующий выходной файл
	outFile, err := writeToOutputFile(fileIndex)
	if err != nil {
		file.Close()
		return 0, nil, nil, fmt.Errorf("failed to create output file: %w", err)
	}

	outWrite := bufio.NewWriter(outFile)

	// Количество наборов входных данных
	numSets = numOfSetsInt

	return numSets, rdr, outWrite, nil
}

// Функция для чтения JSON строк и записи результатов
func readJson(numRows int, rdr *bufio.Reader, outwrt *bufio.Writer) {
	// Инициализация структуры для избежания пустого указателя
	var root FoldersPath

	// Чтение строк с JSON данными
	var jsonBuilder strings.Builder
	for j := 1; j <= numRows; j++ {
		line, err := rdr.ReadString('\n')
		if err != nil {
			fmt.Println("ошибка при чтении строки JSON:", err)
			return
		}
		jsonBuilder.WriteString(line)
	}

	// Разбор JSON
	if err := json.Unmarshal([]byte(jsonBuilder.String()), &root); err != nil {
		fmt.Println("ошибка при разборе JSON:", err)
		return
	}

	// Поиск зараженных файлов
	infected := make(map[string]bool)
	findInfected(&root, "", infected)

	// Запись количества зараженных файлов в выходной файл
	_, err := fmt.Fprintf(outwrt, "%d\n", len(infected))
	if err != nil {
		fmt.Println("ошибка при записи в файл:", err)
	}
}

// Функция для чтения количества строк с описанием директорий
func readRows(rdr *bufio.Reader) int {
	strOfRow, err := rdr.ReadString('\n')
	if err != nil {
		fmt.Println("Error in reading row with sets of rows", err)
		return 0
	}
	strOfRows := strings.TrimSuffix(strOfRow, "\n")
	numOfRows, err := strconv.Atoi(strOfRows)
	if err != nil {
		fmt.Println("Error in type converting", err)
		return 0
	}
	return numOfRows
}

// Рекурсивная функция для поиска зараженных файлов и директорий
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
		for _, file := range node.Files {
			infected[fullPath+"/"+file] = true
		}
		for _, subfolder := range node.Folders {
			markAllInfected(&subfolder, fullPath, infected)
		}
	}
}

// Функция для проверки, является ли директория зараженной
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

// Рекурсивная функция для пометки всех файлов и папок внутри зараженной директории
func markAllInfected(node *FoldersPath, currentPath string, infected map[string]bool) {
	fullPath := currentPath + "/" + node.Dir

	for _, file := range node.Files {
		infected[fullPath+"/"+file] = true
	}
	for i := range node.Folders {
		markAllInfected(&node.Folders[i], fullPath, infected)
	}
}
