package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ********* Advent of Code 2024 *********
// --- Day 7: Bridge Repair --- Puzzle 1
// https://adventofcode.com/2024/day/7

type Case struct {
	value   int
	numbers []int
}

func main() {

	data := GetData("data.txt")

	answer := day7_1(data)
	fmt.Println(answer)
}

func day7_1(cases []Case) (result int) {

	for _, c := range cases {

		allCombinations := GetAllCombinations([]string{"+", "*"}, len(c.numbers)-1)

		fmt.Print("For case ", c, " ---> ")
		fmt.Println("Combinations:", allCombinations)

		for _, combination := range allCombinations {
			expression := []string{}
			for i := 0; i < len(c.numbers); i++ {
				if i == len(c.numbers)-1 {
					expression = append(expression, strconv.Itoa(c.numbers[i]))
				} else {
					expression = append(expression, strconv.Itoa(c.numbers[i]), string(rune(combination[i])))
				}
			}
			fmt.Print(expression)
			fmt.Println("--->", calcExpression(expression))
		}
	}

	// fmt.Println(calcExpression2([]string{"2", "+", "2", "*", "2"}))
	// fmt.Println(calcExpression2([]string{"2", "*", "2", "*", "3"}))
	// fmt.Println(calcExpression2([]string{"2", "+", "2", "+", "3"}))

	return result
}

// Функция вычисляет выражение expr, переданное слайсом строк
// Не является полноценным парсером выражений, это ограниченный прототип только для условий задачи
// Колхоз, но нет порыва писать AST...
func calcExpression(expr []string) (res int) {

	addition := []int{}
	multiplication := []int{}

	//TODO <<<
	// Делим по символу "+" 11*6+16*20*22+5+6 ---> 11*6   16*20*22   5   6
	// Перемножаем каждый
	// Складываем все вместе

	// извлекаем перемножаемые числа, извлеченные значения наменяются на ""
	for i := 0; i < len(expr); i++ {

		if expr[i] == "*" {

			val1, err := strconv.Atoi(expr[i+1])
			if err != nil {
				val1 = 1
			}

			val2, err := strconv.Atoi(expr[i-1])
			if err != nil {
				val2 = 1
			}

			multiplication = append(multiplication, val1, val2)

			expr[i-1] = ""
			expr[i] = ""
			expr[i+1] = ""
		}
	}

	// извлекаем слагаемые числа
	for i := 0; i < len(expr); i++ {

		if expr[i] == "+" {

			val1, err := strconv.Atoi(expr[i+1])
			if err != nil {
				val1 = 0
			}

			val2, err := strconv.Atoi(expr[i-1])
			if err != nil {
				val2 = 0
			}

			addition = append(addition, val1, val2)

			expr[i-1] = ""
			expr[i] = ""
			expr[i+1] = ""
		}
	}

	var mul int

	if len(multiplication) == 0 {
		mul = 0
	} else {
		mul = 1
		for _, v := range multiplication {
			mul *= v
		}
	}

	// fmt.Println(expr, " MULT>>>>>>>>>>>>>>", multiplication)
	// fmt.Println(expr, " ADD>>>>>>>>>>>>>>", addition)

	for _, v := range addition {
		res += v
	}

	res += mul

	return res
}

// Функция возвращает все комбинации длиной k, которые можно составить из набора символов set;
// Для set = ["A", "B"], k = 3;
// [AAA AAB ABA ABB BAA BAB BBA BBB]
func GetAllCombinations(set []string, k int) (list []string) {
	n := len(set)
	list = GetAllCombinationsRec(set, "", n, k)

	return list
}

// Рекурсия для получения всех комбинаций
func GetAllCombinationsRec(set []string, prefix string, n int, k int) (res []string) {

	// базовый случай - выход из рекурсии
	if k == 0 {
		return []string{prefix}
	}

	// один за одним добавляем все символы из набора и рекурсивно вызываем для длины k-1
	for i := range n {
		newPrefix := prefix + set[i]
		res = append(res, GetAllCombinationsRec(set, newPrefix, n, k-1)...)
	}

	return res
}

// Функция извлекает из текстового файла все условия задачи.
func GetData(filename string) (cases []Case) {

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		var str []string
		var nums []int

		part1 := strings.Split(scanner.Text(), ":")[0]
		part2 := strings.Split(strings.Trim(strings.Split(scanner.Text(), ":")[1], " "), " ")

		str = append(str, part1)
		str = append(str, part2...)

		for _, v := range str {
			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			nums = append(nums, num)
		}

		var currentCase = new(Case)
		currentCase.value = nums[0]
		currentCase.numbers = nums[1:]

		cases = append(cases, *currentCase)
	}

	file.Close()

	return cases
}
