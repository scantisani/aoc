package part2

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

type problemBlock struct {
	start int
	end   int
	op    operation
}

func buildProblems(lines []string) []Problem {
	var problems []Problem

	lastLine := lines[len(lines)-1]
	blocks := extractBlocks(lastLine)

	for _, block := range blocks {
		problem := buildProblem(lines[:len(lines)-1], block)
		problems = append(problems, problem)
	}

	return problems
}

func buildProblem(lines []string, block problemBlock) Problem {
	var numbers []int

	for i := block.end; i >= block.start; i-- {
		var digits []rune
		for _, line := range lines {
			digit := line[i]
			if digit != ' ' {
				digits = append(digits, rune(digit))
			}
		}

		if len(digits) > 0 {
			number, _ := strconv.Atoi(string(digits))
			numbers = append(numbers, number)
		}
	}

	return Problem{numbers, block.op}
}

func extractBlocks(line string) []problemBlock {
	r := regexp.MustCompile("[+*]\\s+")
	splits := r.FindAllStringIndex(line, -1)

	var blocks []problemBlock
	for _, split := range splits {
		start := split[0]
		end := split[1] - 1 // use an inclusive instead of exclusive range

		newBlock := problemBlock{start: start, end: end}
		if line[start] == '*' {
			newBlock.op = multiply
		}

		blocks = append(blocks, newBlock)
	}
	return blocks
}
