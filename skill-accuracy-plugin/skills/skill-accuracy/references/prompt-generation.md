# Prompt Generation

Generate prompts that test whether the instruction asset actually controls behavior.

Use these categories:

## Baseline

Straightforward tasks that should succeed if the instruction is functioning.

## Constraint Probes

Prompts that directly test hard rules such as:

- required language
- required format
- brevity
- must ask before acting
- must avoid a prohibited pattern

## Temptation Cases

Prompts that make breaking the rule attractive.

Examples:

- ask for TypeScript when the skill requires JavaScript
- ask for a long explanation when the skill requires brevity
- ask for direct action when the skill requires clarifying questions

## Edge Cases

Prompts with ambiguity, missing information, or partially conflicting goals.

## Domain Transfer

Prompts that move the same rule into a different domain.

Example:

- from code generation to documentation
- from frontend tasks to API tasks

## Suggested Output

For each prompt, capture:

- `id`
- `intent`
- `prompt`
- `targeted_rules`
- `why_it_matters`

## Repo Scan Guidance

When prompts are generated from discovered repo targets, adapt the set by target type:

- `skill`: emphasize workflow rules, tool use, output format, and refusal behavior
- `agent`: emphasize persona pressure, source use, citation rules, and scope boundaries
- `repo-policy`: emphasize policy precedence, escalation behavior, and defaults
- `instruction-file`: emphasize explicit must / must-not rules

Also adapt by discovery confidence:

- `high`: generate a normal compact prompt set
- `medium`: prefer prompts that validate whether the file is truly instruction-bearing

For repo scans, keep prompt volume low and coverage broad so the scoreboard remains readable.
