package main

import (
	"bufio"
	"fmt"
	"os"
)

// ********* Advent of Code 2024 *********
// --- Day 4: Ceres Search --- Puzzle 1
// https://adventofcode.com/2024/day/4

func main() {

	data := GetMatrix("data2.txt")
	pattern := "XMAS"

	answer := day4_1(data, pattern)
	fmt.Println(answer)
}

func day4_1(matrix [][]string, pattern string) int {

	var result int
	var allWays []string // все строки, столбцы, прямые и обратные диагонали из матрицы в виде строк

	// добавить к набору все строки из матрицы
	// allWays = append(allWays, GetLines(matrix)...)

	// добавить к набору все столбцы из матрицы
	// allWays = append(allWays, GetColumns(matrix)...)

	// добавить к набору все прямые диагонали из матрица
	// allWays = append(allWays, GetForwardDiag(matrix)...)

	// добавить к набору все обратные диагонали из матрицы
	allWays = append(allWays, GetBackwardDiag(matrix)...)

	// искать в каждом элемента набора pattern и перевернутый pattern
	// strings.Count(elem, "XMAS"))
	// strings.Count(elem, "SAMX"))

	fmt.Println(allWays)

	return result
}

// Функция извлекает из матрицы все строки
func GetLines(matrix [][]string) []string {
	var result []string

	sizeX := len(matrix[0]) // размерность матрицы по оси X (количество столбцов)
	sizeY := len(matrix)    // размерность матрицы по оси Y (количество строк)

	for y := 0; y < sizeY; y++ {
		var line string
		for x := 0; x < sizeX; x++ {
			line += matrix[y][x]
		}
		result = append(result, line)
	}

	return result
}

// Функция извлекает из матрицы все столбцы
func GetColumns(matrix [][]string) []string {
	var result []string

	sizeX := len(matrix[0]) // размерность матрицы по оси X (количество столбцов)
	sizeY := len(matrix)    // размерность матрицы по оси Y (количество строк)

	for x := 0; x < sizeX; x++ {
		var column string
		for y := 0; y < sizeY; y++ {
			column += matrix[y][x]
		}
		result = append(result, column)
	}

	return result
}

// Функция извлекает из матрицы все прямые диагонали /
func GetForwardDiag(matrix [][]string) []string {
	var result []string

	// последовательность сбора диагоналей
	// 		0 1 2 3
	// 		1 2 3 4
	// 		2 3 4 5
	// 		3 4 5 6
	// 		4 5 6 7
	// 		5 6 7 8

	sizeX := len(matrix[0]) // размерность матрицы по оси X (количество столбцов)
	sizeY := len(matrix)    // размерность матрицы по оси Y (количество строк)

	// перебор диагоналей начинающихся от первого столбца (0-5)
	for line := 0; line < sizeY; line++ {
		var diagonal string

		x := 0
		y := line

		for {
			diagonal += matrix[y][x]
			y--
			x++

			if (x == sizeX) || (y == -1) { //выход за границы матрицы
				break
			}
		}

		result = append(result, diagonal)
	}

	// перебор диагоналей начинающихся от нижней строки (6-8)
	for column := 1; column < sizeX; column++ {
		var diagonal string

		x := column
		y := sizeY - 1

		for {
			diagonal += matrix[y][x]
			y--
			x++

			if (x == sizeX) || (y == -1) { //выход за границы матрицы
				break
			}
		}

		result = append(result, diagonal)
	}

	return result
}

// Функция извлекает из матрицы все обратные диагонали \
func GetBackwardDiag(matrix [][]string) []string {
	var result []string

	// последовательность сбора диагоналей
	// 		3 2 1 0
	// 		4 3 2 1
	// 		5 4 3 2
	// 		6 5 4 3
	// 		7 6 5 4
	// 		8 7 6 5

	sizeX := len(matrix[0]) // размерность матрицы по оси X (количество столбцов)
	sizeY := len(matrix)    // размерность матрицы по оси Y (количество строк)

	// перебор диагоналей начинающихся от последнего столбца (0-5)
	for line := 0; line < sizeY; line++ {
		var diagonal string

		x := sizeX - 1
		y := line

		for {
			diagonal = matrix[y][x] + diagonal // диагональ заполняется справа налево
			y--
			x--

			if (x == -1) || (y == -1) { //выход за границы матрицы
				break
			}
		}

		result = append(result, diagonal)
	}

	// перебор диагоналей начинающихся от нижней строки (6-8)
	for column := sizeX - 2; column >= 0; column-- {
		var diagonal string

		x := column
		y := sizeY - 1

		for {
			diagonal = matrix[y][x] + diagonal // диагональ заполняется справа налево
			y--
			x--

			if (x == -1) || (y == -1) { //выход за границы матрицы
				break
			}
		}

		result = append(result, diagonal)
	}

	return result
}

// Функция разбирает текстовый файл в двумерную матрицу символов
func GetMatrix(filename string) [][]string {

	var result [][]string

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		var element []string

		for _, v := range scanner.Text() {
			element = append(element, string(v))
		}

		result = append(result, element)
	}

	file.Close()

	return result
}
