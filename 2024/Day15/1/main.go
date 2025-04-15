package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// ********* Advent of Code 2024 *********
// --- Day 15: Warehouse Woes --- Puzzle 1
// https://adventofcode.com/2024/day/15

const INPUT_FILE_PATH string = "data.txt"
const OUTPUT_FILE_PATH string = "output.txt"

func main() {
	start := time.Now()

	matrix, moveSet := GetData(INPUT_FILE_PATH)
	answer := day15_1(matrix, moveSet)
	fmt.Println(answer)

	fmt.Printf("%s \n", time.Since(start)) // время выполнения функции
}

func day15_1(matrix [][]string, moveSet []string) (result int) {

	for _, v := range matrix {
		fmt.Println(v)
	}

	fmt.Println(moveSet)

	return result
}

// Функция извлекает из файла filename набор исходных данных
func GetData(filename string) (matrix [][]string, moveSet []string) {

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var isMatrixFinished bool // признак того, что матрицу с игровым полем полностью извлекли из файла и начинается вторая часть файла

	for scanner.Scan() {

		if scanner.Text() == "" {
			isMatrixFinished = true
			continue
		}

		if !isMatrixFinished {
			line := strings.Split(scanner.Text(), "")
			matrix = append(matrix, line)
		} else {
			line := strings.Split(scanner.Text(), "")
			moveSet = append(moveSet, line...)
		}
	}

	file.Close()

	return matrix, moveSet
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
