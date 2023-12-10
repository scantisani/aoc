package part1

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func SumOfPredictionsFromInput() int {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	return SumOfPredictions(lines[:200]) // skip last (blank) line
}

func SumOfPredictions(histories []string) int {
	sum := 0
	for _, history := range histories {
		sum += PredictionFor(history)
	}
	return sum
}

func PredictionFor(history string) int {
	values := parseValues(history)

	extrapolations := ExtrapolationsFor(values)
	return predictedValueFor(values, extrapolations)
}

func predictedValueFor(currentValues []int, extrapolations [][]int) int {
	currentLast := currentValues[len(currentValues)-1]

	if len(extrapolations) == 1 {
		extrapolation := extrapolations[0]
		extrapolationLast := extrapolation[len(extrapolation)-1]

		return currentLast + extrapolationLast
	}

	return currentLast + predictedValueFor(extrapolations[0], extrapolations[1:])
}

func parseValues(history string) []int {
	matches := regexp.MustCompile(`-?\d+`).FindAllString(history, -1)

	values := make([]int, 0)
	for _, match := range matches {
		value, _ := strconv.Atoi(match)
		values = append(values, value)
	}

	return values
}

func ExtrapolationsFor(values []int) [][]int {
	extrapolations := make([][]int, 0)
	currentExtrapolation := values

	for !allZero(currentExtrapolation) {
		currentExtrapolation = ExtrapolationFor(currentExtrapolation)
		extrapolations = append(extrapolations, currentExtrapolation)
	}

	return extrapolations
}

func ExtrapolationFor(values []int) []int {
	extrapolation := make([]int, 0)

	for i := 1; i < len(values); i++ {
		diff := values[i] - values[i-1]
		extrapolation = append(extrapolation, diff)
	}

	return extrapolation
}

func allZero(values []int) bool {
	for _, value := range values {
		if value != 0 {
			return false
		}
	}

	return true
}
