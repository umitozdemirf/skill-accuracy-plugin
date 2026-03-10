package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Run struct {
	Prompt string `json:"prompt"`
	Output string `json:"output"`
}

type Summary struct {
	TotalRuns        int                `json:"total_runs"`
	DistinctOutputs  int                `json:"distinct_outputs"`
	ConsistencyRate  float64            `json:"consistency_rate"`
	TopOutputPreview string             `json:"top_output_preview"`
	ByPrompt         map[string]Metrics `json:"by_prompt"`
}

type Metrics struct {
	TotalRuns       int     `json:"total_runs"`
	DistinctOutputs int     `json:"distinct_outputs"`
	ConsistencyRate float64 `json:"consistency_rate"`
}

func main() {
	inputPath := flag.String("input", "", "Path to a JSON array of {prompt, output} objects")
	flag.Parse()

	if *inputPath == "" {
		fmt.Fprintln(os.Stderr, "summarize_runs requires -input")
		os.Exit(1)
	}

	raw, err := os.ReadFile(*inputPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var runs []Run
	if err := json.Unmarshal(raw, &runs); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	summary := summarize(runs)
	encoded, err := json.MarshalIndent(summary, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(string(encoded))
}

func summarize(runs []Run) Summary {
	byPrompt := map[string][]string{}
	allOutputs := map[string]int{}
	for _, run := range runs {
		normalized := normalize(run.Output)
		byPrompt[run.Prompt] = append(byPrompt[run.Prompt], normalized)
		allOutputs[normalized]++
	}

	summary := Summary{
		TotalRuns: len(runs),
		ByPrompt:  map[string]Metrics{},
	}

	topCount := 0
	for output, count := range allOutputs {
		if count > topCount {
			topCount = count
			summary.TopOutputPreview = truncate(output, 140)
		}
	}
	summary.DistinctOutputs = len(allOutputs)
	if len(runs) > 0 {
		summary.ConsistencyRate = float64(topCount) / float64(len(runs))
	}

	for prompt, outputs := range byPrompt {
		counts := map[string]int{}
		maxCount := 0
		for _, output := range outputs {
			counts[output]++
			if counts[output] > maxCount {
				maxCount = counts[output]
			}
		}
		summary.ByPrompt[prompt] = Metrics{
			TotalRuns:       len(outputs),
			DistinctOutputs: len(counts),
			ConsistencyRate: float64(maxCount) / float64(len(outputs)),
		}
	}

	return summary
}

func normalize(value string) string {
	return strings.Join(strings.Fields(strings.ToLower(value)), " ")
}

func truncate(value string, limit int) string {
	if len(value) <= limit {
		return value
	}
	return value[:limit] + "..."
}
