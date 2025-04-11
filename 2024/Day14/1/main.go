package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

// Структура представляет поле (двумерную матрицу), обладающее шириной lenX клеток и высотой lenY клеток
type Field struct {
	lenX int
	lenY int
}

// Структура представляет робота
// Position: точка в которой робот располагается в данное время
// Velocity: смещение на которое робот перемещается за 1 ход (1 секунда)
type Robot struct {
	Position Point
	Velocity Point
	Field    Field
}

func (r *Robot) Move() {
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

func day14_1(robots []Robot) (result int) {

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

		field := Field{101, 103} // размер поля по условиям задачи
		position := Point{nums[0], nums[1]}
		velocity := Point{nums[2], nums[3]}
		robot := Robot{position, velocity, field}

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
