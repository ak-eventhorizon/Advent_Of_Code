package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

// ********* Advent of Code 2024 *********
// --- Day 2: Red-Nosed Reports --- Puzzle 2
// https://adventofcode.com/2024/day/2

func main() {

	reports := GetDataFromFile("data.txt")

	answer := day2_2(reports)
	fmt.Println(answer)

}

func day2_2(list []string) int {

	safeCount := 0

	// преобразование всех строк в слайсы целых чисел -- ["7 6 4 2 1" , "1 2 7 8 9"] -> [[7 6 4 2 1] , [1 2 7 8 9]]
	var reportsInt [][]int

	for _, v := range list {

		var sliceOfStrings []string
		var sliceOfInts []int

		sliceOfStrings = strings.Split(v, " ")

		for _, v := range sliceOfStrings {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
			}
			sliceOfInts = append(sliceOfInts, num)
		}

		reportsInt = append(reportsInt, sliceOfInts)
	}

	// проверка каждого элемента на безопасность по условиям задачи
	for _, v := range reportsInt {
		if isSafe(v, true) {
			safeCount++
		}
	}

	return safeCount
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
		line := strings.Trim(scanner.Text(), " ")
		result = append(result, line)
	}

	file.Close()

	return result
}

// Функция проверяет является ли переданный слайс целых чисел безопасным по условиям задачи.
// Параметр tolerance - разрешает игнорирование одного ошибочного элемента в последовательности
func isSafe(input []int, tolerance bool) bool {

	var isIncreasing bool
	var isDecreasing bool

	var results []bool // содержит результаты анализа пар значений из входной последовательности
	var result bool

	// проверка на безопасность последовательности по описанным в задании правилам
	for i := range input {

		condition1 := true
		condition2 := true

		// если это последний элемент слайса - прервать цикл
		if i == len(input)-1 {
			break
		}

		current := input[i]
		next := input[i+1]
		diff := math.Abs(float64(current) - float64(next))

		// разница текущего и следующего элемента
		if diff == 0 || diff > 3 {
			condition1 = false
		}

		// определение направления изменения значений при анализе первого элемента
		if i == 0 {
			if current > next {
				isDecreasing = true
			} else if current < next {
				isIncreasing = true
			} else {
				condition1 = false
			}
		}

		if (current > next && isDecreasing) || (current < next && isIncreasing) {
			condition2 = true
		} else {
			condition2 = false
		}

		if condition1 && condition2 {
			results = append(results, true)
		} else {
			results = append(results, false)
		}
	}

	if slices.Contains(results, false) {

		// ветка отрабатывает если исходная последовательность небезопасная и разрешено право на одну ошибку
		if tolerance {

			// проверка исходной последовательности без каждого элемента поочередно, если хоть одна из комбинаций безопасна - признать исходную комбинацию безопасной
			var midResults []bool
			for i := 0; i < len(input); i++ {
				if isSafe(slices.Delete(slices.Clone(input), i, i+1), false) {
					midResults = append(midResults, true)
				} else {
					midResults = append(midResults, false)
				}
			}

			if slices.Contains(midResults, true) {
				result = true
			} else {
				result = false
			}

		} else {
			result = false
		}
	} else {
		result = true
	}

	return result
}
