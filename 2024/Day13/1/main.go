package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// ********* Advent of Code 2024 *********
// --- Day 13: Claw Contraption --- Puzzle 1
// https://adventofcode.com/2024/day/13

const INPUT_FILE_PATH string = "data.txt"
const OUTPUT_FILE_PATH string = "output.txt"

// Структура представляет точку на двумерной плоскости
type Point struct {
	x int
	y int
}

// Структура представляет игровой автомат, имеющий кнопки A и B, клешню Claw и приз Prize
// При нажатии кнопки А: происходит смещение клешни на AX по оси X и на AY по оси Y
// При нажатии кнопки B: происходит смещение клешни на BX по оси X и на BY по оси Y
type Machine struct {
	AX, AY int
	BX, BY int
	Claw   Point
	Prize  Point
}

func (m *Machine) PushButtonA() {
	m.Claw.x += m.AX
	m.Claw.y += m.AY
}

func (m *Machine) PushButtonB() {
	m.Claw.x += m.BX
	m.Claw.y += m.BY
}

func (m Machine) Display() {
	fmt.Printf("A(%v,%v) B(%v,%v) Claw (%v,%v) Prize(%v,%v)\n",
		m.AX,
		m.AY,
		m.BX,
		m.BY,
		m.Claw.x,
		m.Claw.y,
		m.Prize.x,
		m.Prize.y)
}

func main() {
	start := time.Now()

	newClaw := Point{0, 0}
	newPrize := Point{90, 90}
	newMachine := Machine{5, 5, 3, 3, newClaw, newPrize}
	newMachine.Display()

	// input := GetData(INPUT_FILE_PATH)
	// answer := day13_1(input)
	// fmt.Println(answer)

	fmt.Printf("%s \n", time.Since(start)) // время выполнения функции
}

func day13_1(input [][]string) (result int) {

	return result
}

// Функция извлекает из файла filename матрицу исходных данных
func GetData(filename string) (matrix [][]string) {

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
		matrix = append(matrix, line)
	}

	file.Close()

	return matrix
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
