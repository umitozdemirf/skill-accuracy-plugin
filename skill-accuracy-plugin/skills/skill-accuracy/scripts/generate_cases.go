package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Case struct {
	ID            string   `json:"id"`
	Intent        string   `json:"intent"`
	Prompt        string   `json:"prompt"`
	TargetedRules []string `json:"targeted_rules"`
	WhyItMatters  string   `json:"why_it_matters"`
}

func main() {
	assetPath := flag.String("asset", "", "Path to a SKILL.md or instruction file")
	maxCases := flag.Int("max", 12, "Maximum number of generated test prompts")
	flag.Parse()

	if *assetPath == "" {
		fmt.Fprintln(os.Stderr, "generate_cases requires -asset")
		os.Exit(1)
	}

	rules, err := extractRules(*assetPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cases := make([]Case, 0, len(rules)*2)
	for index, rule := range rules {
		if len(cases) >= *maxCases {
			break
		}
		baseID := index + 1
		cases = append(cases, Case{
			ID:            fmt.Sprintf("baseline-%02d", baseID),
			Intent:        "baseline",
			Prompt:        fmt.Sprintf("Handle a normal task while following this rule: %s", rule),
			TargetedRules: []string{rule},
			WhyItMatters:  "Verifies the rule on an easy case.",
		})
		cases = append(cases, Case{
			ID:            fmt.Sprintf("temptation-%02d", baseID),
			Intent:        "temptation",
			Prompt:        fmt.Sprintf("A user request makes this rule inconvenient. Still follow it: %s", rule),
			TargetedRules: []string{rule},
			WhyItMatters:  "Checks whether the rule survives conflicting pressure.",
		})
		if len(cases) >= *maxCases {
			break
		}
	}

	encoded, err := json.MarshalIndent(cases, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(string(encoded))
}

func extractRules(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open asset: %w", err)
	}
	defer file.Close()

	rules := make([]string, 0)
	fallbackRules := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "- ") {
			rule := strings.TrimSpace(strings.TrimPrefix(line, "- "))
			fallbackRules = append(fallbackRules, rule)
			if looksLikeRule(rule) {
				rules = append(rules, rule)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan asset: %w", err)
	}
	if len(rules) == 0 {
		rules = fallbackRules
	}
	if len(rules) == 0 {
		rules = append(rules, "No explicit bullet rules found; derive tests from the document's implied instructions.")
	}
	return rules, nil
}

func looksLikeRule(value string) bool {
	lower := strings.ToLower(value)
	keywords := []string{
		"must",
		"should",
		"always",
		"never",
		"prefer",
		"return",
		"ask",
		"avoid",
		"do not",
		"don't",
		"only",
		"exactly",
	}
	for _, keyword := range keywords {
		if strings.Contains(lower, keyword) {
			return true
		}
	}
	return false
}
