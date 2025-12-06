package part1

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Solve() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return SumSolutions(lines)
}

func SumSolutions(lines []string) int {
	problems := buildProblems(lines)

	sum := 0
	for _, problem := range problems {
		sum += problem.solve()
	}
	return sum
}

type Problem struct {
	arguments []int
	op        operation
}

func (p *Problem) addArgument(argument int) {
	p.arguments = append(p.arguments, argument)
}

func (p *Problem) solve() int {
	if p.op == add {
		sum := 0
		for _, arg := range p.arguments {
			sum += arg
		}
		return sum
	} else {
		sum := 1
		for _, arg := range p.arguments {
			sum *= arg
		}
		return sum
	}
}

type operation int

const (
	add      operation = iota
	multiply operation = iota
)

func buildProblems(lines []string) []Problem {
	var problems []Problem

	firstNumbers := extractNumbers(lines[0])
	for _, firstNumber := range firstNumbers {
		problems = append(problems, Problem{arguments: []int{firstNumber}})
	}

	for _, line := range lines[1:] {
		isFinal := strings.Contains(line, "+")
		if isFinal {
			operations := extractOperations(line)
			for i, operation := range operations {
				problems[i].op = operation
			}
		} else {
			numbers := extractNumbers(line)
			for i, number := range numbers {
				problems[i].addArgument(number)
			}
		}
	}

	return problems
}

func extractNumbers(line string) []int {
	r := regexp.MustCompile("\\s+")
	splits := r.Split(line, -1)

	var nums []int
	for _, split := range splits {
		num, _ := strconv.Atoi(split)
		nums = append(nums, num)
	}
	return nums
}

func extractOperations(line string) []operation {
	r := regexp.MustCompile("\\s+")
	splits := r.Split(line, -1)

	var operations []operation
	for _, split := range splits {
		if split == "+" {
			operations = append(operations, add)
		} else {
			operations = append(operations, multiply)
		}
	}
	return operations
}
