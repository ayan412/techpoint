package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"
)

const (
	filePath       string = "./67_6/"   // Папка, содержащая входные файлы для теста
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

// readMatrixFromFile reads the matrix from a given file
func readMatrixFromFile(rdr *bufio.Reader) ([][]int, error) {

	numAi, numBi := readTwonumber(rdr)
	// fmt.Println(numAi, numBi)
	// Срез (из кол-ва рядов) срезов - [[] [] [] [] [] [] []]

	mtx := make([][]int, numAi)
	for i := 0; i < numAi; i++ {
		lineStr, err := rdr.ReadString('\n')
		//fmt.Print(lineStr)
		if err != nil {
			return nil, err
		}

		lineStr = strings.TrimSpace(lineStr)

		values := strings.Split(lineStr, "")
		//fmt.Println(values)

		// if len(values) > numBi {
		// 	return nil, fmt.Errorf("invalid number of values in row %d", i+1)
		// }
		
		//fmt.Println(lineStr)
		//line, _ := strconv.Atoi((string(lineStr[:len(lineStr)-1])))
		row := make([]int, numBi)
		for index, val := range values {
			//fmt.Println(val)
			num, err := strconv.Atoi(val)
			//fmt.Println(num)
			if err != nil {
				return nil, err
			}
			row[index] = num
			//mtx[i-1] = row
		}
		mtx[i] = row
	}
	//fmt.Println(mtx)
	return mtx, nil
}

func readTwonumber(reader *bufio.Reader) (a, b int) {
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read line: %v", err)
	}
	line = strings.TrimSpace(line)
	numbers := strings.Split(line, " ")
	if len(numbers) != 2 {
		log.Fatalf("Expected two numbers in the line, got: %v", numbers)
	}
	a, err = strconv.Atoi(numbers[0])
	if err != nil {
		log.Fatalf("Failed to convert first number to int: %v", err)
	}
	b, err = strconv.Atoi(numbers[1])
	if err != nil {
		log.Fatalf("Failed to convert second number to int: %v", err)
	}
	return a, b
	// for {
	// 	numAstr, err := reader.ReadString(' ')
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		} else {
	// 			fmt.Println("Error reading numAstr:", err)
	// 			return
	// 		}
	// 	}

	// 	numBstr, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		} else {
	// 			fmt.Println("Error reading numBstr:", err)
	// 			return
	// 		}
	// 	}

	// 	numAi, _ := strconv.Atoi((string(numAstr[:len(numAstr)-1])))
	// 	numBi, _ := strconv.Atoi((string(numBstr[:len(numBstr)-1])))
	// 	return numAi, numBi

	// }
	// return
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
	for i := 9; i <= 9; i++ {
		// Чтение набора данных из файла
		numSets, rdr, _, err := readSets(i)
		if err != nil {
			log.Fatalf("Failed to read sets from file: %v\n", err)
		}

		// Обработка каждого набора данных
		for j := 1; j <= numSets; j++ {
			matrix, err := readMatrixFromFile(rdr)
			if err != nil {
				log.Fatalf("Failed to read matrix from file: %v\n", err)
			}
			bestRow, bestCol := findBestRemoval(matrix)
			fmt.Printf("%d %d\n", bestRow, bestCol)
		}
	}
}

// findWorst returns the minimun
func findWorstGrade(matrix [][]int) int {
	worst := math.MaxInt32
	for _, row := range matrix {
		for _, grade := range row {
			if grade < worst {
				worst = grade
			}
		}
	}
	return worst
}

func removeRowAndColumn(matrix [][]int, rowIndex, colIndex int) [][]int {
	newMatrix := make([][]int, 0, len(matrix)-1)
	for i, row := range matrix {
		if i != rowIndex {
			newRow := make([]int, 0, len(matrix)-1)
			for j, grade := range row {
				if j != colIndex {
					newRow = append(newRow, grade)
				}
			}
			newMatrix = append(newMatrix, newRow)
		}
	}
	return newMatrix
}

// findBestRemoval determines which row or column removal maximizes the minimum grade
func findBestRemoval(matrix [][]int) (int, int) {
	//bestMatrix := matrix
	bestWorstGrade := findWorstGrade(matrix)
	bestRow, bestCol := -1, -1

	// Test removing each row and each column combination
	for i := range matrix {
		for j := range matrix[0] {
			tempMatrix := removeRowAndColumn(matrix, i, j)
			worstGrade := findWorstGrade(tempMatrix)
			if worstGrade > bestWorstGrade {
				bestWorstGrade = worstGrade
				//bestMatrix = tempMatrix
				bestRow = i
				bestCol = j
			}
		}
	}
	return bestRow + 1, bestCol + 1
}
