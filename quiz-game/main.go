package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type Problem struct {
	question string
	answer   string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in format 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
problemLoop:
	for i, problem := range problems {
		fmt.Printf("Problem: #%d: %s = ", i+1, problem.question)
		answerChannel := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChannel <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemLoop
		case answer := <-answerChannel:
			if answer == problem.answer {
				correct++
			}
		}

	}

	fmt.Printf("You scored %d out of %d.", correct, len(problems))
}

func parseLines(lines [][]string) []Problem {
	result := make([]Problem, len(lines))

	for i, line := range lines {
		result[i] = Problem{
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
