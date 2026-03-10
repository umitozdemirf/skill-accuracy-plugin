# Repo Report Format

Use this reference for repo-scan and quick-analysis outputs across Claude and Codex.

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

The first table shown to the user should be the main scoreboard for scoreable targets. Do not start with a discovery-only table that omits `overall_score`.

## Markdown Report

Use this structure:

1. `Repo Summary`
2. `Possible Targets Not Scored`
3. `Per-Target Findings`
4. `Cross-Repo Anti-Patterns`
5. `Recommended Next Actions`

## Quick Analysis Report

Use this shorter structure:

1. `Quick Scoreboard`
2. `Short Findings`
3. `Next Actions`

Quick-analysis should favor brevity over exhaustive justification.

Recommended quick scoreboard fields:

- `path`
- `type`
- `confidence`
- `scope`
- `overall`
- `clarity`
- `testability`
- `risk`
- `priority`

## Scoreboard Fields

Each row should include:

- `path`
- `type`
- `confidence`
- `platform`
- `scope`
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

For scoreable targets, `Repo Summary` must render a table using the scoreboard fields directly:

- `path`
- `type`
- `confidence`
- `platform`
- `scope`
- `overall_score`
- `clarity`
- `testability`
- `consistency_risk`
- `honesty_conflict_risk`
- `format_control`
- `priority`

Items that are not instruction assets or are not scoreable should be moved to `Possible Targets Not Scored` with:

- `path`
- `type`
- `confidence`
- `platform`
- `scope`
- `reason_not_scored`

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
- `platform`
- `scope`
- `scorecard`
- `priority`
- `findings`

Use `possible_targets` for low-confidence discoveries that were not included in the main scoreboard.

For Codex reports, group the output into:

- `Repo Targets`
- `User Skill Targets`

Keep the grouping visible in both quick-analysis and full repo-scan outputs.
