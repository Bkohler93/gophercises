package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bkohler93/gophercises-quiz/platform"
)

func main() {
	// read in csv (default to problems.csv)
	filePath := flag.String("f", "../problems.csv", "sets the filepath to search for the csv problem file")
	flag.Parse()

	q, err := platform.NewQuizFromCsv(*filePath)
	if err != nil {
		log.Fatal(err)
	}

	err = runPartOne(&q)
	if err != nil {
		log.Fatalf("error running quiz - %v", err)
	}

	// report results
	fmt.Printf("Number of questions: %d - Number correct: %d\n", len(q.Problems), q.NumRight)
}

func runPartOne(q *platform.Quiz) error {
	reader := bufio.NewReader(os.Stdin)

	for i, p := range q.Problems {
		fmt.Printf("Q: %s = ", p)
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		line = line[:len(line)-1] //remove newline character
		ans, err := strconv.Atoi(line)
		for err != nil {
			fmt.Printf("Q: %s = ", p)
			line, err = reader.ReadString('\n')
			line = line[:len(line)-1]
			ans, err = strconv.Atoi(line)
		}

		if ans == q.Ans[i] {
			q.NumRight++
		}
	}
	return nil
}
