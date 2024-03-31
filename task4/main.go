package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	//"reflect"
	"strconv"
	"strings"
)

const (
	filePath       string = "./63_4/"   // Каталог, содержащий входные файлы
	outputFilePath string = "./output/" // Каталог, в который будут записываться выходные файлы
)

// Функция для чтения двух чисел из файла
func readTwoNumbers(reader *bufio.Reader) (a, b int) {
	for {
		numAstr, err := reader.ReadString(' ')
		if err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				break
			} else {
				fmt.Println("Error reading numAstr:", err)
				return
			}
		}

		numBstr, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				break
			} else {
				fmt.Println("Error reading numBstr:", err)
				return
			}
		}

		// Удаление символа новой строки '\n'
		numAi, err := strconv.Atoi((string(numAstr[:len(numAstr)-1])))
		if err != nil {
			fmt.Printf("Error while trimming:%s", err)
		}
		numBi, err := strconv.Atoi((string(numBstr[:len(numBstr)-1])))
		if err != nil {
			fmt.Printf("Error while trimming:%s", err)
		}
		return numAi, numBi

	}
	return
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

func main() {
	// Переменная для хранения пути к текущему файлу
	var filePathFull string

	// Цикл для обработки каждого файла в каталоге "59_3"
	for i := 1; i <= 1; i++ {
		iStr := strconv.Itoa(i)
		// Путь до файла
		filePathFull = fmt.Sprintf("%s%s", filePath, iStr)

		// Открытие файла
		f, err := os.Open(filePathFull)
		if err != nil {
			fmt.Println("Error while opening the file", err)
			break
		}
		defer f.Close() // Закрытие файла по окончании работы с ним

		// Создание Reader для чтения файла
		rdr := bufio.NewReader(f)

		// Чтение количества наборов входных данных в файле
		numOfSetsStr, err := rdr.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading amount of sets:", err)
			break
		}
		//Преобразование количества наборов из строки в число
		numOfSetsStrTrim := strings.TrimSuffix(numOfSetsStr, "\n")
		numOfSetsInt, _ := strconv.Atoi(numOfSetsStrTrim)
		fmt.Println("amount of sets:", numOfSetsInt)

		// Создание Writer для записи в выходной файл
		outWrite := bufio.NewWriter(writeOutFile(i))
		defer outWrite.Flush() // Сброс буфера при окончании работы с ним

		// Получение размеров склада: строки и столбцы склада
		rowStrWithNL, err := rdr.ReadString('\n')
		fmt.Println(rowStrWithNL)
		if err != nil {
			fmt.Println("Error in reading row with dimensions:", err)
		}
		
		rowStr := strings.TrimSpace(rowStrWithNL)
		fmt.Println(rowStr)
		// Преобразование строки в числа
		numbersStr := strings.Split(rowStr, " ")
		numbers := make([]int, len(numbersStr))

		for i, numStr := range numbersStr {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Ошибка преобразования числа:", err)
				return
			}
			numbers[i] = num
		}

		// Выводим результат
		fmt.Println(numbers[1])

		if err != nil {
			fmt.Println("Error in reading the slice of bytes:", err)
		}
		// размеры
		verticDim, horizonDim := checkDim(rowStr)
		fmt.Println("строки и столбцы:", verticDim, horizonDim)
		for j := 1; j <= verticDim; j++ {
			rowsOfStore, err := rdr.ReadString('\n')
			if err != nil {
				fmt.Println("ERRRORS:", err)
			}
			fmt.Printf("%v", rowsOfStore)
		}
	}
}

func checkDim(numOfbytes string) (vertic, horizon int) {
	// если входной срез меньше 3
	if len(numOfbytes) < 3 {
		fmt.Println("Wrong dimensions")
		return 0, 0
	}
	// Преобразование байтов в целое число
	verticalDim, _ := strconv.Atoi(string(numOfbytes[0]))
	horizontalDim, _ := strconv.Atoi(string(numOfbytes[1]))

	return verticalDim, horizontalDim
}

//придумать цикл, который бы ориентировался на длину среза байтов > 5 байта
