package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// ********* Advent of Code 2024 *********
// --- Day 2: Red-Nosed Reports --- Puzzle 1
// https://adventofcode.com/2024/day/2

func main() {

	reports := GetDataFromFile("data.txt")

	answer := day2_1(reports)
	fmt.Println(answer)

}

func day2_1(list []string) int {

	safeCount := 0

	for _, v := range list {
		if IsLineSafe(v) {
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

// Функция проверяет является ли переданная строка безопасной с точки зрения формулировки задачи
func IsLineSafe(line string) bool {

	var sliceOfStrings []string
	var sliceOfInt []int
	var isIncreasing bool
	var isDecreasing bool
	var result bool

	// преобразование слайса строк в слайс целых чисел -- "7 6 4 2 1" -> [7 6 4 2 1]
	sliceOfStrings = strings.Split(line, " ")

	for _, v := range sliceOfStrings {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
		}
		sliceOfInt = append(sliceOfInt, num)
	}

	// проверка на безопасность последовательности по описанным в задании правилам
	result = true

	for i := range sliceOfInt {

		//если это последний элемент слайса - прервать цикл
		if i == len(sliceOfInt)-1 {
			break
		}

		current := sliceOfInt[i]
		next := sliceOfInt[i+1]
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
		} else {
			if current > next && isDecreasing == true {
				result = true
			} else if current < next && isIncreasing == true {
				result = true
			} else {
				result = false
				break
			}
		}

		fmt.Printf("%v and %v = %v \n", sliceOfInt[i], sliceOfInt[i+1], result)
	}

	fmt.Println("**********************")
	return result
}
