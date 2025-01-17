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

	data := GetMatrix("data.txt")
	pattern := "XMAS"

	answer := day4_1(data, pattern)
	fmt.Println(answer)
}

func day4_1(matrix [][]string, pattern string) int {

	var result int

	// TODO

	// добавить к набору все строки из матрицы
	// set append GetLines(matrix)

	// добавить к набору все столбцы из матрицы
	// set append GetColumns(matrix)

	// добавить к набору все прямые диагонали из матрица
	// set append GetForwardDiag(matrix)

	// добавить к набору все обратные диагонали из матрицы
	// set append GetBackwardDiag(matrix)

	// искать в каждом элемента набора pattern и перевернутый pattern
	// strings.Count(elem, "XMAS"))
	// strings.Count(elem, "SAMX"))

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
