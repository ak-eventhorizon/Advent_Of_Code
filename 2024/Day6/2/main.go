package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

// ********* Advent of Code 2024 *********
// --- Day 6: Guard Gallivant --- Puzzle 2
// https://adventofcode.com/2024/day/6

// Объект, представляющий собой игровое поле, его состояние и объекты, которые на нем находятся
type Field struct {
	state [][]string // состояние игрового поля (матрица символов из файла)

	unit struct { // Стражник "^|v|>|<", который перемещается по полю
		x            int    // координата по оси X
		y            int    // координата по оси Y
		direction    string // направление следующего шага "^" | "v" | ">" | "<"
		isOnTheField bool   // признак того, что объект находится на поле и не вышел за его пределы
	}

	obstacle struct { // Препятствие "O", которое можно разместить на поле
		x         int  // координата по оси X
		y         int  // кооридната по оси Y
		hitTop    int  // если с одной стороны было более 1 столкновения - значит стражник ходит кругами
		hitRight  int  // если с одной стороны было более 1 столкновения - значит стражник ходит кругами
		hitBottom int  // если с одной стороны было более 1 столкновения - значит стражник ходит кругами
		hitLeft   int  // если с одной стороны было более 1 столкновения - значит стражник ходит кругами
		isSet     bool // признак, что препятствие было установлено на поле
	}
}

// Заполнить структуру из содержимого поля layout
func (f *Field) fillFrom(layout [][]string) {

	// заполнение матрицы state клонированием каждого элемента, поскольку иначе слайсы передаются по ссылке и везде используется один и тот же фактический слайс
	var tmpMatrix [][]string
	for _, v := range layout {
		tmpMatrix = append(tmpMatrix, slices.Clone(v))
	}

	f.state = tmpMatrix

	// находжение на поле стражника и установка его координат и направления
	for y, line := range layout {
		for x, char := range line {
			if char == "^" || char == ">" || char == "v" || char == "<" {
				f.unit.x = x
				f.unit.y = y
				f.unit.direction = char
				f.unit.isOnTheField = true
				break // прерывание цикла поиска - подразумевается, что стражник на поле только один
			}
		}
	}

	if f.unit.isOnTheField == false {
		panic("No units on the field!")
	}
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
		f.unit.isOnTheField = false
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
	if f.unit.isOnTheField {
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
		if f.obstacle.isSet {
			f.state[f.obstacle.y][f.obstacle.x] = "."
			f.obstacle.isSet = false
		}

		f.state[y][x] = "O"
		f.obstacle.isSet = true
		f.obstacle.x = x
		f.obstacle.y = y
		f.obstacle.hitBottom = 0
		f.obstacle.hitLeft = 0
		f.obstacle.hitRight = 0
		f.obstacle.hitTop = 0
	}
}

func main() {

	startLayout := GetData("data.txt")

	answer := day6_2(startLayout)
	fmt.Println(answer)
}

func day6_2(layout [][]string) (result int) {

	// копия исходного игрового поля для сохранения на нем всех подходящих препятствий
	fieldWithUsefulObstacles := new(Field)
	fieldWithUsefulObstacles.fillFrom(layout)

	// копия исходного поля для вычисления маршрута стражника без стороннего вмешательства
	tmpField := new(Field)
	tmpField.fillFrom(layout)

	route := getDefaultRoute(tmpField)

	// поочередно ставим препятствие на каждую точку исходного маршрута стражника и проверяем помогло-ли
	for _, point := range route {
		// экземпляр поля, на которое будем устанавливать препятствие и проверять его
		field := new(Field)
		field.fillFrom(layout)

		field.setObstacle(point[0], point[1])
		if isObstacleUseful(field) {
			fieldWithUsefulObstacles.state[point[1]][point[0]] = "O"
		}
	}

	result = CalcObstacles(fieldWithUsefulObstacles.state)
	SaveData(fieldWithUsefulObstacles.state, "output.txt")

	return result
}

// Функция возвращает набор координат, описывающих все шаги стражника до выхода с поля начальной конфигурации [x,y]
func getDefaultRoute(f *Field) (result [][]int) {

	// ограничение на количество ходов, для исключения бесконечного хождения по кругу внутри поля
	// представляется достаточным сделать ограничение количества равным площади поля
	moveLimit := len(f.state[0]) * len(f.state)

	for i := range moveLimit {

		point := make([]int, 2)

		f.moveUnit()

		// прерываем цикл, если стражник вышел с поля
		if f.unit.isOnTheField == false {
			break
		} else {
			point[0] = f.unit.x
			point[1] = f.unit.y
		}

		// если достигнут лимит - значит стражник ходит кругами по исходному полю
		if i == moveLimit-1 {
			panic("Infinite Guardian loop on initial Field")
		}

		result = append(result, point)
	}

	return result
}

// Функция устанавливает препятствие в указанные координаты и проверяет создает ли оно петлю движения стражника
// true - препятствие зацикливает стражника, false - препятствие не зацикливает стражника
func isObstacleUseful(f *Field) (result bool) {

	// ограничение на количество ходов, для исключения бесконечного хождения по кругу внутри поля
	// представляется достаточным сделать ограничение количества равным площади поля
	moveLimit := len(f.state[0]) * len(f.state)

	for {

		f.moveUnit()
		moveLimit--

		// прерываем цикл, если стражник вышел с поля
		if f.unit.isOnTheField == false {
			result = false
			break
		}

		// прерываем цикл, если стражник ходит по кругу, в котором циклически участвует наше препятствие
		if f.obstacle.hitBottom > 1 || f.obstacle.hitLeft > 1 || f.obstacle.hitRight > 1 || f.obstacle.hitTop > 1 {
			result = true
			break
		}

		// если достигнут лимит - значит стражник ходит по кругу, в котором циклически не участвует наше препятствие
		if moveLimit == 0 {
			result = true
			break
		}
	}

	return result
}

// Функция вычисляет количество препятствий "O" на поле.
func CalcObstacles(layout [][]string) (count int) {

	for _, line := range layout {
		for _, char := range line {
			if char == "O" {
				count++
			}
		}
	}

	return count
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
