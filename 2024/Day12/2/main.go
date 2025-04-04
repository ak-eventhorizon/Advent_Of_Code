package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

// ********* Advent of Code 2024 *********
// --- Day 12: Garden Groups --- Puzzle 2
// https://adventofcode.com/2024/day/12

const INPUT_FILE_PATH string = "data.txt"
const OUTPUT_FILE_PATH string = "output.txt"

func main() {
	start := time.Now()

	input := GetData(INPUT_FILE_PATH)
	answer := day12_2(input)
	fmt.Println(answer)

	fmt.Printf("%s \n", time.Since(start)) // измерение времени выполнения функции
}

func day12_2(input [][]string) (result int) {

	os.Truncate(OUTPUT_FILE_PATH, 0) //очистка файла

	gardens := extractRegions(input)

	for _, garden := range gardens {

		for i := 0; i < len(garden); i++ {
			line := strings.Join(garden[i], "")
			SaveToFile(line, OUTPUT_FILE_PATH)
		}

		price := calcFencePrice(garden)
		result += price

		SaveToFile("Price: "+strconv.Itoa(price)+"\n", OUTPUT_FILE_PATH)
	}

	return result
}

// Функция вычисляет стоимость забора для полученной матрицы
// Стоимость = количество_клеток_региона * количество_сторон_региона
// +-------+
// |A A A A|	6 * 6 = 36
// |A+-----+
// |A|
// +-+
func calcFencePrice(matrix [][]string) (result int) {

	area := 0  // количество клеток, занятых регионом
	sides := 0 // количество сторон региона

	for _, line := range matrix {
		for _, char := range line {
			if char != "." {
				area++
			}
		}
	}

	sides = countCorners(matrix)

	return area * sides
}

// Функция считает количество углов в матрице, количество сторон фигуры всегда равно количеству углов в ней
func countCorners(matrix [][]string) (result int) {

	// Дополнить исходную матрицу пустыми строкой/столбцом с каждой стороны чтобы при анализе клеток не было необходимости проверять выход за границы матрицы
	var copyMatrix [][]string
	sizeCopyMatrixY := len(matrix) + 2
	sizeCopyMatrixX := len(matrix[0]) + 2

	for i := 0; i < sizeCopyMatrixY; i++ {
		var line []string

		if i == 0 || i == sizeCopyMatrixY-1 { // первая и последняя строка матрицы заполняется пустыми строками
			for range sizeCopyMatrixX {
				line = append(line, ".")
			}
			copyMatrix = append(copyMatrix, line)
		} else {
			line = append(line, ".")
			line = append(line, matrix[i-1]...)
			line = append(line, ".")
			copyMatrix = append(copyMatrix, line)
		}
	}

	// Последовательный анализ каждой клетки (XX)- на предмет касается ли она угла (изнутри и/или снаружи)
	// UL UU UR
	// LL XX RR
	// DL DD DR
	for y, line := range copyMatrix {
		for x, cell := range line {

			XX := cell

			if XX == "." { // пропускаем пустые клетки
				continue
			} else {
				UL := copyMatrix[y-1][x-1]
				UU := copyMatrix[y-1][x]
				UR := copyMatrix[y-1][x+1]
				RR := copyMatrix[y][x+1]
				DR := copyMatrix[y+1][x+1]
				DD := copyMatrix[y+1][x]
				DL := copyMatrix[y+1][x-1]
				LL := copyMatrix[y][x-1]

				// * . *
				// . x *
				// * * *
				if UU == LL && XX != UU {
					result++
				}

				// . x *
				// x x *
				// * * *
				if XX == UU && XX == LL && XX != UL {
					result++
				}

				// * . *
				// * x .
				// * * *
				if UU == RR && XX != UU {
					result++
				}

				// * x .
				// * x x
				// * * *
				if XX == UU && XX == RR && XX != UR {
					result++
				}

				// * * *
				// * x .
				// * . *
				if DD == RR && XX != DD {
					result++
				}

				// * * *
				// * x x
				// * x .
				if XX == RR && XX == DD && XX != DR {
					result++
				}

				// * * *
				// . x *
				// * . *
				if LL == DD && XX != LL {
					result++
				}

				// * * *
				// x x *
				// . x *
				if XX == DD && XX == LL && XX != DL {
					result++
				}
			}
		}
	}

	return result
}

// Функция извлекает обособленные регионы из исходной матрицы
//
// .....    ...     result[0]        result[1]        result[2]
// OOOOO    -->      OOOOO            .....            .....
// OXXOO    -->      O..OO            .XX..            .....
// OOOOO    -->      OOOOO            .....            .....
// OOOXO    -->      OOO.O            .....            ...X.
// OOOOO    -->      OOOOO            .....            .....
func extractRegions(matrix [][]string) (result map[int][][]string) {

	result = map[int][][]string{}
	counter := 0

	lenY := len(matrix)
	lenX := len(matrix[0])

	//делаем копию исходной матрицы, чтобы не вносить в нее изменения; далее работаем с копией
	var copyMatrix [][]string
	for _, line := range matrix {
		copyMatrix = append(copyMatrix, slices.Clone(line))
	}

	//анализ каждой клетки исходного поля
	for y, line := range copyMatrix {
		for x, char := range line {

			// пропускаем клетку, если она пустая
			if char == "." {
				continue
			}

			// создаем пустую матрицу (заполненную ".") по размеру исходной, на которую будем наносить каждый регион и сохранять его отдельно
			var blankMatrix [][]string
			for range lenY {
				line := []string{}
				for range lenX {
					line = append(line, ".")
				}
				blankMatrix = append(blankMatrix, line)
			}

			// поиском в ширину вычиляются все координаты региона текущей клетки
			start := []int{x, y}
			currentCells := [][]int{}
			currentCells = append(currentCells, start)
			regionCoords := [][]int{}
			regionCoords = append(regionCoords, start)

			// пока следующий шаг вычислений находит соседей - продолжаем вычисление
			for len(currentCells) > 0 {
				currentCells = stepBFS(currentCells, copyMatrix)
				regionCoords = append(regionCoords, currentCells...)
			}

			for _, coord := range regionCoords {
				x := coord[0]
				y := coord[1]
				blankMatrix[y][x] = char // сохраняем регион в чистую матрицу
				copyMatrix[y][x] = "."   // в исходной поле заменяем все клетки найденного региона на "." чтобы больше их не анализировать
			}

			// добавление региона к результату
			result[counter] = blankMatrix
			counter++
		}
	}

	return result
}

// Функция получает набор координат и возвращает набор координат всех соседних клеток (по горизонтали и вертикали), имеющих идентичнео содержимое
// Реализует один шаг алгоритма поиска в ширину (Breadth-First Search)
func stepBFS(currentCells [][]int, matrix [][]string) (nextCells [][]int) {

	for _, cell := range currentCells {
		x := cell[0]
		y := cell[1]
		nextCells = append(nextCells, calcNextStep(x, y, matrix)...)
	}

	// удалению дублей, поскольку одна клетка может быть достигнута на этом шаге несколькими разными путями
	nextCells = removeDuplicates(nextCells)

	return nextCells
}

// На основе координат исходной клетки x,y вычисляются коордираты ее соседей с таким же значением по вертикали и горизонтали
func calcNextStep(x, y int, matrix [][]string) (result [][]int) {

	regionName := matrix[y][x] // сохраняем содержимое клетки (имя региона)
	matrix[y][x] = "."         // очищаем клетку, как обработанную
	lenY := len(matrix)
	lenX := len(matrix[0])

	// проверка соседа сверху
	if y-1 >= 0 && matrix[y-1][x] == regionName {
		result = append(result, []int{x, y - 1})
	}
	// проверка соседа снизу
	if y+1 < lenY && matrix[y+1][x] == regionName {
		result = append(result, []int{x, y + 1})
	}
	// проверка соседа слева
	if x-1 >= 0 && matrix[y][x-1] == regionName {
		result = append(result, []int{x - 1, y})
	}
	// проверка соседа справа
	if x+1 < lenX && matrix[y][x+1] == regionName {
		result = append(result, []int{x + 1, y})
	}

	return result
}

// Функция удаляет дубли из произвольного слайса слайсов целых чисел
func removeDuplicates(source [][]int) (result [][]int) {

	// превращение слайса слайсов интов в слайс строк, который можно отсортировать и удалить дубли в сортированном
	strSlice := []string{}
	for _, elem := range source {
		strElem := strings.Trim(strings.Replace(fmt.Sprint(elem), " ", "-", -1), "[]") // [3 4] --> "3-4"
		strSlice = append(strSlice, strElem)
	}

	slices.Sort(strSlice)               // сортировка
	strSlice = slices.Compact(strSlice) // удаление дубликатов

	// превращение слайса строк обратно в слайс слайсов интов
	for _, s := range strSlice {
		elements := strings.Split(s, "-")
		element := []int{}
		for _, v := range elements {
			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			element = append(element, num)
		}
		result = append(result, element)
	}

	return result
}

// Функция извлекает из файла filename матрицу исходных данных
func GetData(filename string) (matrix [][]string) {

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var line []string
		for _, char := range scanner.Text() {
			line = append(line, string(char))
		}
		matrix = append(matrix, line)
	}

	file.Close()

	return matrix
}

// Функция дописывает строку line в файл filename
func SaveToFile(line string, filename string) {

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err2 := file.WriteString(line + "\n")

	if err2 != nil {
		panic(err2)
	}
}
