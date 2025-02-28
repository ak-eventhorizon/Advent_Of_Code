package main

import (
	"bufio"
	"fmt"
	"os"
)

// ********* Advent of Code 2024 *********
// --- Day 6: Guard Gallivant --- Puzzle 2
// https://adventofcode.com/2024/day/6

// Объект - препятствие "O", которое можно разместить на поле
type Obstacle struct {
	x         int // координата по оси X
	y         int // кооридната по оси Y
	hitTop    int // если с одной стороны прошло более 1 столкновения - значит стражник ходит кругами
	hitRight  int // если с одной стороны прошло более 1 столкновения - значит стражник ходит кругами
	hitBottom int // если с одной стороны прошло более 1 столкновения - значит стражник ходит кругами
	hitLeft   int // если с одной стороны прошло более 1 столкновения - значит стражник ходит кругами
}

// Объект - стражник "^|v|>|<", который перемещается по полю
type Unit struct {
	x          int    // координата по оси X
	y          int    // координата по оси Y
	onTheField bool   // признак того, что объект находится на поле и не вышел за его пределы
	direction  string // направление следующего шага "^" | "v" | ">" | "<"
}

// Объект, представляющий собой игровое поле, его состояние и объекты, которые на нем находятся
type Field struct {
	state    [][]string // состояние игрового поля (матрица символов из файла)
	unit     *Unit      // стражник, который перемещается по полю
	obstacle *Obstacle  // препятствие, добавленное на игровое поле
}

// Сделать шаг объекта по полю
func (f *Field) moveUnit() {

	var nextX, nextY int // координаты следующей позиции юнита

	fieldSizeX := len(f.state[0])
	fieldSizeY := len(f.state)

	switch f.unit.direction {
	case "^":
		nextX = f.unit.x
		nextY = f.unit.y - 1
	case "v":
		nextX = f.unit.x
		nextY = f.unit.y + 1
	case "<":
		nextX = f.unit.x - 1
		nextY = f.unit.y
	case ">":
		nextX = f.unit.x + 1
		nextY = f.unit.y
	default:
		nextX = f.unit.x
		nextY = f.unit.y
	}

	// если шаг сделан за пределы поля
	if ((nextX < 0) || (nextX >= fieldSizeX)) || ((nextY < 0) || (nextY >= fieldSizeY)) {
		f.unit.x = nextX
		f.unit.y = nextY
		f.unit.onTheField = false
		return
	}

	// если шаг сделан в препятствие "#" - изменить направление следующего шага на 90 градусов вправо
	if f.state[nextY][nextX] == "#" {
		switch f.unit.direction {
		case "^":
			f.unit.direction = ">"
		case ">":
			f.unit.direction = "v"
		case "v":
			f.unit.direction = "<"
		case "<":
			f.unit.direction = "^"
		}
		return
	}

	// если шаг сделан в препятствие "O" - изменить направление следующего шага на 90 градусов вправо и увеличить счетчик столкновений с препятствием
	if f.state[nextY][nextX] == "O" {
		switch f.unit.direction {
		case "^":
			f.unit.direction = ">"
			f.obstacle.hitBottom++
		case ">":
			f.unit.direction = "v"
			f.obstacle.hitLeft++
		case "v":
			f.unit.direction = "<"
			f.obstacle.hitTop++
		case "<":
			f.unit.direction = "^"
			f.obstacle.hitRight++
		}
		return
	}

	// сделать следующий шаг
	if f.unit.onTheField {
		f.unit.x = nextX
		f.unit.y = nextY
		f.state[nextY][nextX] = f.unit.direction
		return
	}
}

// Установить на поле препятствие "O" в координаты x, y. Если препятствие было установлено ранее - оно будет заменено.
func (f *Field) setObstacle(x, y int) {

	// Если клетка на поле свободна
	if f.state[y][x] == "." {

		// Убрать ранее установленное препятствие, если оно было
		if f.obstacle != nil {
			f.state[f.obstacle.y][f.obstacle.x] = "."
			f.obstacle = nil
		}

		f.state[y][x] = "O"
		f.obstacle = new(Obstacle)
		f.obstacle.x = x
		f.obstacle.y = y
	}
}

// Вывод статуса юнита
func (f *Field) status() {
	fmt.Println("Unit:", f.unit.x, f.unit.y, "Direction:", f.unit.direction)
	fmt.Println("Obstacle:", f.obstacle.x, f.obstacle.y)
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

	// инициализация нового игрового поля
	field := new(Field)
	field.state = layout

	// находжение на поле стражника и инициализация соответствующего объекта
	guardian := new(Unit)
	for y, line := range layout {
		for x, char := range line {
			if char == "^" || char == ">" || char == "v" || char == "<" {
				guardian.x = x
				guardian.y = y
				guardian.direction = char
				guardian.onTheField = true
				break // прерывание цикла поиска - подразумеваем, что стражник на поле один
			}
		}
	}

	if guardian.onTheField == false {
		panic("No units on the field")
	}

	field.unit = guardian

	// запуск движения по полю
	for range moveLimit {
		if field.unit.onTheField == false {
			break // прервать движение, если объект вышел за пределы поля
		}
		// field.status() // отчет о каждом шаге для отладки
		field.moveUnit()
	}

	result = CalcVisited(field.state)
	SaveData(field.state, "output.txt")

	return result
}

// Функция устанавливает препятствие в указанные координаты и проверяет создает ли оно петлю движения стражника
// true - препятствие зацикливает стражника, false - препятствие не зацикливает стражника
func isObstacleUseful(f Field, x int, y int) (result bool) {

	// ограничение на количество ходов, для исключения бесконечного хождения по кругу внутри поля
	// представляется достаточным сделать ограничение количества равным площади поля
	moveLimit := len(f.state[0]) * len(f.state)

	// TODO -- ПЕРЕПИСАТЬ day6_2
	f.setObstacle(x, y)
	for i := range moveLimit {

		f.moveUnit()

		// прерываем цикл, если стражник вышел с поля
		if f.unit.onTheField == false {
			result = false
			break
		}

		// прерываем цикл, если стражник ходит по кругу, созданному нашим препятствием
		if f.obstacle.hitBottom > 1 || f.obstacle.hitLeft > 1 || f.obstacle.hitRight > 1 || f.obstacle.hitTop > 1 {
			result = true
			break
		}

		// если достигнут лимит - значит стражник ходит кругами по исходному полю
		if i == moveLimit {
			panic("Infinite Guardian loop on current Field")
		}
	}

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

		// для каждой строки, кроме последней допечатываем перевод строки
		if i != len(layout)-1 {
			file.WriteString("\n")
		}
	}

	file.Close()
}
