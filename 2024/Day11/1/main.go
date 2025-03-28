package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ********* Advent of Code 2024 *********
// --- Day 11: Plutonian Pebbles --- Puzzle 1
// https://adventofcode.com/2024/day/11

func main() {

	input := GetData("data.txt")
	answer := day11_1(input)

	fmt.Println(answer)
}

func day11_1(input []string) (result int) {

	return result
}

// Функция извлекает из файла filename исходняе данные для задачи
func GetData(filename string) (result []string) {

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		line := strings.Split(scanner.Text(), " ")
		result = append(result, line...)
	}

	file.Close()

	return result
}

// Функция сохраняет полученную строку input в файл filename
func SaveData(input []string, filename string) {

	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	line := strings.Join(input, " ")
	file.WriteString(line)

	file.Close()
}
