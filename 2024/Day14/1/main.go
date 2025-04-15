package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ********* Advent of Code 2024 *********
// --- Day 14: Restroom Redoubt --- Puzzle 1
// https://adventofcode.com/2024/day/14

const INPUT_FILE_PATH string = "data.txt"
const OUTPUT_FILE_PATH string = "output.txt"

// Структура представляет координаты на двумерной плоскости
type Point struct {
	x int
	y int
}

// Структура представляет поле (двумерную матрицу), обладающее шириной lenX клеток, высотой lenY клеток и содержимым layout
type Field struct {
	lenX   int
	lenY   int
	layout [][]string
}

func (f *Field) Init(lenX, lenY int) {
	f.lenX = lenX
	f.lenY = lenY

	for range lenY {
		line := []string{}
		for range lenX {
			line = append(line, "0")
		}
		f.layout = append(f.layout, line)
	}
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

	_, err3 := file.WriteString("\n")
	if err3 != nil {
		panic(err3)
	}

}

// Структура представляет робота
// Position: точка в которой робот располагается в данное время
// Velocity: смещение на которое робот перемещается за 1 ход (1 секунда)
// Field: указатель на поле, на котором располагается робот
type Robot struct {
	Position Point
	Velocity Point
	Field    *Field
}

func (r *Robot) SetOnField(field *Field) {
	r.Field = field

	targetCellValue := r.Field.layout[r.Position.y][r.Position.x]
	intValue, err := strconv.Atoi(targetCellValue)
	if err != nil {
		panic(err)
	}

	intValue += 1
	targetCellValue = strconv.Itoa(intValue)

	r.Field.layout[r.Position.y][r.Position.x] = targetCellValue
}

func (r *Robot) Move() {

	// уменьшаем на единицу значения клетки, из которой робот уходит
	startCell := r.Field.layout[r.Position.y][r.Position.x]
	cellValue, err := strconv.Atoi(startCell)
	if err != nil {
		panic(err)
	}
	cellValue -= 1
	startCell = strconv.Itoa(cellValue)
	r.Field.layout[r.Position.y][r.Position.x] = startCell

	// вычисление клетки, в которую робот перемещается
	x := r.Position.x + r.Velocity.x
	y := r.Position.y + r.Velocity.y

	// при достижении границы поля - появление с противоположной стороны
	if x >= r.Field.lenX {
		r.Position.x = x - r.Field.lenX
	} else if x < 0 {
		r.Position.x = r.Field.lenX + x
	} else {
		r.Position.x = x
	}

	// при достижении границы поля - появление с противоположной стороны
	if y >= r.Field.lenY {
		r.Position.y = y - r.Field.lenY
	} else if y < 0 {
		r.Position.y = r.Field.lenY + y
	} else {
		r.Position.y = y
	}

	// увеличиваем на единицу значения клетки, в которую робот попадает
	endCell := r.Field.layout[r.Position.y][r.Position.x]
	cellValue, err = strconv.Atoi(endCell)
	if err != nil {
		panic(err)
	}
	cellValue += 1
	endCell = strconv.Itoa(cellValue)
	r.Field.layout[r.Position.y][r.Position.x] = endCell
}

func (r *Robot) Display() {
	fmt.Printf("Position(%v,%v) -- Velocity(%v,%v) -- Field(%v,%v)\n",
		r.Position.x, r.Position.y,
		r.Velocity.x, r.Velocity.y,
		r.Field.lenX, r.Field.lenY)
}

func main() {
	start := time.Now()

	input := GetData(INPUT_FILE_PATH)
	answer := day14_1(input)
	fmt.Println(answer)

	fmt.Printf("%s \n", time.Since(start)) // время выполнения функции
}

func day14_1(input []Robot) (result int) {

	field := new(Field)  // создание поля, на кототором будет происходить все движение
	field.Init(101, 103) // размер поля по условию задачи 101x103

	robots := []Robot{} // всех роботов копируем в этот слайс, чтобы input оставить неизменным

	// размещаем роботов на поле и сохраняем в рабочий слайс
	for _, robot := range input {
		robot.SetOnField(field)
		robots = append(robots, robot)
	}

	// 100 итераций движения роботов на поле
	for _, robot := range robots {
		for range 100 {
			robot.Move()
		}
	}

	field.SaveToFile(OUTPUT_FILE_PATH)

	result = calcSafetyFactor(*field)
	return result
}

// Функция производит вычисление поля по четвертям в соответствии с формулировкой задачи
// Разделение на четверти матрицы 11x7:
//
//	00000  00000
//	0Q100  00Q20
//	00000  00000
//
//	00000  00000
//	0Q300  00Q40
//	00000  00000
func calcSafetyFactor(f Field) (result int) {

	separatorX := int(math.Floor(float64(f.lenX / 2))) // средний столбец в матрице (разделитель пополам по оси X)
	separatorY := int(math.Floor(float64(f.lenY / 2))) // средняя строка в матрице (разделитель пополам по оси Y)

	var Q1, Q2, Q3, Q4 int

	for y := range f.lenY {
		for x := range f.lenX {

			if f.layout[y][x] != "0" {
				value, err := strconv.Atoi(f.layout[y][x])
				if err != nil {
					panic(err)
				}

				if x < separatorX && y < separatorY { // попадает в первую четверть Q1
					Q1 += value
				} else if x > separatorX && y < separatorY { // попадает во вторую четверть Q2
					Q2 += value
				} else if x < separatorX && y > separatorY { // попадает в третью четверть Q3
					Q3 += value
				} else if x > separatorX && y > separatorY { // попадает в четвертую четверть Q4
					Q4 += value
				}
			}
		}
	}

	result = Q1 * Q2 * Q3 * Q4
	return result
}

// Функция извлекает из файла filename набор исходных данных
func GetData(filename string) (robots []Robot) {

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		pattern := regexp.MustCompile(`(-?\d+(?:\.\d+)?)`) // выхватывает числа и отрицательные числа
		substrings := pattern.FindAllString(scanner.Text(), -1)

		var nums []int
		for _, s := range substrings {
			num, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}

		position := Point{nums[0], nums[1]}
		velocity := Point{nums[2], nums[3]}
		robot := Robot{position, velocity, nil} // на этом этапе роботы еще не размещены на поле

		robots = append(robots, robot)
	}

	file.Close()

	return robots
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
