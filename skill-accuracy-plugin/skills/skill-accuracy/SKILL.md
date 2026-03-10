---
name: skill-accuracy
description: Analyze, scan, rewrite, and compare skills or agent files inside Claude Code by discovering instruction assets, generating test prompts, checking adherence and consistency, and improving instruction quality.
---

# Skill Accuracy

Use this skill when the user wants Claude Code to evaluate, scan, improve, or compare another skill, system prompt, agent file, repo policy file, or instruction bundle rather than directly solve the target task.

This skill is generic. It is not limited to coding skills.

Applicable targets include:

- coding skills
- documentation or writing skills
- research assistants
- support workflows
- planning agents
- repo guidance files such as `CLAUDE.md`
- any instruction file where repeated behavior matters

## What This Skill Does

When triggered, do this:

1. Read the instruction asset the user wants analyzed.
2. If the target is a repo or directory, discover likely instruction assets automatically.
3. Classify each target and record discovery confidence.
4. Extract explicit rules, implicit goals, output constraints, and likely failure points.
5. Generate a compact but diverse test prompt set.
6. If execution is requested and possible, simulate or run the same prompts multiple times through the target workflow.
7. Score adherence and consistency.
8. Classify weak instructions.
9. Propose sharper wording or structure changes.

The thing being evaluated is the target instruction asset as used by Claude Code, not the base model in isolation.

## Modes

Choose the lightest mode that answers the question.

### Mode 1: Analyze

Use when the user wants:

- a critique of the skill or agent file
- a proposed test plan
- suggested improvements without executing trials

In this mode:

- inspect the file
- list rules and ambiguities
- generate candidate prompts
- describe what would be measured

### Mode 2: Repo Scan

Use when the user wants:

- a repository-level audit
- automatic discovery of skill, agent, or instruction files
- a ranked scoreboard across multiple targets

In this mode:

- scan the repo for high-confidence targets first
- add medium-confidence targets only when content signals support them
- classify each discovered target
- analyze each target conservatively
- produce a repo-level scoreboard and priority summary

### Mode 3: Rewrite

Use when the user wants the target asset updated rather than only reviewed.

In this mode:

- start from a prior analysis or perform one quickly
- identify the highest-impact fixes first
- preserve the target's intent and structure where possible
- back up the original when editing local files
- patch only the requested scope if the user constrained the rewrite

Typical rewrite requests:

- "apply the recommended rewrites"
- "fix only the high-impact issues"
- "make this skill more deterministic"

### Mode 4: Compare

Use when the user wants to know whether one version is better than another.

In this mode:

- compare v1 vs v2 or original vs rewritten
- identify which rules improved or regressed
- compare predicted or observed scores
- summarize what changed materially

### Mode 5: Repeated prompt trials

Use when the user wants consistency measurement inside Claude Code and repeated testing is practical.

In this mode:

- generate or accept a prompt set
- run or simulate each prompt multiple times
- store raw outputs
- compare outputs for adherence and consistency

### Mode 6: Workspace task trials

Use when the target skill or agent changes files, executes commands, or performs multi-step work.

In this mode:

- create isolated workspaces per run
- replay the same task multiple times
- score resulting files, diffs, and outputs

## Inputs

Accept any of these:

- a local `SKILL.md`
- a local skill directory
- an agent or system prompt file
- a repo root or directory containing instruction files
- the name of a skill already available in Claude Code
- a pasted instruction block

If the user does not provide a test set, generate one.

## Typical User Lifecycle

The default single-target product flow should be:

1. `analyze`
2. `rewrite`
3. `compare`

If the user starts with analysis only, do not force them through the rest.
But when the user asks how to improve the target, move naturally into rewrite mode.

For repository-level work, the default flow should be:

1. `scan`
2. `analyze`
3. optional `rewrite`
4. optional `compare`

## Target Resolution

Before analyzing, resolve what the user means by "the target skill" or "the target agent".

Possible target forms:

- a file path such as `/repo/skills/foo/SKILL.md`
- a local skill directory containing `SKILL.md`
- a repository root
- a directory subtree within a repository
- a previously installed Claude skill by name
- a pasted skill or agent body

When the target is a local path:

1. Read the `SKILL.md` or instruction file directly.
2. If there are referenced files, only open the ones needed for evaluation.

When the target is a repo or directory:

1. Discover instruction assets automatically.
2. Start with high-confidence path and naming signals.
3. Add medium-confidence candidates only when content signals justify them.
4. Normalize each discovery into `path`, `type`, `confidence`, and `reason_signals`.
5. Analyze the discovered targets and present a scoreboard.

When the target is a named Claude skill:

1. Locate the installed skill metadata and `SKILL.md`.
2. Read only the files needed to understand the target behavior.
3. If the skill cannot be located, say so briefly and continue with the material the user provided.

When the target is ambiguous, ask for the exact skill name or path.

## Example Invocations

Typical Claude Code requests that should trigger this skill:

- `Use skill-accuracy to analyze /repo/skills/frontend/SKILL.md`
- `Use skill-accuracy to evaluate the installed skill "test-driven-development"`
- `Use skill-accuracy to inspect this agent prompt and generate a consistency test plan`
- `Use skill-accuracy to review this documentation-writing skill and suggest stronger wording`
- `Use skill-accuracy to apply the recommended fixes to this skill`
- `Use skill-accuracy to compare the rewritten skill against the original`
- `Use skill-accuracy to scan this repo for skill and agent files and rank them by reliability`
- `Use skill-accuracy to audit this repo and show me the most brittle instruction files`
- `Use skill-accuracy to scan only ./agents and ./skills and summarize the findings`

## Repo Discovery Workflow

When the user asks for a repo audit, prefer a conservative hybrid discovery strategy.

### High-confidence discovery

Start with obvious signals such as:

- `SKILL.md`
- `CLAUDE.md`
- files under `agents/`
- instruction-oriented prompt files
- skill directories with expected structure

### Content-assisted detection

Add medium-confidence candidates only when the file content shows instruction-like behavior such as:

- imperative rules
- stepwise workflows
- tool-use constraints
- output format rules
- explicit must / must-not language

Use [`references/repo-discovery.md`](references/repo-discovery.md) for the detailed heuristics.

### Classification

Classify discovered targets conservatively into:

- `skill`
- `agent`
- `repo-policy`
- `instruction-file`
- `instruction-bundle`

If confidence is too low, keep the item out of the main scoreboard and mention it separately as a possible target.

## Prompt Generation Workflow

When generating test prompts, aim for coverage rather than volume.

Create prompts that probe:

- explicit must-do rules
- explicit must-not-do rules
- formatting constraints
- edge cases and ambiguity
- conflicting temptations
- normal tasks that should succeed easily

Use [`references/prompt-generation.md`](references/prompt-generation.md) when you need the detailed prompt taxonomy.

Default prompt set size:

- 6 to 12 prompts for a light pass
- 12 to 25 prompts for a stronger pass

For a first pass, always prefer a compact set that covers:

- one baseline case
- one temptation case
- one edge case
- one format or policy constraint case

When the target came from repo discovery, tune the prompt set by target type and discovery confidence.

Use [`references/prompt-generation.md`](references/prompt-generation.md) for repo-scan prompt guidance.

## Scoring Workflow

Use a hybrid approach:

1. Rule-based checks first
2. Semantic judgment second, only when needed

Primary metrics:

- `adherence_rate`
- `consistency_rate`
- `format_compliance`
- `failure_pattern_frequency`

Every analysis must also produce a `scorecard`, even in analysis-only mode.

The scorecard should define:

- named metrics
- metric weights
- critical-fail rules
- pass/fail conditions per test prompt
- an overall scoring formula

Default scorecard dimensions:

- `source_selection_score`
- `citation_score`
- `honesty_score`
- `instruction_adherence_score`
- `format_stability_score`
- `consistency_score`

If execution did not happen yet, mark the relevant metrics as `predicted` rather than omitting them.

Use [`references/scoring-rubric.md`](references/scoring-rubric.md) for metric definitions and interpretation.

For repo scan mode, the user-facing report should be a scoreboard, not just a list of notes.

Each scoreboard row should include:

- `overall_score`
- `confidence`
- `type`
- `clarity`
- `testability`
- `consistency_risk`
- `honesty_conflict_risk`
- `format_control`
- `priority`
- concise explanation

At repo level, also include:

- total targets found
- confidence distribution
- highest-risk targets
- strongest targets
- common anti-patterns
- recommended next actions

Use [`references/repo-report-format.md`](references/repo-report-format.md) for the markdown and JSON report contract.

## Improvement Workflow

When a skill or agent is inconsistent, do not just say "the model is flaky".

Try to attribute weakness to one of:

- ambiguous wording
- weak priority structure
- incomplete edge-case coverage
- conflicting instructions
- agent limitation
- interaction effect between wording and agent

Then propose instruction changes that are narrower and more enforceable.

Use [`references/rewrite-heuristics.md`](references/rewrite-heuristics.md) for rewrite patterns.

## Rewrite Policy

When rewriting a target:

- preserve the target's purpose
- prefer small, high-impact edits before broad rewrites
- make rule priority more explicit
- reduce ambiguity
- reduce format variance
- reduce hallucination pressure

When editing a local file:

1. back up the original if the user asked for edits or if the change is meaningful
2. apply the rewrite
3. summarize what changed
4. recommend re-evaluation

If the user does not want the file changed yet, provide a patch proposal instead.

## Compare Policy

When comparing versions:

- name the baseline and candidate clearly
- compare rule quality, not just wording differences
- compare predicted or observed scores side by side
- call out regressions explicitly
- end with a verdict: better, worse, or mixed

Useful outputs:

- metric deltas
- resolved critical issues
- newly introduced risks
- recommended next fixes

## Script Usage

The main workflow should happen inside Claude Code. Use the bundled Go scripts only as optional helpers when repeated execution, case generation, or summarization would be tedious to do manually.

Available scripts:

- `scripts/generate_cases.go`
  - extracts rules and produces candidate test prompts from an instruction file
- `scripts/run_claude_trials.go`
  - runs the generated prompts repeatedly through Claude Code and stores raw outputs
- `scripts/summarize_runs.go`
  - summarizes repeated trial outputs and computes simple consistency signals

Run them with:

```bash
go run ./scripts/generate_cases.go -asset /path/to/SKILL.md
go run ./scripts/run_claude_trials.go -asset /path/to/SKILL.md -cases /path/to/cases.json -runs 10 -output /path/to/runs.json
go run ./scripts/summarize_runs.go -input /path/to/runs.json
```

The scripts are helpers, not the primary interface. The primary interface is Claude Code invoking this skill on another skill, repo, or agent.

## How To Operate Inside Claude Code

When this skill is invoked inside Claude Code, follow this sequence:

1. Resolve the target skill or agent.
2. Determine the requested mode: analyze, repo scan, rewrite, compare, repeated trials, workspace trials, or scan.
3. Inspect the target instructions and extract concrete rules.
4. If the target is a repo, discover and classify candidate instruction assets first.
5. Generate a compact test set when evaluation is needed.
6. If the user asked for analysis-only, stop at critique plus proposed tests and scorecard.
7. If the user asked for rewrite, patch the target or propose a patch.
8. If the user asked for compare, evaluate the versions side by side.
9. If the user asked for execution, perform repeated trials using the most practical available workflow.
10. Summarize adherence, consistency, weak instructions, and rewrite suggestions.

If the user says something like:

- "analyze this skill"
- "rewrite this skill"
- "compare these two versions"
- "measure how consistent this agent is"
- "test this installed Claude skill"

then this skill should treat that as a request to evaluate the target instruction asset, not to solve the target task itself.

If the user asks to fix the target, this skill should remain in evaluation mode while making edits. The goal is still to improve instruction quality, not to act as the target agent.

## Analysis vs Execution

Default to analysis-first.

Choose `analyze` when:

- the user asked for critique or design only
- the target skill is not easily executable
- the execution path is unclear
- running repeated trials would be expensive or noisy

Choose `repo scan` when:

- the user asks for a repository audit
- the user wants automatic discovery of instruction assets
- the user wants a scoreboard across multiple targets

Choose `execution` when:

- the user explicitly wants repeated trials
- the target workflow can be exercised clearly
- the repeated runs are likely to produce diagnostic evidence

Choose `rewrite` when:

- the user explicitly wants changes applied
- the analysis already identified clear high-impact fixes
- the target file is editable and the user wants an updated version

Choose `compare` when:

- there are two candidate versions
- the user asks whether a revision is actually better
- the user wants a regression check after edits

If execution is skipped, still provide:

- extracted rules
- proposed prompt set
- proposed scorecard
- likely weak instructions
- recommended rewrites

If rewrite is skipped, still provide:

- a proposed patch plan
- the highest-impact changes first

If compare is skipped because only one version exists, say so and continue with analyze or rewrite.

## Output Format

When reporting single-target results, structure the answer in this order:

1. `What was tested`
2. `Extracted rules`
3. `Proposed scorecard`
4. `Observed or predicted adherence`
5. `Observed or predicted consistency`
6. `Likely weak instructions`
7. `Recommended rewrites`

If execution did not happen, clearly say it was an analysis-only pass.

For analysis-only passes, explicitly include:

- metric weights
- critical fail conditions
- prompt-by-prompt pass conditions
- predicted weak metrics

For rewrite passes, explicitly include:

- what changed
- why each change matters
- whether a re-evaluation is recommended

For compare passes, explicitly include:

- version labels
- metric deltas
- final verdict

For repo scan passes, explicitly include:

1. `Repo summary`
2. `Discovery overview`
3. `Scoreboard`
4. `Per-target findings`
5. `Cross-repo anti-patterns`
6. `Recommended next actions`

## Constraints

- Do not assume the asset is coding-related.
- Do not require a user-provided test set.
- Prefer a small, representative test set over a large noisy one.
- Separate "skill weakness" from "agent weakness" as an inference, not a certainty.
- If execution tooling is unavailable, still provide prompt generation and evaluation design.
- In repo scan mode, keep discovery conservative and explainable.
- Prefer markdown scoreboard output plus JSON-structured data when the user wants artifacts.
