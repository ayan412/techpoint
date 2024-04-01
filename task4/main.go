package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	filePath       string = "./63_4/"   // Каталог, содержащий входные файлы
	outputFilePath string = "./output/" // Каталог, в который будут записываться выходные файлы
)

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
		defer outWrite.Flush()

		for j := 1; j <= numOfSetsInt; j++ {
			readDim(rdr)
		}

	}
}

func checkDim(numbersStr []string) (vertic, horizon int) {
	// если входной срез меньше 3
	if len(numbersStr) > 2 {
		fmt.Println("Wrong dimensions")
		return 0, 0
	}
	// Преобразование байтов в целое число
	verticalDim, _ := strconv.Atoi(numbersStr[0])
	horizontalDim, _ := strconv.Atoi(numbersStr[1])

	return verticalDim, horizontalDim
}

func readDim(rdr *bufio.Reader) {

	// Получение размеров склада: строки и столбцы склада
	rowStrWithNL, err := rdr.ReadString('\n')
	//fmt.Printf("rowStrWithNL:%vlentgh:%v\n", rowStrWithNL, len(rowStrWithNL)) // rowStrWithNL:23 99 (+\n) lentgh:6
	if err != nil {
		fmt.Println("Error in reading row with dimensions:", err)
	}

	rowStr := strings.TrimSpace(rowStrWithNL)
	//fmt.Printf("Length after TRIM:%v\n", len(rowStr)) //Length after TRIM:5
	// Преобразование строки
	numbersStr := strings.Split(rowStr, " ")
	//fmt.Printf("numOfSetsStr:%v lentgh:%v\n", numbersStr, len(numbersStr)) //numbersStr:[23 99] lentgh:2

	// считываем и выводим только размеры нужного склада в зав-ти от длины его строк
	verticDim, horizonDim := checkDim(numbersStr)
	fmt.Println("строки и столбцы:", verticDim, horizonDim)
	for j := 1; j <= verticDim; j++ {
		rowsOfStore, err := rdr.ReadString('\n')
		if err != nil {
			fmt.Println("ERRRORS:", err)
		}
		fmt.Printf("%v", rowsOfStore)
	}

}

// внести вывод в двумерный масссив или срез и уже работать с этим типом данных
