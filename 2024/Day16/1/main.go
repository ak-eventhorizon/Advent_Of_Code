package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// ********* Advent of Code 2024 *********
// --- Day 16: Reindeer Maze --- Puzzle 1
// https://adventofcode.com/2024/day/16

const INPUT_FILE_PATH string = "data.txt"
const OUTPUT_FILE_PATH string = "output.txt"

// Структура представляет лабиринт (двумерную матрицу)
// sX, sY - координаты точки старта
// eX, eY - координаты точки выходв
// layout - содержимое матрицы: # - стена, . - пустая клетка, S - старт, E - выход
type Field struct {
	sX, sY int
	eX, eY int
	layout [][]string
}

func (f *Field) SaveToFile(filename string) {

	os.Truncate(filename, 0) // очистка файла

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	for i, line := range f.layout {
		s := strings.Join(line, "")

		if i == 0 {
			SaveLineToFile(s, OUTPUT_FILE_PATH, false)
		} else {
			SaveLineToFile(s, OUTPUT_FILE_PATH, true)
		}
	}
}

func main() {
	start := time.Now()

	field := GetData(INPUT_FILE_PATH)
	answer := day16_1(field)
	fmt.Println(answer)

	fmt.Printf("%s \n", time.Since(start)) // время выполнения функции
}

func day16_1(field Field) (result int) {
	//TODO
	field.SaveToFile(OUTPUT_FILE_PATH)
	return result
}

// Функция извлекает из файла filename набор исходных данных
func GetData(filename string) (field Field) {

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		field.layout = append(field.layout, line)
	}

	file.Close()

	// Поиск координат старта и выхода
	for y, line := range field.layout {
		for x, char := range line {
			switch char {
			case "S":
				field.sX = x
				field.sY = y
			case "E":
				field.eX = x
				field.eY = y
			}
		}
	}

	return field
}

// Функция дописывает строку line в файл filename (признак isNewLine определяет - начинать ли ее с сновой строки)
func SaveLineToFile(line string, filename string, isNewLine bool) {

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	if isNewLine {
		_, err = file.WriteString("\n" + line)
	} else {
		_, err = file.WriteString(line)
	}

	if err != nil {
		panic(err)
	}
}
