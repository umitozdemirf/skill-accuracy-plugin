package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Case struct {
	ID            string   `json:"id"`
	Intent        string   `json:"intent"`
	Prompt        string   `json:"prompt"`
	TargetedRules []string `json:"targeted_rules"`
	WhyItMatters  string   `json:"why_it_matters"`
}

type Run struct {
	CaseID         string   `json:"case_id"`
	Intent         string   `json:"intent"`
	Prompt         string   `json:"prompt"`
	TargetedRules  []string `json:"targeted_rules,omitempty"`
	RunIndex       int      `json:"run_index"`
	Command        []string `json:"command"`
	Output         string   `json:"output"`
	Stderr         string   `json:"stderr,omitempty"`
	ExitCode       int      `json:"exit_code"`
	DurationMillis int64    `json:"duration_millis"`
}

func main() {
	assetPath := flag.String("asset", "", "Path to the target SKILL.md or instruction file")
	casesPath := flag.String("cases", "", "Path to generated cases JSON")
	outputPath := flag.String("output", "", "Path to write run results JSON")
	runs := flag.Int("runs", 5, "Number of repeated runs per prompt")
	model := flag.String("model", "", "Optional Claude model name")
	cwd := flag.String("cwd", "", "Optional working directory for Claude")
	allowedTools := flag.String("allowed-tools", "", "Optional comma-separated allowed tools")
	flag.Parse()

	if *assetPath == "" || *casesPath == "" || *outputPath == "" {
		fmt.Fprintln(os.Stderr, "run_claude_trials requires -asset, -cases, and -output")
		os.Exit(1)
	}
	if *runs <= 0 {
		fmt.Fprintln(os.Stderr, "-runs must be > 0")
		os.Exit(1)
	}

	assetBytes, err := os.ReadFile(*assetPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	systemPrompt := strings.TrimSpace(string(assetBytes))

	caseBytes, err := os.ReadFile(*casesPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var cases []Case
	if err := json.Unmarshal(caseBytes, &cases); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	results := make([]Run, 0, len(cases)**runs)
	for _, item := range cases {
		for runIndex := 1; runIndex <= *runs; runIndex++ {
			command := buildCommand(systemPrompt, item.Prompt, *model, *allowedTools)
			cmd := exec.Command(command[0], command[1:]...)
			if *cwd != "" {
				cmd.Dir = *cwd
			}

			started := time.Now()
			output, err := cmd.CombinedOutput()
			duration := time.Since(started)

			exitCode := 0
			stderr := ""
			text := strings.TrimSpace(string(output))
			if err != nil {
				if exitErr, ok := err.(*exec.ExitError); ok {
					exitCode = exitErr.ExitCode()
				} else {
					stderr = err.Error()
				}
			}

			results = append(results, Run{
				CaseID:         item.ID,
				Intent:         item.Intent,
				Prompt:         item.Prompt,
				TargetedRules:  item.TargetedRules,
				RunIndex:       runIndex,
				Command:        command,
				Output:         text,
				Stderr:         stderr,
				ExitCode:       exitCode,
				DurationMillis: duration.Milliseconds(),
			})
		}
	}

	encoded, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := os.WriteFile(*outputPath, encoded, 0o644); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func buildCommand(systemPrompt, taskPrompt, model, allowedTools string) []string {
	command := []string{"claude", "--print", "--append-system-prompt", systemPrompt}
	if model != "" {
		command = append(command, "--model", model)
	}
	if allowedTools != "" {
		command = append(command, "--allowedTools", allowedTools)
	}
	command = append(command, taskPrompt)
	return command
}
