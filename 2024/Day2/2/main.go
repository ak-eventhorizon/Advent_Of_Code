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

	// преобразование всех строк в слайсы целых чисел -- "7 6 4 2 1" -> [7 6 4 2 1]
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
		if IsSafe(v) {
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

// Функция проверяет является ли переданный слайс целых чисел безопасным
func IsSafe(input []int) bool {

	var isIncreasing bool
	var isDecreasing bool
	var result bool

	// проверка на безопасность последовательности по описанным в задании правилам
	for i := range input {

		// если это последний элемент слайса - прервать цикл
		if i == len(input)-1 {
			break
		}

		current := input[i]
		next := input[i+1]
		diff := math.Abs(float64(current) - float64(next))

		// разница текущего и следующего элемента
		if diff == 0 || diff > 3 {
			result = false
			break
		}

		// определение направления изменения значений при анализе первого элемента
		if i == 0 {
			if current > next {
				isDecreasing = true
			} else if current < next {
				isIncreasing = true
			} else {
				result = false
				break
			}
		}

		if current > next && isDecreasing {
			result = true
		} else if current < next && isIncreasing {
			result = true
		} else {
			result = false
			break
		}

		// fmt.Printf("%v and %v -> %v \n", input[i], input[i+1], result)
	}

	return result
}

// TODO
// Функция проверяет является ли переданный слайс целых чисел безопасным
// количество допустимых элементов для удаления указывается в параметре tolerance
func isSafeWithException(input []int, tolerance int) bool {

	var result bool

	test := []int{1, 2, 3, 4, 5, 6}
	temp := slices.Delete(test, 0, 1) // удаление элемента из слайса
	fmt.Println(temp)

	return result
}
