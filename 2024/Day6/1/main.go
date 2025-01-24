package main

import (
	"bufio"
	"fmt"
	"os"
)

// ********* Advent of Code 2024 *********
// --- Day 6: Guard Gallivant --- Puzzle 1
// https://adventofcode.com/2024/day/6

func main() {

	startLayout := GetData("data.txt")

	answer := day6_1(startLayout)
	fmt.Println(answer)
}

func day6_1(layout [][]string) (result int) {

	result, resultLayout := CalculateRoute(layout)
	SaveData(resultLayout, "output.txt")

	return result
}

// Функция на основании начального состояния layout вычисляет маршрут движения по условиям задачи.
// Возвращает длину пройденного пути pathLenght и финальное состояние поля после вычисления.
func CalculateRoute(layout [][]string) (pathLenght int, finalLayout [][]string) {

	// TODO

	return pathLenght, finalLayout
}

// Функция извлекает из текстового файла карту задачи.
func GetData(filename string) (layout [][]string) {

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
		layout = append(layout, line)
	}

	file.Close()

	return layout
}

// Функция сохраняет полученную карту layout в файл filename.
func SaveData(layout [][]string, filename string) {

	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	for i, line := range layout {
		var s string
		for _, char := range line {
			s += char
		}
		file.WriteString(s)

		if i != len(layout)-1 {
			file.WriteString("\n")
		}
	}

	file.Close()
}
