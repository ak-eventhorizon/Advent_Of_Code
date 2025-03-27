package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// ********* Advent of Code 2024 *********
// --- Day 10: Hoof It --- Puzzle 1
// https://adventofcode.com/2024/day/10

func main() {

	input := GetData("data.txt")
	answer := day10_1(input)

	fmt.Println(answer)
}

func day10_1(fieldMap [][]int) (result int) {

	// поиск на карте всех точек старта "0"
	startPoints := [][]int{}
	for y, line := range fieldMap {
		for x := range line {
			if fieldMap[y][x] == 0 {
				startPoints = append(startPoints, []int{x, y})
			}
		}
	}

	// вычисление для каждой точки количества возможных из нее маршрутов
	for _, startPoint := range startPoints {
		routes := countRoutes(startPoint, fieldMap)
		result += routes
	}

	return result
}

// Функция возвращает количество финальных точек "9", до которых можно добраться по полю field,
// из точки start[xy] по указанным в задаче правилам построения маршрута 0->1->2->3->4->5->6->7->8->9
func countRoutes(start []int, field [][]int) (routes int) {

	currentCells := [][]int{}
	currentCells = append(currentCells, start)

	// для прохождения всего пути по условиям задачи нужно 9 шагов
	for i := 1; i <= 9; i++ {
		currentCells = stepBFS(currentCells, field)
	}

	routes = len(currentCells)

	return routes
}

// Функция реализует один шаг алгоритма поиска в ширину (Breadth-First Search) на поле field
// получает на вход слой текущих клеток (слайс координат [xy]), возвращает слой клеток на которые можно перейти из текущих (слайс координат [xy])
func stepBFS(currentCells [][]int, field [][]int) (nextCells [][]int) {

	for _, cell := range currentCells {
		nextCells = append(nextCells, getNextCells(cell, field)...)
	}

	// удалению дублей, поскольку одна клетка может быть достигнута на этом шаге несколькими разными путями
	nextCells = removeDuplicates(nextCells)

	return nextCells
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

// Функция возвращает координаты клеток [[xy] [xy]], в которые можно сделать шаг из указанной клетки [xy] на поле field
func getNextCells(initCell []int, field [][]int) (nextCells [][]int) {

	fieldSizeX := len(field[0])
	fieldSizeY := len(field)
	cellX := initCell[0]
	cellY := initCell[1]
	cellValue := field[cellY][cellX]

	// пробуем шаг вверх
	nextCellX := cellX
	nextCellY := cellY - 1

	if nextCellY >= 0 {
		nextCellValue := field[nextCellY][nextCellX]
		if nextCellValue == cellValue+1 {
			nextCell := []int{nextCellX, nextCellY}
			nextCells = append(nextCells, nextCell)
		}
	}

	// пробуем шаг вниз
	nextCellX = cellX
	nextCellY = cellY + 1

	if nextCellY < fieldSizeY {
		nextCellValue := field[nextCellY][nextCellX]
		if nextCellValue == cellValue+1 {
			nextCell := []int{nextCellX, nextCellY}
			nextCells = append(nextCells, nextCell)
		}
	}

	// пробуем шаг влево
	nextCellX = cellX - 1
	nextCellY = cellY

	if nextCellX >= 0 {
		nextCellValue := field[nextCellY][nextCellX]
		if nextCellValue == cellValue+1 {
			nextCell := []int{nextCellX, nextCellY}
			nextCells = append(nextCells, nextCell)
		}
	}

	// пробуем шаг вправо
	nextCellX = cellX + 1
	nextCellY = cellY

	if nextCellX < fieldSizeX {
		nextCellValue := field[nextCellY][nextCellX]
		if nextCellValue == cellValue+1 {
			nextCell := []int{nextCellX, nextCellY}
			nextCells = append(nextCells, nextCell)
		}
	}

	return nextCells
}

// Функция извлекает из файла filename карту поля
func GetData(filename string) (layout [][]int) {

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var line []int
		for _, char := range scanner.Text() {

			n, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			line = append(line, n)
		}
		layout = append(layout, line)
	}

	file.Close()

	return layout
}

// Функция сохраняет полученную карту layout в файл filename.
func SaveData(layout [][]int, filename string) {

	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	for i, line := range layout {
		var s string
		for _, char := range line {
			s += strconv.Itoa(char)
		}
		file.WriteString(s)

		if i != len(layout)-1 {
			file.WriteString("\n")
		}
	}

	file.Close()
}
