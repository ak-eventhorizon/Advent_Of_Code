package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"time"
)

// ********* Advent of Code 2024 *********
// --- Day 12: Garden Groups --- Puzzle 1
// https://adventofcode.com/2024/day/12

func main() {

	input := GetData("data.txt")
	answer := day12_1(input)

	fmt.Println(answer)
}

func day12_1(input [][]string) (result int) {
	start := time.Now()

	gardens := extractIdenticalGardens(input)

	os.Truncate("output.txt", 0) //очистка файла
	for _, v := range gardens {  // запись в файл всех матриц с садами каждого вида
		SaveData(v, "output.txt")
	}

	fmt.Printf("%s \n", time.Since(start)) // измерение времени выполнения функции
	return result
}

// Функция разбирает исходную матрицу на структуру нескольких матриц, каждая из которых содержит карту только своего вида
func extractIdenticalGardens(matrix [][]string) (result map[string][][]string) {

	result = map[string][][]string{}

	lenY := len(matrix)
	lenX := len(matrix[0])

	//анализ каждой клетки исходного поля
	for y, line := range matrix {
		for x, char := range line {

			_, isKeyExistInMap := result[char]

			// если такого ключа в мапе еще нет - его нужно создать и заполнить пустой матрицей
			if !isKeyExistInMap {
				// создаем пустую матрицу (заполненную ".") по размеру исходной, на которую будем наносить каждый вид сада и сохранять его отдельно
				var blankMatrix [][]string
				for range lenY {
					line := []string{}
					for range lenX {
						line = append(line, ".")
					}
					blankMatrix = append(blankMatrix, line)
				}
				result[char] = slices.Clone(blankMatrix)
			}

			result[char][y][x] = char
		}
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

// Функция добавляет полученную матрицу matrix в файл filename
func SaveData(matrix [][]string, filename string) {

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	for i, line := range matrix {

		var s string
		for _, char := range line {
			s += char
		}
		file.WriteString(s + "\n")

		if i == len(matrix)-1 {
			file.WriteString("\n")
		}
	}

	file.Close()
}
