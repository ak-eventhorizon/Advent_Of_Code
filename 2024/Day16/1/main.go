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

	for _, line := range f.layout {
		s := strings.Join(line, "")
		_, err2 := file.WriteString(s + "\n")
		if err2 != nil {
			panic(err2)
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
			if char == "S" {
				field.sX = x
				field.sY = y
			} else if char == "E" {
				field.eX = x
				field.eY = y
			}
		}
	}

	return field
}

// Функция дописывает строку line в файл filename
func SaveToFile(line string, filename string) {

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err2 := file.WriteString(line + "\n")

	if err2 != nil {
		panic(err2)
	}
}
