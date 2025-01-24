package main

import (
	"bufio"
	"fmt"
	"os"
)

// ********* Advent of Code 2024 *********
// --- Day 6: Guard Gallivant --- Puzzle 1
// https://adventofcode.com/2024/day/6

// Объект, который перемещается по полю
type Unit struct {
	x          int        // координата по оси X
	y          int        // координата по оси Y
	onTheField bool       // признак того, что объект находится на поле и не вышел за его пределы
	direction  string     // направление следующего шага "^" | "v" | ">" | "<"
	field      [][]string // поле, на котором располагается объект и препятствия
}

// Сделать один шаг объекта по полю
func (u *Unit) move() {

	var nextX, nextY int // координаты следующей позиции юнита

	fieldSizeX := len(u.field[0])
	fieldSizeY := len(u.field)

	switch u.direction {
	case "^":
		nextX = u.x
		nextY = u.y - 1
	case "v":
		nextX = u.x
		nextY = u.y + 1
	case "<":
		nextX = u.x - 1
		nextY = u.y
	case ">":
		nextX = u.x + 1
		nextY = u.y
	default:
		nextX = u.x
		nextY = u.y
	}

	// если шаг сделан за пределы поля
	if ((nextX < 0) || (nextX >= fieldSizeX)) || ((nextY < 0) || (nextY >= fieldSizeY)) {
		u.x = nextX
		u.y = nextY
		u.onTheField = false
		return
	}

	// если шаг сделан в препятствие - изменить направление следующего шага на 90 градусов вправо
	if u.field[nextY][nextX] == "#" {
		switch u.direction {
		case "^":
			u.direction = ">"
		case ">":
			u.direction = "v"
		case "v":
			u.direction = "<"
		case "<":
			u.direction = "^"
		}
		return
	}

	// сделать следующий шаг
	if u.onTheField {
		u.x = nextX
		u.y = nextY
		u.field[u.y][u.x] = u.direction
		return
	}
}

func (u Unit) status() {
	fmt.Println("Pos:", u.x, u.y, "Dir:", u.direction)
}

func main() {

	startLayout := GetData("data.txt")

	answer := day6_1(startLayout)
	fmt.Println(answer)
}

func day6_1(layout [][]string) (result int) {

	// ограничение на количество ходов, для исключения бесконечного хождения по кругу внутри поля
	// представляется достаточным сделать ограничение количества равным площади поля
	moveLimit := len(layout[0]) * len(layout)

	// находжение на поле стражника и инициализация соответствующего объекта
	guardian := new(Unit)

	for y, line := range layout {
		for x, char := range line {
			if char == "^" || char == ">" || char == "v" || char == "<" {
				guardian.x = x
				guardian.y = y
				guardian.direction = char
				guardian.onTheField = true
				guardian.field = layout
				break // прерывание цикла поиска - подразумеваем, что стражник на поле один
			}
		}
	}

	if guardian.onTheField == false {
		panic("No units on the field")
	}

	// запуск движения по полю
	for i := 0; i < moveLimit; i++ {
		if guardian.onTheField == false {
			break // прервать движение, если объект вышел за пределы поля
		}
		// guardian.status() // отчет о каждом шаге для отладки
		guardian.move()
	}

	result = CalcVisited(layout)
	SaveData(layout, "output.txt")

	return result
}

// Функция на основании состояния поля layout вычисляет количество посещенных клеток.
func CalcVisited(layout [][]string) (visited int) {

	for _, line := range layout {
		for _, char := range line {
			if char == "^" || char == ">" || char == "v" || char == "<" {
				visited++
			}
		}
	}

	return visited
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
