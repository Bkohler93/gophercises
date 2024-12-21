package platform

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Quiz struct {
	Problems []string
	Ans      []int
	NumRight int
}

func NewQuizFromCsv(filePath string) (Quiz, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return Quiz{}, fmt.Errorf("error opening file - %v", err)
	}
	defer f.Close()

	ps := make([]string, 0)
	as := make([]int, 0)
	reader := csv.NewReader(f)

	for {
		r, err := reader.Read()
		if err != nil && err != io.EOF {
			return Quiz{}, fmt.Errorf("error reading csv - %v", err)
		}

		if err == io.EOF {
			break
		}

		ps = append(ps, r[0])
		a, err := strconv.Atoi(r[1])
		if err != nil {
			log.Fatalf("error reading in answer from csv - %v", err)
		}
		as = append(as, a)
	}

	q := Quiz{
		Problems: ps,
		Ans:      as,
		NumRight: 0,
	}
	return q, nil
}
