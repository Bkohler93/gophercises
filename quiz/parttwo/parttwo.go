package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/bkohler93/gophercises-quiz/platform"
)

func main() {
	filePath := flag.String("f", "../problems.csv", "sets the filepath to search for the csv problem file")
	time := flag.Int("t", 30, "sets the time limit for the quiz")

	flag.Parse()

	q, err := platform.NewQuizFromCsv(*filePath)
	if err != nil {
		log.Fatal(err)
	}

	err = runPartTwo(&q, *time)
	if err != nil {
		log.Fatalf("error running quiz - %v", err)
	}

	// report results
	fmt.Printf("Number of questions: %d - Number correct: %d\n", len(q.Problems), q.NumRight)
}

func runPartTwo(q *platform.Quiz, t int) error {
	timer := time.NewTimer(time.Second * time.Duration(t))

	ch := make(chan bool)

	go func() {
		reader := bufio.NewReader(os.Stdin)

		for i, p := range q.Problems {
			fmt.Printf("Q: %s = ", p)
			line, _ := reader.ReadString('\n')
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
		ch <- true
	}()

	for {
		select {
		case <-timer.C:
			fmt.Println("\ntime's up!")
			return nil
		case <-ch:
			return nil
		}
	}
}
