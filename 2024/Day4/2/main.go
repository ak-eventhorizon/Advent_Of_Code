package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"unicode/utf8"
)

// ********* Advent of Code 2024 *********
// --- Day 4: Ceres Search --- Puzzle 2
// https://adventofcode.com/2024/day/4

// Подход для решения: перебор всех маленьких матриц в большой.
// Размер маленькой матрицы NxN ,где N - это длина паттерна. В случае с паттерном "MAS" - малая матрица 3x3

//  1 1 1 *		* 2 2 2		* * * *		* * * *		* * * *		* * * *
//  1 1 1 *		* 2 2 2		3 3 3 *		* 4 4 4		* * * *		* * * *
//  1 1 1 *		* 2 2 2		3 3 3 *		* 4 4 4		5 5 5 *		* 6 6 6
//  * * * *		* * * *		3 3 3 *		* 4 4 4		5 5 5 *		* 6 6 6
//  * * * *		* * * *		* * * *		* * * *		5 5 5 *		* 6 6 6

// В каждой маленькой матрице проверить диагонали на предмет наличия в них прямого и обратного паттерна.

func main() {

	data := GetMatrix("data.txt")
	pattern := "MAS"

	answer := day4_2(data, pattern)
	fmt.Println(answer)
}

func day4_2(matrix [][]string, pattern string) int {

	var result int
	subMatrixSize := utf8.RuneCountInString(pattern) // количество символов в строке pattern

	sizeX := len(matrix[0]) // размерность матрицы по оси X (количество столбцов)
	sizeY := len(matrix)    // размерность матрицы по оси Y (количество строк)

	// обход всех ячеек матрицы
	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			subMatrix, err := GetSubMatrixFromMatrix(x, y, subMatrixSize, matrix)

			// пропуск итерации если подматрица вываливается за границы родительской
			if err != nil {
				continue
			}

			fwDiag := GetForwardMainDiag(subMatrix)
			bwDiag := GetBackwardMainDiag(subMatrix)

			if (pattern == fwDiag || ReverseStr(pattern) == fwDiag) &&
				(pattern == bwDiag || ReverseStr(pattern) == bwDiag) {
				result++
			}
		}
	}

	return result
}

// Функция возвращает подматрицу размерностью size, с началом в точке (x,y), содержащуюся в матрице matrix.
// Функция возвращает ошибку, если подматрицу с указанными параметрами выделить нельзя, т.е. она выходит за границы исходной матрицы.
func GetSubMatrixFromMatrix(x int, y int, size int, matrix [][]string) ([][]string, error) {

	// x=2, y=3, size=4

	//          x
	//      * * v * * * *
	//      * * v * * * *
	//      * * v * * * *
	//    y > > S I Z E *
	//      * * I X X X *
	//      * * Z X X X *
	//      * * E X X X *
	//      * * * * * * *

	var result [][]string
	sizeX := len(matrix[0]) // размерность матрицы по оси X (количество столбцов)
	sizeY := len(matrix)    // размерность матрицы по оси Y (количество строк)

	if x+size > sizeX || y+size > sizeY {
		return nil, errors.New("Sub-matrix is out of range of parent matrix")
	}

	for yi := y; yi <= y+size-1; yi++ { // y
		var line []string
		for xj := x; xj <= x+size-1; xj++ { // x
			line = append(line, matrix[yi][xj])
		}
		result = append(result, line)
	}

	return result, nil
}

// Функция извлекает из квадратной матрицы прямую диагональ /
func GetForwardMainDiag(matrix [][]string) string {
	var result string

	sizeX := len(matrix[0]) // размерность матрицы по оси X (количество столбцов)
	sizeY := len(matrix)    // размерность матрицы по оси Y (количество строк)

	x := 0
	y := sizeY - 1

	for {
		result += matrix[y][x]
		y--
		x++

		if (x == sizeX) || (y == -1) { //выход за границы матрицы
			break
		}
	}

	return result
}

// Функция извлекает из квадратной матрицы обратную диагональ \
func GetBackwardMainDiag(matrix [][]string) string {
	var result string

	sizeX := len(matrix[0]) // размерность матрицы по оси X (количество столбцов)
	sizeY := len(matrix)    // размерность матрицы по оси Y (количество строк)

	x := 0
	y := 0

	for {
		result += matrix[y][x]
		y++
		x++

		if (x == sizeX) || (y == sizeY) { //выход за границы матрицы
			break
		}
	}

	return result
}

// Функция возвращает перевернутую строку
func ReverseStr(s string) string {
	var result string

	for _, v := range s {
		result = string(v) + result
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
