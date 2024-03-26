package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"runtime/pprof"
)

const (
	filePath       string = "./51_2/"   // Каталог, содержащий входные файлы
	outputFilePath string = "./output/" // Каталог, в который будут записываться выходные файлы
)

// Функция для измерения времени выполнения
func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

// Функция для отслеживания времени выполнения функции
func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

// Функция для вычисления остатка от деления
func penny(price, percent float64) float64 {
	pen1 := (price * percent) / 100
	pen2 := pen1 - float64(int64(pen1))
	return pen2
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

// Функция для вычисления суммы остатков от деления в каждом наборе
func writeSum(reader *bufio.Reader, numAi, numBi int) float64 {
	var outSum float64
	for i := 1; i <= numAi; i++ {
		numAstrC, _ := reader.ReadString('\n')
		numAiC, _ := strconv.Atoi((string(numAstrC[:len(numAstrC)-1])))
		num := penny(float64(numAiC), float64(numBi))
		outSum += num
	}

	return outSum
}

// Функция для чтения двух чисел из файла
func readTwonumber(reader *bufio.Reader) (a, b int) {
	for {
		numAstr, err := reader.ReadString(' ')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("Error reading numAstr:", err)
				return
			}
		}

		numBstr, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("Error reading numAstr:", err)
				return
			}
		}

		numAi, _ := strconv.Atoi((string(numAstr[:len(numAstr)-1])))
		numBi, _ := strconv.Atoi((string(numBstr[:len(numBstr)-1])))
		return numAi, numBi

	}
	return
}

// Функция для чтения входного файла и записи выходных файлов
func readInput() {
	defer duration(track("readInput"))

	// Переменная для хранения пути к текущему файлу
	var filePathFull string

	// Цикл для обработки каждого файла в каталоге "51_2"
	for i := 1; i <= 19; i++ {
		iStr := strconv.Itoa(i)
		filePathFull = filePath + iStr

		// Открытие файла
		f, err := os.Open(filePathFull)
		defer f.Close() // Закрытие файла по окончании работы с ним
		if err != nil {
			fmt.Println("Error opening file:", err)
			break
		}

		// Создание Reader для чтения файла
		rd := bufio.NewReader(f)

		// Чтение количества наборов чисел в файле
		numOfSetsStr, err := rd.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading num of sets:", err)
			continue
		}

		// Преобразование количества наборов из строки в число
		num := strings.TrimSuffix(numOfSetsStr, "\n")
		num1, _ := strconv.Atoi(num)

		// Создание Writer для записи в выходной файл
		outWrite := bufio.NewWriter(writeOutFile(i))
		defer outWrite.Flush() // Сброс буфера при окончании работы с ним

		// Обработка каждого набора чисел
		for j := 1; j <= num1; j++ {
			numAi, numBi := readTwonumber(rd)
			fmt.Fprintf(outWrite, "%.2f\n", writeSum(rd, numAi, numBi))
		}
	}
}

func main() {
	f, err := os.Create("memprofile.prof")
    if err != nil {
        fmt.Println("He удалось создать файл профиля:", err)
        return
    }
    defer f.Close()

    pprof.WriteHeapProfile(f)
	readInput()
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// const (
// 	filePath       string = "./51_2/"
// 	outputFilePath string = "./output/"
// )

// func duration(msg string, start time.Time) {
// 	log.Printf("%v:%v\n", msg, time.Since(start))
// }

// func track(msg string) (string, time.Time) {
// 	return msg, time.Now()
// }

// func penny(price, percent float64) float64 {
// 	pen1 := (price * percent) / 100
// 	pen2 := pen1 - float64(int64(pen1))
// 	return pen2
// }

// func writeOutFile(iter int) *os.File {

// 	pathToFile := fmt.Sprintf("%s%d.txt", outputFilePath, iter)

// 	outFile, err := os.OpenFile(pathToFile, os.O_WRONLY|os.O_CREATE, 0666)
// 	//defer outFile.Close()
// 	if err != nil {
// 		fmt.Println("File does not exists or cannot be created!!", err)
// 		os.Exit(1)
// 	}
// 	return outFile
// }

// func writeSum(reader *bufio.Reader, numAi, numBi int) float64 {
// 	var outSum float64
// 	for i := 1; i <= numAi; i++ {
// 		numAstrC, _ := reader.ReadString('\n')
// 		numAiC, _ := strconv.Atoi((string(numAstrC[:len(numAstrC)-1])))
// 		num := penny(float64(numAiC), float64(numBi))
// 		outSum += num
// 	}

// 	return outSum
// }

// func readTwonumber(reader *bufio.Reader) (a, b int) {
// 	for {
// 		numAstr, err := reader.ReadString(' ')
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			} else {
// 				fmt.Println("Error reading numAstr:", err)
// 				return
// 			}
// 		}

// 		numBstr, err := reader.ReadString('\n')
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			} else {
// 				fmt.Println("Error reading numAstr:", err)
// 				return
// 			}
// 		}

// 		numAi, _ := strconv.Atoi((string(numAstr[:len(numAstr)-1])))
// 		numBi, _ := strconv.Atoi((string(numBstr[:len(numBstr)-1])))
// 		return numAi, numBi

// 	}
// 	return
// }

// // reading file
// func readInput() {

// 	defer duration(track("readInput"))

// 	//Path to file
// 	var filePathFull string

// 	//cycle for each file names in folder "51_2"
// 	for i := 1; i <= 19; i++ {
// 		iStr := strconv.Itoa(i)
// 		filePathFull = filePath + iStr

// 		//opening the file
// 		f, err := os.Open(filePathFull)
// 		defer f.Close()
// 		if err != nil {
// 			fmt.Println("ERRORRRR", err)
// 		}

// 		//creating Reader
// 		rd := bufio.NewReader(f)

// 		// getting first number in file = number of sets
// 		numOfSets, err := rd.ReadString('\n')
// 		if err != nil {
// 			fmt.Println("Error reading num of sets:", err)
// 			continue
// 		}

// 		num := strings.TrimSuffix(numOfSets, "\n")
// 		num1, _ := strconv.Atoi(num)

// 		outWrite := bufio.NewWriter(writeOutFile(i))
// 		defer outWrite.Flush()

// 		for j := 1; j <= num1; j++ {
// 			numAi, numBi := readTwonumber(rd)
// 			fmt.Fprintf(outWrite, "%.2f\n", writeSum(rd, numAi, numBi))
// 		}
// 	}
// }

// func main() {

// 	readInput()

// }
