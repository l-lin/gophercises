package problem

import (
	"encoding/csv"
	"io"
)

// CsvProblemsParser extracts problems from a csv content
type CsvProblemsParser struct {
}

// Parse problems from a csv file
func (pe *CsvProblemsParser) Parse(in io.Reader) ([]*Problem, error) {
	reader := csv.NewReader(in)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	pbs := make([]*Problem, 0)
	for _, line := range lines {
		pb := &Problem{
			Question: line[0],
			Answer:   line[1],
		}
		pbs = append(pbs, pb)
	}
	return pbs, nil
}
