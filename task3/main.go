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

	// Переменная для хранения пути к текущему файлу
	var filePathFull string

	// Цикл для обработки каждого файла в каталоге "59_3"
	for i := 1; i <= 18; i++ {
		iStr := strconv.Itoa(i)
		// Путь до файла
		filePathFull = fmt.Sprintf("%s%s", filePath, iStr)

		// Открытие файла
		f, err := os.Open(filePathFull)
		if err != nil {
			fmt.Println("Error opening the file", err)
			continue
		}
		defer f.Close() // Закрытие файла по окончании работы с ним

		// Создание Reader для чтения файла
		rd := bufio.NewReader(f)

		// Чтение количества наборов в файле
		numOfSetsStr, err := rd.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading num of sets:", err)
			continue
		}

		//Преобразование количества наборов из строки в число
		numOfSetsStrTrim := strings.TrimSuffix(numOfSetsStr, "\n")
		numOfSetsInt, _ := strconv.Atoi(numOfSetsStrTrim)
		fmt.Println("first number:", numOfSetsInt)

		// Создание Writer для записи в выходной файл
		outWrite := bufio.NewWriter(writeOutFile(i))
		defer outWrite.Flush() // Сброс буфера при окончании работы с ним

		// Обработка каждого набора строк (задач)
		for j := 1; j <= numOfSetsInt; j++ {
			setOftasks, err := rd.ReadBytes('\n')
			if err != nil {
				fmt.Println("error in reading line", err)
				break
			}
			// Запись результата в выходной файл
			fmt.Fprintf(outWrite, "%s\n", checkTaskStatus(setOftasks))
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

// Функция для проверки корректности последовательности задач в каждой строке
func checkTaskStatus(setOftasks []byte) string {
	// Запись статуса задач в мапу
	statusCode := make(map[string]string)
	// Счетчик ошибок в строке
	ErrCount := 0
	// Цикл для работы с каждой задачой в текущей строке
	for _, val := range setOftasks {
		switch val {
		case 'M':
			// Задачи НЕТ (empty) ИЛИ (пере)запущенную задачу можно только отменить ИЛИ она должна быть запущена ИЛИ запустить снова после закрытия
			if _, ok := statusCode["task"]; !ok || statusCode["task"] == "cancel" || statusCode["task"] == "close" {
				statusCode["task"] = "start"
			} else {
				ErrCount++
			}
		case 'R':
			// запущенную задачу можно перезапустить
			if _, ok := statusCode["task"]; ok && statusCode["task"] == "start" {
				statusCode["task"] = "restart"
			} else {
				ErrCount++
			}
		case 'C':
			// запущенную задачу можно отменить ИЛИ перезапущенная задача отменяется
			if _, ok := statusCode["task"]; ok && statusCode["task"] == "start" || statusCode["task"] == "restart" {
				statusCode["task"] = "cancel"
			} else {
				ErrCount++
			}
		case 'D':
			// запущенную задачу можно завершить
			if _, ok := statusCode["task"]; ok && statusCode["task"] == "start" {
				statusCode["task"] = "close"
			} else {
				ErrCount++
			}
		case '\n':
			// Для случая символа новой строки - символ есть И перед ним был статус "Завершено"
			if _, ok := statusCode["task"]; ok && statusCode["task"] == "close" {
			} else {
				ErrCount++
			}
		}

	}
	if ErrCount >= 1 {
		return "NO"
	} else {
		return "YES"
	}
}
