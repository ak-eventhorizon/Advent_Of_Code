package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

// ********* Advent of Code 2024 *********
// --- Day 8: Resonant Collinearity --- Puzzle 1
// https://adventofcode.com/2024/day/8

// Поле, с расстановкой объектов на нем
type Field struct {
	state [][]string // карта поля, на котором располагается объекты
}

func main() {

	startLayout := GetData("data.txt")
	answer := day8_1(startLayout)
	fmt.Println(answer)
}

func day8_1(layout [][]string) (result int) {

	puzzle := new(Field)
	puzzle.state = slices.Clone(layout)

	fmt.Println(puzzle.state)

	return result
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
