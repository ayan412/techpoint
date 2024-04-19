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

		fmt.Println("reflect wHouse:", reflect.TypeOf(allSlices[0]))
		//fmt.Println(len(allSlices[0][0]))
		//fmt.Println(allSlices[0][0])
		//fmt.Println(allSlices[0])

		wHouse := allSlices[0]
		//fmt.Println("wH:", wHouse)

		//Start walking the maze run
		steps := run(wHouse, point{2, 3}, point{len(wHouse) - 1, len(wHouse[0]) - 1})
		fmt.Println(steps)
		//Give a path according to steps
		wHouse = changeMatrix(wHouse, steps)

		fmt.Println("steps:")
		for x := range steps {
			for y := range steps[i] {
				fmt.Printf("%4d", steps[x][y])
			}
			fmt.Println()

			fmt.Println("changed maze:")
			for i := range wHouse {
				for j := range wHouse[i] {
					fmt.Printf("%s", wHouse[i][j])
				}
				fmt.Println()
			}
		}
		//robotA := point{2, 3}
		//robotB := Position{3, 4}

		// pathA := bfs(robotA, point{0, 0}, "a")
		// // := bfs(robotB, Position{len(allSlices[0]) - 1, len(allSlices[0][0]) - 1}, "b")

		// for point, robot := range pathA.route {
		// 	wHouse[point.x][point.y] = robot
		// }

		// // for point, robot := range pathB.route {
		// // 	wHouse[point.x][point.y] = robot
		// // }

		// for _, row := range wHouse {
		// 	fmt.Println("row", row)
		// }

		// shortPath(wHouse, robotA, point{0, 0})
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
	//fmt.Println("len A", sumA)
	// Сумма разницы чисел между МАХ и В
	sumB := 0
	for i := 0; i < len(resultBSlice); i++ {
		sumB += resultBSlice[i]
	}
	//fmt.Println("len B", sumB)
	if sumA > sumB {
		// алгоритм по которому будет дописываться путь робота в точку 0;0

	} else {
		// алгоритм по которому будет дописываться путь робота в точку MAX
	}

	// Выводим разность
	//fmt.Println("Разность чисел между MAX и A:", resultASlice)
	//fmt.Println("Разность чисел между MAX и B:", resultBSlice)

}

// Define the coordinate structure
type point struct {
	x, y int
}

// Define the difference between the top, bottom, left, and right elements of the current element
var directions = [4]point{
	{-1, 0}, 
	{0,-1}, 
	{1, 0}, 
	{0, 1},
}

// увеличить значения координат смещением на одну позицию по гор-ти и вертикали, но как по какому принципу выбор в какое напр-е сначала идти???
func (p point) add(direction point) point {
	return point{p.x + direction.x, p.y + direction.y}
}

// Judging the situation where next cannot go: 1. Encounter obstacles, go out of the boundary 2. Have already gone through, can’t go back 3. Can’t go back in a circle
func (next point) noAccess(steps [][]int, maze [][]string, start point) bool {
	// касается промежуточной матрицы
	if next.x < 0 || next.x >= len(maze) {
		return true
	}
	if next.y < 0 || next.y >= len(maze) {
		return true
	}
	// касается текущей матрицы
	if maze[next.x][next.y] == "#" || maze[next.x][next.y] == "B" || maze[next.x][next.y] == "A" {
		return true
	}
	// ????
	if steps[next.x][next.y] != 0 {
		return true
	}
	//??? ЧТО ЭТО ПРОВЕРЯЕТ ????
	if next.x == start.x && next.y == start.y {
		return true
	}
	return false
}

// здесь в исходном примере из правого нижнего в левый верхний начинается поиск
func changeMatrix(maze [][]string, steps [][]int) [][]string {
	var current = point{len(maze) - 1, len(maze[0]) - 1}
	// CHANGE!!! должна браться координата любого робота !!!
	var start = point{2, 3}
	for current != start {
		//Find the surrounding nodes, whether it is the value of the current node -1
		for _, direction := range directions {
			next := current.add(direction)
			// ПОЧЕМУ вычитаем 1 ????!!!!
			// Judge whether it is legal
			if next.x >= 0 && next.x < len(maze) && next.y >= 0 && next.y < len(maze[0]) && steps[next.x][next.y] == steps[current.x][current.y]-1 {
				//Modify maze to 6
				maze[current.x][current.y] = "b"
				current = next
			}
		}
	}
	maze[2][3] = "b"
	return maze
}

func run(maze [][]string, start, end point) [][]int {
	//Generate steps matrix
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	//Generate a queue and put the first node in the queue
	Q := []point{start}

	//If the queue is not empty, it means there is no end
	for len(Q) > 0 {
		//Remove the head element and delete it
		peek := Q[0]
		fmt.Println("Queue is:", Q)

		Q = Q[1:]

		if peek == end {
			break
		}
		//пытаемся найти соседей
		for _, direction := range directions {
			// коорд-тА соседА текущей точки с использ-ем смещения
			next := peek.add(direction)
			/*
			TRUE - переходит к следующей итерации цикла, независимо от какого-либо условия. 
			Это может быть полезно, если вам нужно пропустить выполнение остальной части текущей итерации и перейти к следующей
			FALSE - этот код фактически не делает ничего. 
			Он игнорируется компилятором, потому что код внутри блока if никогда не будет выполнен из-за условия false.
			*/
			if next.noAccess(steps,maze,start) {
				continue
			}
			//Assign a value to steps
			steps[next.x][next.y] = steps[peek.x][peek.y] + 1
			//Put into the queue
			Q = append(Q, next)
		}
	}
	return steps
}

// type robotPath struct {
// 	success bool
// 	route   map[point]string
// }

// var wHouse [][]string

// func isValidCell(x, y int) bool {
// 	if x < 0 || x >= len(wHouse) || y < 0 || y >= len(wHouse[0]) { // x строки, y столбцы
// 		return false
// 	} else {
// 		if wHouse[x][y] == "#" || wHouse[x][y] == "A" || wHouse[x][y] == "B" {
// 			return false
// 		}
// 	}
// 	//fmt.Println("wHouse", wHouse)
// 	return true

// }

// func shortPath(wHouse [][]string, start, end point) {
// 	h := len(wHouse)
// 	w := len(wHouse[0])
// 	fmt.Println(h, w)
// 	visit := make(map[point]bool) // карта для посещенных точек
// 	queu := []point{start} // очередь из соседних точек

// 	for len(queu) > 0 {
// 		curren := queu[0] // FIFO zero index = first
// 		fmt.Println("(THIS IS queu[0]:", queu[0])
// 		queu = queu[1:] // оставшаяся очередь

// 		if curren == end {
// 			break
// 		}

// 		neighbo := []point{{curren.x - 1, curren.y}, {curren.x + 1, curren.y}, {curren.x, curren.y + 1}, {curren.x, curren.y - 1}}

// 		for _,nei := range neighbo {
// 			if isValidCell(nei.x, nei.y) && !visit[nei] {
// 				visit[nei] = true
// 				queu = append(queu, nei)
// 			}
// 		}
// 	}
// }

// //Функция обхода в ширину (BFS)
// func bfs(start, end point, name string) robotPath {
// 	visited := make(map[point]bool) // карта для посещенных точек
// 	queue := []point{start} // очередь из соседних точек
// 	route := make(map[point]string)

// 	for len(queue) > 0 {
// 		current := queue[0]
// 		queue = queue[1:]

// 		if current == end {
// 			return robotPath{true, route}
// 		}

// 		// Перемещаемся во все соседние клетки
// 		neighbors := []point{{current.x - 1, current.y}, {current.x + 1, current.y}, {current.x, current.y - 1}, {current.x, current.y + 1}}
// 		for _, neighbor := range neighbors {
// 			if isValidCell(neighbor.x, neighbor.y) && !visited[neighbor] {
// 				visited[neighbor] = true
// 				queue = append(queue, neighbor)
// 				route[neighbor] = name
// 			}
// 		}
// 	}
// 	return robotPath{false, nil}
// }
