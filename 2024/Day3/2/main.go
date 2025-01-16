package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// ********* Advent of Code 2024 *********
// --- Day 3: Mull It Over --- Puzzle 2
// https://adventofcode.com/2024/day/3

func main() {

	data := GetDataFromFileOneLine("data.txt")

	answer := day3_2(data)
	fmt.Println(answer)
}

func day3_2(dataLine string) int {

	var result int

	// собираем все пары чисел из строки в один слайс
	pairs := GetPairs(dataLine)

	// вычисляем сумму произведений каждой пары
	for _, v := range pairs {
		result += v[0] * v[1]
	}

	return result
}

// Функция возвращает все наборы пар чисел, извлеченных из строки
func GetPairs(line string) [][]int {

	var result [][]int

	pattern := regexp.MustCompile(`(don't\(\)).*?(do\(\))`) // "don't()*************do()"
	line = pattern.ReplaceAllString(line, "")               // удаление всех вхождений из исходной строки

	pattern = regexp.MustCompile(`(don't\(\)).*`) // "don't()************конец_строки"
	line = pattern.ReplaceAllString(line, "")     // удаление непарного фрагмента, начинающегося с don't()******* если такой есть

	pattern = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`) // "mul(число,число)"
	matches := pattern.FindAllStringSubmatch(line, -1)

	for _, v := range matches {

		var pair []int

		str1, str2 := v[1], v[2]

		num1, err := strconv.Atoi(str1)
		if err != nil {
			fmt.Println(err)
		}
		pair = append(pair, num1)

		num2, err := strconv.Atoi(str2)
		if err != nil {
			fmt.Println(err)
		}
		pair = append(pair, num2)

		result = append(result, pair)
	}

	return result
}

// Функция собирает текстовый файл с исходными данными в одну строку
func GetDataFromFileOneLine(filename string) string {

	var result string

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		result += scanner.Text()
	}

	file.Close()

	return result
}
