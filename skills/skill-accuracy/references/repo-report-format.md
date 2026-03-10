# Repo Report Format

Use this reference for repo-scan outputs.

## Output Channels

The repo scan should support:

- concise terminal summary
- markdown report for humans
- JSON report for integrations

Suggested persisted outputs:

- `reports/skill-accuracy/latest.md`
- `reports/skill-accuracy/latest.json`

## Terminal Summary

Keep terminal output short and decision-oriented:

- total targets found
- confidence distribution
- highest-priority targets
- strongest targets
- top cross-repo anti-patterns

## Markdown Report

Use this structure:

1. `Repo Summary`
2. `Discovery Overview`
3. `Scoreboard`
4. `Per-Target Findings`
5. `Cross-Repo Anti-Patterns`
6. `Recommended Next Actions`

## Scoreboard Fields

Each row should include:

- `path`
- `type`
- `confidence`
- `overall_score`
- `clarity`
- `testability`
- `consistency_risk`
- `honesty_conflict_risk`
- `format_control`
- `priority`
- `notes`

## Repo Summary Fields

The summary should include:

- `total_targets`
- `targets_by_type`
- `targets_by_confidence`
- `highest_priority_targets`
- `strongest_targets`
- `common_anti_patterns`
- `recommended_actions`

## JSON Contract

The JSON report should include:

- `repo_path`
- `generated_at`
- `summary`
- `targets`
- `possible_targets`

Each target object should include:

- `path`
- `type`
- `confidence`
- `reason_signals`
- `scorecard`
- `priority`
- `findings`

Use `possible_targets` for low-confidence discoveries that were not included in the main scoreboard.
