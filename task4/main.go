package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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

		allSlices := make([][][]string, numOfSetsInt)
		for j := 1; j <= numOfSetsInt; j++ {
			allSlices[j-1] = readDim(rdr)
		}
		fmt.Println(reflect.TypeOf(allSlices[0]))
		fmt.Println(len(allSlices[0][0]))
		fmt.Println(allSlices[0][0])
		//fmt.Println(allSlices[0])

		robotA := robotPosition{2, 3}
		robotB := robotPosition{3, 4}

		pathA := bfs(robotA, robotPosition{0, 0}, "a")
		pathB := bfs(robotB, robotPosition{len(wHouse) - 1, len(wHouse[0]) - 1}, "b")

		for point, robot := range pathA.route {
			wHouse[point.x][point.y] = robot
		}

		for point, robot := range pathB.route {
			wHouse[point.x][point.y] = robot
		}

		for _, row := range wHouse {
			fmt.Println(row)
		}
	}
}

func checkDim(numbersStr []string) (vertic, horizon int) {
	// если входной срез больше 2 ошибка
	if len(numbersStr) > 2 {
		fmt.Println("Wrong dimensions")
		return 0, 0
	}
	// Преобразование байтов в целое число
	verticalDim, _ := strconv.Atoi(numbersStr[0])
	horizontalDim, _ := strconv.Atoi(numbersStr[1])

	return verticalDim, horizontalDim
}

func readDim(rdr *bufio.Reader) [][]string {

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
	var mdArray [3][2]int
	result := make(map[string][]int)

	verticDim, horizonDim := checkDim(numbersStr)
	//fmt.Println("строки и столбцы:", verticDim, horizonDim)
	mdArray[0][0] = verticDim
	mdArray[0][1] = horizonDim
	result["MAX"] = mdArray[0][:]

	// Срез (из кол-ва рядов) срезов - [[] [] [] [] [] [] []]
	matrix := make([][]string, verticDim)
	//fmt.Println(matrix)

	for j := 1; j <= verticDim; j++ {
		rowOfWh, err := rdr.ReadString('\n')
		rowOfWh = strings.Trim(rowOfWh, "\n")
		//fmt.Println(rowOfWh)
		if err != nil {
			fmt.Println("ERRRORS:", err)
		}
		slice := make([]string, horizonDim)
		for index, value := range rowOfWh {

			slice[index] = string(value)

			switch value {
			case 'A':
				mdArray[1][0] = j
				mdArray[1][1] = index + 1
			case 'B':
				mdArray[2][0] = j
				mdArray[2][1] = index + 1
			}
		}
		matrix[j-1] = slice
	}

	result["A"] = mdArray[1][:]
	result["B"] = mdArray[2][:]
	fmt.Println(result)

	subtracPositions(result)
	return matrix
}

func subtracPositions(result map[string][]int) {
	// Извлекаем срезы для MAX, A и B
	maxSlice := result["MAX"]
	aSlice := result["A"]
	bSlice := result["B"]

	// Срез для хранения разности чисел между MAX и A
	resultASlice := []int{}

	// Разность чисел между MAX и A
	for i := 0; i < len(maxSlice); i++ {
		resultASlice = append(resultASlice, maxSlice[i]-aSlice[i])
	}

	// Срез для хранения разности чисел между MAX и B
	resultBSlice := []int{}

	// Разность чисел между MAX и B
	for i := 0; i < len(maxSlice); i++ {
		resultBSlice = append(resultBSlice, maxSlice[i]-bSlice[i])
	}
	// Сумма разницы чисел между МАХ и А
	sumA := 0
	for i := 0; i < len(resultASlice); i++ {
		sumA += resultASlice[i]
	}
	fmt.Println("len A", sumA)
	// Сумма разницы чисел между МАХ и В
	sumB := 0
	for i := 0; i < len(resultBSlice); i++ {
		sumB += resultBSlice[i]
	}
	fmt.Println("len B", sumB)
	if sumA > sumB {
		// алгоритм по которому будет дописываться путь робота в точку 0;0

	} else {
		// алгоритм по которому будет дописываться путь робота в точку MAX
	}

	// Выводим разность
	fmt.Println("Разность чисел между MAX и A:", resultASlice)
	fmt.Println("Разность чисел между MAX и B:", resultBSlice)

}

type robotPosition struct {
	x, y int
}

type robotPath struct {
	success bool
	route   map[robotPosition]string
}

var wHouse [][]string

func isValidCell(x, y int) bool {
	if x < 0 || x >= len(wHouse) || y < 0 || y >= len(wHouse[0]) {
		return false
	} else {
		if wHouse[x][y] == "x" {
			return false
		}
	}
	return true
}

// Функция обхода в ширину (BFS)
func bfs(start, end robotPosition, name string) robotPath {
	visited := make(map[robotPosition]bool)
	queue := []robotPosition{start}
	route := make(map[robotPosition]string)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			return robotPath{true, route}
		}

		// Перемещаемся во все соседние клетки
		neighbors := []robotPosition{{current.x - 1, current.y}, {current.x + 1, current.y}, {current.x, current.y - 1}, {current.x, current.y + 1}}
		for _, neighbor := range neighbors {
			if isValidCell(neighbor.x, neighbor.y) && !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
				route[neighbor] = name
			}
		}
	}
	return robotPath{false, nil}
}
