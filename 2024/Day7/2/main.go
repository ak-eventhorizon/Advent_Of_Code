package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// ********* Advent of Code 2024 *********
// --- Day 7: Bridge Repair --- Puzzle 2
// https://adventofcode.com/2024/day/7

type Case struct {
	value   int
	numbers []int
}

func main() {

	data := GetData("data.txt")
	answer := day7_2(data)
	fmt.Println(answer)
}

// Функция возвращает результат по условиям задачи
func day7_2(cases []Case) (result int) {

	for _, c := range cases {

		operatorsCombinations := GetAllCombinations([]string{"+", "*", "|"}, len(c.numbers)-1)

		// fmt.Print("For case ", c, " ---> ")                 // DEBUG print
		// fmt.Println("Combinations:", operatorsCombinations) // DEBUG print

		for _, operators := range operatorsCombinations {
			expression := []string{}
			for i := 0; i < len(c.numbers); i++ {
				if i == len(c.numbers)-1 { // последний элемент выражения
					expression = append(expression, strconv.Itoa(c.numbers[i]))
				} else {
					expression = append(expression, strconv.Itoa(c.numbers[i]), string(operators[i]))
				}
			}

			exprResult := calcExpression(expression)

			// fmt.Print(expression)               // DEBUG print
			// fmt.Print("--------> ", exprResult) // DEBUG print

			if c.value == exprResult {
				// fmt.Println(" <---- BINGO!!!") // DEBUG print
				result += exprResult
				break
			} else {
				// fmt.Println("") // DEBUG print
			}
		}
	}

	return result
}

// Функция вычисляет выражение expr, переданное слайсом строк
func calcExpression(expr []string) (res int) {

	// ЭТОТ КОД РАБОТАЕТ БЕЗ УЧЕТА ПРЕИМУЩЕСТВА ОПЕРАЦИИ УМНОЖЕНИЯ ПЕРЕД СЛОЖЕНИЕМ - просто слева направо, как указано в условиях задачи
	// копия исходного слайса для его изменения
	tmpExpression := slices.Clone(expr)

	// цикл по знакам операций, в слайсе tmpExpr - знак операции это каждый второй элемент, начиная со второго -- [11 * 6 + 16 * 20]
	for i := 1; i < len(tmpExpression)-1; i = i + 2 {

		opA, err := strconv.Atoi(tmpExpression[i-1]) // операнд А - до знака
		if err != nil {
			panic(err)
		}
		opB, err := strconv.Atoi(tmpExpression[i+1]) // операнд B - после знака
		if err != nil {
			panic(err)
		}

		// действие над операндами зависит от оператора между ними, результат помещается во второй операнд для следующего цикла
		switch tmpExpression[i] {
		case "+":
			tmpExpression[i+1] = strconv.Itoa(opA + opB)
		case "*":
			tmpExpression[i+1] = strconv.Itoa(opA * opB)
		case "|":
			tmpExpression[i+1] = strconv.Itoa(opA) + strconv.Itoa(opB)
		}
	}

	tmp, err := strconv.Atoi(tmpExpression[len(tmpExpression)-1])
	if err != nil {
		panic(err)
	}

	res = tmp

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

// Рекурсия для получения всех комбинаций символов из набора
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
