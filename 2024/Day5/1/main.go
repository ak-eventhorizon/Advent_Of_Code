package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ********* Advent of Code 2024 *********
// --- Day 5: Print Queue --- Puzzle 1
// https://adventofcode.com/2024/day/5

func main() {

	rules, updates := GetData("data.txt")

	fmt.Println(rules)
	fmt.Println(updates)

	answer := day5_1(rules, updates)
	fmt.Println(answer)
}

func day5_1(rules [][]int, updates [][]int) (result int) {

	for _, update := range updates {
		if IsCorrect(update, rules) {
			result += GetMiddleNumber(update)
		}
	}
	return result
}

// Функция проверяет соответствует ли слайс s списку правил rules.
func IsCorrect(s []int, rules [][]int) (result bool) {

	// TODO

	return result
}

// Функция возвращает средний элемент слайса нечетной длины.
// Если слайс четной длины - среднего элемента нет и функция вернет 0.
func GetMiddleNumber(s []int) (middleElement int) {

	if len(s)%2 == 0 {
		middleElement = 0
	} else {
		i := (len(s) - 1) / 2
		middleElement = s[i]
	}

	return middleElement
}

// Функция извлекает из текстового файла набор правил и набор обновлений по условиям задачи.
func GetData(filename string) (rules [][]int, updates [][]int) {

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var isRulesParseFinished bool

	for scanner.Scan() {

		if scanner.Text() == "" {
			isRulesParseFinished = true
			continue
		}

		var rule []int
		var update []int

		if !isRulesParseFinished {
			str := strings.Split(scanner.Text(), "|")
			for _, v := range str {
				val, err := strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
				rule = append(rule, val)
			}
			rules = append(rules, rule)
		} else {
			str := strings.Split(scanner.Text(), ",")
			for _, v := range str {
				val, err := strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
				update = append(update, val)
			}
			updates = append(updates, update)
		}
	}

	file.Close()

	return rules, updates
}
