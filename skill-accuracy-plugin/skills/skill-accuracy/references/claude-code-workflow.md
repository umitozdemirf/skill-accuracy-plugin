# Claude Code Workflow

Use Claude Code as the first execution target and the primary editing surface for this skill.

## Execution Pattern

Treat the target skill or instruction file as extra system guidance.

Recommended flow:

1. Read the target asset.
2. Determine whether the request is analyze, rewrite, compare, or execute.
3. Generate a compact test case set when evaluation is needed.
4. Run the same prompt repeatedly through Claude Code when execution is requested.
5. Store raw outputs if running trials.
6. Summarize consistency and failure patterns.
7. If rewrite was requested, patch the target and recommend re-evaluation.

## Practical Invocation

The helper script uses Claude Code in non-interactive mode.

Assumed invocation shape:

```bash
claude --print --append-system-prompt "..." "task prompt"
```

Prerequisite:

- Claude Code must be installed
- Claude Code must already be logged in

Quick local checks:

```bash
claude --version
claude --print "Say hello in one sentence."
```

If Claude Code returns `Not logged in · Please run /login`, authenticate first and retry.

Optional flags can be passed through when needed, such as:

- current working directory
- allowed tools
- model

## What To Look For

When validating Claude behavior, focus on:

- whether hard rules survive repeated runs
- whether the same prompt yields different decisions
- whether format constraints drift
- whether failure patterns repeat on temptation prompts
- whether the rewrite actually improved predicted or observed scores

## Caution

A repeated failure on Claude Code does not automatically mean the skill is bad.
It may still be:

- ambiguous wording
- weak prioritization
- an execution/tooling limitation
- a prompt that is not diagnostic enough

## Typical Lifecycle

Common real usage:

1. `analyze` a target skill
2. `rewrite` the target using the highest-impact fixes
3. `compare` the rewritten version against the original
4. optionally `execute` repeated trials for extra confidence
