// Дописать комментарии 26032024

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	filePath       string = "./59_3/"   // Каталог, содержащий входные файлы
	outputFilePath string = "./output/" // Каталог, в который будут записываться выходные файлы
)

func main() {

	//	var noErrors bool

	// Переменная для хранения пути к текущему файлу
	var filePathFull string

	// Цикл для обработки каждого файла в каталоге "59_3"
	for i := 1; i <= 18; i++ {
		iStr := strconv.Itoa(i)
		filePathFull = fmt.Sprintf("%s%s", filePath, iStr)
		//.Println(filePathFull)

		// Открытие файла
		f, err := os.Open(filePathFull)
		if err != nil {
			fmt.Println("Error opening the file", err)
			continue
		}
		defer f.Close() // Закрытие файла по окончании работы с ним

		// Создание Reader для чтения файла
		rd := bufio.NewReader(f)

		// Чтение количества наборов чисел в файле
		numOfSetsStr, err := rd.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading num of sets:", err)
			continue
		}

		//Преобразование количества наборов из строки в число
		num := strings.TrimSuffix(numOfSetsStr, "\n")
		numInt, _ := strconv.Atoi(num)
		fmt.Println("first number:", numInt)

		// Создание Writer для записи в выходной файл
		outWrite := bufio.NewWriter(writeOutFile(i))
		defer outWrite.Flush() // Сброс буфера при окончании работы с ним

		//fmt.Fprintf(outWrite, "%s", checkCode(rd, numInt))
		// Обработка каждого набора чисел
		//checkCode(rd, numInt)

		for j := 1; j <= numInt; j++ {
			setOftasks, err := rd.ReadBytes('\n')
			if err != nil {
				fmt.Println("error in reading line", err)
				break
			}
			//fmt.Println(setOftasks)
			fmt.Fprintf(outWrite, "%s\n", checkTaskStatus(setOftasks))
			//return result

		}

	}
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

// Функция 
func checkTaskStatus(setOftasks []byte) string {
	statusCode := make(map[string]string)
	ErrCount := 0
	for _, val := range setOftasks {
		//fmt.Println(val)
		switch val {
		case 'M':
			// (пере)запущенную задачу можно только отменить ИЛИ она должна быть запущена ИЛИ запустить снова после закрытия
			if _, ok := statusCode["task"]; !ok || statusCode["task"] == "cancel" || statusCode["task"] == "close" {
				statusCode["task"] = "start"
				//fmt.Printf("previous was: %v\n", mm)
			} else {
				//fmt.Println("ERRRRROOOOR")
				ErrCount++
			}
		case 'R':
			// запущенную задачу можно перезапустить
			if _, ok := statusCode["task"]; ok && statusCode["task"] == "start" {
				statusCode["task"] = "restart"
				//fmt.Printf("previous was: %v\n", rr)
			} else {
				//fmt.Println("ERRRRROOOOR")
				ErrCount++
			}
		case 'C':
			// запущенную задачу можно отменить И перезапущенная задача отменяется
			if _, ok := statusCode["task"]; ok && statusCode["task"] == "start" || statusCode["task"] == "restart" {
				statusCode["task"] = "cancel"
				//fmt.Printf("previous was: %v\n", cc)
			} else {
				//fmt.Println("ERRRRROOOOR")
				ErrCount++
			}
		case 'D':
			// запущенную задачу можно завершить
			if _, ok := statusCode["task"]; ok && statusCode["task"] == "start" {
				statusCode["task"] = "close"
				//fmt.Printf("current is: %v\n", string(val))
				//fmt.Printf("previous was: %v\n", dd)
			} else {
				//fmt.Println("ERRRRROOOOR")
				ErrCount++
			}
		case '\n':
			//
			if _, ok := statusCode["task"]; ok && statusCode["task"] == "close" {
				//fmt.Printf("NEW LINE - previous was: %v\n", ss)
			} else {
				//fmt.Println("END OF LINE ERRRRROOOOR!!!!! - TASK WAS NOT CLOSED")
				ErrCount++
			}
			//fmt.Println("Количество ошибок в наборе задач:", ErrCount)
		}

	}
	if ErrCount >= 1 {
		return "NO"
	} else {
		return "YES"
	}

}

/*
	statusCode := make(map[string]string)
		for _, val := range setOftasks {
			fmt.Println(val)
			switch val {
			case 'M':
				// (пере)запущенную задачу можно только отменить ИЛИ она должна быть запущена ИЛИ запустить снова после закрытия
				if mm, ok := statusCode["task"]; !ok || statusCode["task"] == "cancel" || statusCode["task"] == "close" {
					statusCode["task"] = "start"
					fmt.Printf("previous was: %v\n", mm)
				} else {
					fmt.Println("ERRRRROOOOR")
					continue
				}
			case 'R':
				// запущенную задачу можно перезапустить
				if rr, ok := statusCode["task"]; !ok || statusCode["task"] == "start" {
					statusCode["task"] = "restart"
					fmt.Printf("previous was: %v\n", rr)
				} else {
					fmt.Println("ERRRRROOOOR")
					continue
				}
			case 'C':
				// запущенную задачу можно отменить ИЛИ перезапущенная задача отменяется
				if cc, ok := statusCode["task"]; !ok || statusCode["task"] == "start" || statusCode["task"] == "restart" {
					statusCode["task"] = "cancel"
					fmt.Printf("previous was: %v\n", cc)
				} else {
					fmt.Println("ERRRRROOOOR")
					continue
				}
			case 'D':
				// запущенную задачу можно завершить
				if dd, ok := statusCode["task"]; !ok || statusCode["task"] == "start" {
					statusCode["task"] = "close"
					fmt.Printf("current is: %v\n", string(val))
					fmt.Printf("previous was: %v\n", dd)
				} else {
					fmt.Println("ERRRRROOOOR")
					continue
				}
			case '\n':
				//
				if ss, ok := statusCode["task"]; ok && statusCode["task"] == "close" {
					fmt.Printf("NEW LINEE - previous was: %v\n", ss)
				} else {
					fmt.Println("NEW LINE - ERRRRROOOOR!!!!!")

				}
				}
			}

		}

	}

*/
