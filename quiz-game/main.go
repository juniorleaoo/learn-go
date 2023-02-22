package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	question string
	answer   string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in format 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFilename))
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse the provied CSV file.")
	}

	problems := parseLines(lines)

	correct := 0
	for i, problem := range problems {
		fmt.Printf("Problem: #%d: %s = ", i+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	result := make([]problem, len(lines))

	for i, line := range lines {
		result[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}

	return result
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
