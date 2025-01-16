package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
)

// ********* Advent of Code 2024 *********
// --- Day 3: Mull It Over --- Puzzle 1
// https://adventofcode.com/2024/day/3

func main() {

	data := GetDataFromFile("data.txt")

	answer := day3_1(data)
	fmt.Println(answer)
}

func day3_1(data []string) int {

	var result int
	var pairs [][]int

	// собираем все пары чисел из каждой строки в один слайс
	for _, v := range data {
		pairs = slices.Concat(pairs, GetPairs(v))
	}

	fmt.Println(pairs) // DEBUG PRINT

	// вычисляем сумму произведений каждой пары
	for _, v := range pairs {
		result += v[0] * v[1]
	}

	return result
}

func GetPairs(line string) [][]int {

	// [[1 4] [34 56] [432 490]]
	var result [][]int

	pattern := regexp.MustCompile(`123`)       // TODO
	matches := pattern.FindAllString(line, -1) // -1 означает "искать все вхождения"

	return result
}

// Функция разбирает текстовый файл с исходными данными в слайс строк
func GetDataFromFile(filename string) []string {

	var result []string

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	file.Close()

	return result
}
