package main

import (
	"bufio"
	"fmt"
	"os"
)

// ********* Advent of Code 2024 *********
// --- Day 6: Guard Gallivant --- Puzzle 2
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
	if u.field[nextY][nextX] == "#" || u.field[nextY][nextX] == "O" {
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

// Вывод статуса юнита
func (u Unit) status() {
	fmt.Println("Pos:", u.x, u.y, "Dir:", u.direction)
}

// Объект - препятствие "O", которое можно разместить на поле
type Obstacle struct {
	x    int        // координата по оси X
	y    int        // кооридната по оси Y
	hits HitCounter // если с одной стороны прошло более 1 столкновения - значит стражник ходит кругами
}

// Cчетчик столкновений со сторонами препятствия
type HitCounter struct {
	top    int
	bottom int
	left   int
	right  int
}

func main() {

	startLayout := GetData("data.txt")

	answer := day6_2(startLayout)
	fmt.Println(answer)
}

func day6_2(layout [][]string) (result int) {

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

// Функция устанавливает препятствие в указанные координаты и проверяет создает ли оно петлю длижения стражника
// true - препятсвие зацикливает стражника, false - препятствие не зацикливает стражника
func TestObstacle(x int, y int) (result bool) {

	// TODO
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
