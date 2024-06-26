package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"
)

const (
	filePath       string = "./63_4/"   // Каталог, содержащий входные файлы
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

var result = make(map[string][]int)

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
	f, err := os.Create("memprofile.prof")
	if err != nil {
		fmt.Println("Не удалось создать файл профиля:", err)
		return
	}
	defer f.Close()

	pprof.WriteHeapProfile(f)

	defer duration(track("readInput"))
	// Переменная для хранения пути к текущему файлу
	var filePathFull string

	// Цикл для обработки каждого файла в каталоге "63_4"
	for i := 1; i <= 17; i++ {
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
		//fmt.Println("amount of sets:", numOfSetsInt)

		// Создание Writer для записи в выходной файл
		outWrite := bufio.NewWriter(writeOutFile(i))
		defer outWrite.Flush()

		//Запись готовых путей в выходной файл
		createRoute(numOfSetsInt, rdr, outWrite)
	}
}

func createRoute(numOfSetsInt int, rdr *bufio.Reader, outWrite *bufio.Writer) {
	allSlices := make([][][]string, numOfSetsInt)
	for j := 1; j <= numOfSetsInt; j++ {
		allSlices[j-1] = readDim(rdr)
		wHouse := allSlices[j-1]
		//fmt.Println("wH:", wHouse)

		maxSlice := result["MAX"]
		aSlice := result["A"]
		bSlice := result["B"]

		var robotA, robotB, end point
		robotA.x = aSlice[0] - 1
		robotA.y = aSlice[1] - 1

		robotB.x = bSlice[0] - 1
		robotB.y = bSlice[1] - 1

		end.x = maxSlice[0] - 1
		end.y = maxSlice[1] - 1

		//fmt.Printf("%v%v%v\n", aSlice, bSlice, maxSlice)

		//fmt.Println("robots:", robotA, robotB)

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
		//fmt.Println("len A:", sumA)
		// Сумма разницы чисел между МАХ и В
		sumB := 0
		for i := 0; i < len(resultBSlice); i++ {
			sumB += resultBSlice[i]
		}
		//fmt.Println("len B:", sumB)

		switch {

		case sumA > sumB:
			{
				//fmt.Println("sumA > sumB")
				// алгоритм по которому будет дописываться путь робота A к 0; B к MAX
				//Start walking the maze run
				steps := run(wHouse, robotA, point{0, 0})

				//Give a path according to steps
				wHouse = changeMatrix(wHouse, steps, "a", robotA, point{0, 0})

				//fmt.Println("steps:")
				// for x := range steps {
				// 	for y := range steps[x] {
				// 		fmt.Printf("%4d", steps[x][y])
				// 	}
				// 	fmt.Println()
				// }

				// fmt.Println("changed maze:")
				// for i := range wHouse {
				// 	for j := range wHouse[i] {
				// 		fmt.Printf("%s", wHouse[i][j])
				// 	}
				// 	fmt.Println()
				// }

				steps = run(wHouse, robotB, end)

				wHouse1 := changeMatrix(wHouse, steps, "b", robotB, end)

				//fmt.Println("steps:")
				// for x := range steps {
				// 	for y := range steps[x] {
				// 		fmt.Printf("%4d", steps[x][y])
				// 	}
				// 	fmt.Println()
				// }

				//fmt.Println("changed maze is done:")
				for x := range wHouse1 {
					for y := range wHouse1[x] {
						fmt.Fprintf(outWrite, "%s", wHouse1[x][y])
					}
					fmt.Fprintf(outWrite, "\n")
				}

			}

		case sumB > sumA:

			{
				//fmt.Println("sumB > sumA")
				// алгоритм по которому будет дописываться путь робота B к 0; А к МАХ
				//Start walking the maze run
				steps := run(wHouse, robotA, end)

				//Give a path according to steps
				wHouse = changeMatrix(wHouse, steps, "a", robotA, end)

				//fmt.Println("steps:")
				// for x := range steps {
				// 	for y := range steps[x] {
				// 		fmt.Printf("%4d", steps[x][y])
				// 	}
				// 	fmt.Println()
				// }

				// fmt.Println("changed maze:")
				// for x := range wHouse {
				// 	for y := range wHouse[x] {
				// 		fmt.Printf("%s", wHouse[x][y])
				// 	}
				// 	fmt.Println()
				// }

				steps = run(wHouse, robotB, point{0, 0})

				wHouse1 := changeMatrix(wHouse, steps, "b", robotB, point{0, 0})

				//fmt.Println("steps:")
				// for x := range steps {
				// 	for y := range steps[x] {
				// 		fmt.Printf("%4d", steps[x][y])
				// 	}
				// 	fmt.Println()
				// }

				//fmt.Println("changed maze is done:")
				for x := range wHouse1 {
					for y := range wHouse1[x] {
						fmt.Fprintf(outWrite, "%s", wHouse1[x][y])
					}
					fmt.Fprintf(outWrite, "\n")
				}
			}
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
	//result := make(map[string][]int)

	verticDim, horizonDim := checkDim(numbersStr)
	//fmt.Println("строки и столбцы:", verticDim, horizonDim)
	mdArray[0][0] = verticDim
	mdArray[0][1] = horizonDim
	//fmt.Print(mdArray)
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
	//fmt.Println("result:",result)

	//subtracPositions(result)
	//fmt.Println(mdArray)
	//fmt.Println(matrix)
	return matrix
}

// Координаты вершины
type point struct {
	x, y int
}

// Смещение влево, вправо, вверх, вниз для определения соседей вершины
var directions = [4]point{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

// Метод чтобы увеличить значения координаты смещением на одну позицию для координаты вершины
func (p point) add(direction point) point {
	return point{p.x + direction.x, p.y + direction.y}
}

// Определение допустимых границ и условий для смещения/поиска соседей
func (next point) noAccess(steps [][]int, maze [][]string, start point) bool {
	// координата соседней вершины - проверка на соот-ие границ
	if next.x < 0 || next.x >= len(maze) {
		return true
	}
	if next.y < 0 || next.y >= len(maze[0]) {
		return true
	}

	// обход/пропуск вершины основной матрицы, если она содержит # и самих роботов на основе координат соседней вершины
	if maze[next.x][next.y] == "#" || maze[next.x][next.y] == "B" || maze[next.x][next.y] == "A" {
		return true
	}
	// Если промежуточная матрица не заполнена 0, что гарантирует избежание повторения одних и тех же точек. Аналогично VISITED.
	if steps[next.x][next.y] != 0 {
		return true
	}
	// Если коор-ы соседа совпадают со коор-ми робота из осн-й матрицы
	if next.x == start.x && next.y == start.y {
		return true
	}
	return false
}

// Сборка матрицы с путями роботов
func changeMatrix(maze [][]string, steps [][]int, robot string, start, end point) [][]string {
	//Look up from the lower right corner, if it is less than 1, it is a path node
	var cur = end
	var st = start
	var next point

	stopFor := false

	for cur != st {
		//Начинаем искать с конца до робота
		if steps[cur.x][cur.y] == 1 {
			for _, direction := range directions {
				next = cur.add(direction)
				if next == st {
					maze[cur.x][cur.y] = robot
					stopFor = true
					break
				}
			}
		}

		if stopFor {
			break
		}

		for _, direction := range directions {
			next = cur.add(direction)
			// Отнимаем 1, чтобы найти "соседнюю" точку, которая находится ближе к началу пути И // значение по коор-м next = зн-ю по тек-м коор-м т.е.
			if next.x >= 0 && next.x < len(maze) && next.y >= 0 && next.y < len(maze[0]) && steps[next.x][next.y] == steps[cur.x][cur.y]-1 && maze[next.x][next.y] != "#" {

				// Заменяем на символ робота в ориг-й матрице
				maze[cur.x][cur.y] = robot

				// Текущей коор-й становится next
				cur = next
				// ОЧЕНЬ ВАЖНОЕ УСЛОВИЕ - КОГДА ЦИКЛ НЕ НАЧИНАЕТСЯ С НУЛЯ ПРИ НОВОЙ ИТЕРАЦИИ - for _, direction := range directions
				break
			}
		}
	}
	//fmt.Println("end of CHANGE FUNC")
	return maze
}

func run(maze [][]string, start, end point) [][]int {
	// Иниц-я промеж-й матрицы с нулями на базе основной
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	// Очередь и отправка коор-т стартовой вершины
	Q := []point{start}

	// Пока очередь не опустеет
	for len(Q) > 0 {
		// Работаем только с первым элементом из среза очереди
		cur := Q[0]
		//fmt.Println("cur := Q[0]:", cur)
		// И сразу обрезаем стартовый элемент из среза очереди, чтобы всегда работать с другим первым элементом
		Q = Q[1:]
		//fmt.Println("Q = Q[1:]", Q)

		// !!! continue нужен чтобы не было лишних нулей на местах точек, иначе (break) строить обратный маршрут будет в бесконечном цикле
		if cur == end {
			continue
		}

		// Обход координат соседних вершин
		for _, direction := range directions {
			// коорд-тА соседА текущей точки с использ-ем смещения
			// за раз обход только на одно смещение, а таких будет 4 в методе add
			next := cur.add(direction)
			//fmt.Println("координата next:", next)
			/*
				TRUE - переходит к следующей итерации цикла, независимо от какого-либо условия.
				Это может быть полезно, если нужно пропустить выполнение остальной части текущей итерации и перейти к следующей
				FALSE - этот код фактически не делает ничего.
				Он игнорируется компилятором, потому что код внутри блока if никогда не будет выполнен из-за условия false.
			*/
			// WHY continue?
			if next.noAccess(steps, maze, start) {
				//steps[next.x][next.y] = 0
				continue
			}
			// Если коор-А проходит условия, то присвоить по ней зн-е +1 в промеж-й матрице на базе тек-х коор-т, где везде -1.
			steps[next.x][next.y] = steps[cur.x][cur.y] + 1

			// Помещаем эту NEXT координату в очередь т.к. она соседняя с CUR
			Q = append(Q, next)
			//fmt.Println("Q = append(Q, next)", Q)
		}
	}
	return steps
}
