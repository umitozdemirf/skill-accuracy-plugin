# Repo Discovery

Use this reference when `skill-accuracy` needs to audit a repository or skill root instead of a single file.

## Goal

Find likely instruction assets with enough confidence that the resulting scoreboard remains useful and explainable.

## Quick Analysis First Pass

When the caller explicitly asks for `quick-analysis` in a Claude Code repository, check standard instruction surfaces before generic discovery:

- `CLAUDE.md`
- `.claude/commands/`
- `.claude/agents/`
- `skills/`

If these paths exist, use them as the target set and stop there.
Only fall back to the broader discovery tiers below when none of these surfaces exist.

When the caller explicitly asks for `quick-analysis` in a Codex environment, check these standard surfaces first:

- `AGENTS.md`
- `SKILL.md`
- `skills/`
- `~/.codex/skills/`
- limited packaged-skill paths under `~/.codex/*/skills/`

For Codex, preserve source scope:

- `scope=repo` for repo-local assets
- `scope=user` for user-global assets

## Discovery Tiers

## Platform-Specific Discovery

### Claude

Treat these as high-confidence direct targets unless there is strong evidence otherwise:

- `CLAUDE.md`
- `.claude/commands/`
- `.claude/agents/`
- `skills/`

### Codex

Treat these as high-confidence direct targets unless there is strong evidence otherwise:

- `AGENTS.md`
- `SKILL.md`
- `skills/`
- `~/.codex/skills/`
- `~/.codex/*/skills/`

## Discovery Tiers

### High-confidence discovery

Treat these as direct targets unless there is strong evidence otherwise:

- `SKILL.md`
- `CLAUDE.md`
- `AGENTS.md`
- files under `agents/`
- files under `.claude/agents/`
- files under `.claude/commands/`
- files under directories clearly dedicated to prompts or instructions
- files whose names explicitly include `prompt`, `instruction`, `agent`, or `policy`

## Content-assisted detection

Promote a file to medium-confidence discovery when the content includes multiple instruction-like signals such as:

- repeated imperative verbs
- explicit `must`, `must not`, `always`, `never`, `prefer`, or `only`
- step-by-step operational workflows
- tool usage constraints
- output formatting contracts
- conflict resolution or priority rules

Use content-assisted detection to widen discovery carefully, not to vacuum every markdown file in the repo.

## Confidence Rules

- `high`: path and naming are explicit, or content strongly confirms the role
- `medium`: naming is weak but content shows clear instruction behavior
- `low`: some signals exist but the file may be ordinary documentation

Low-confidence candidates should not be forced into the main scoreboard. Surface them separately as possible targets.

## Normalized Target Shape

Each discovery should be converted into:

- `path`
- `type`
- `confidence`
- `reason_signals`
- `platform`
- `scope`

`reason_signals` should be concrete, for example:

- `filename=SKILL.md`
- `directory=agents/`
- `directory=.claude/agents/`
- `directory=~/.codex/skills/`
- `content=must/never rules`
- `content=tool usage constraints`

## Classification Mapping

Use these initial target classes:

- `skill`
- `agent`
- `repo-policy`
- `instruction-file`
- `instruction-bundle`

Suggested mapping rules:

- `SKILL.md` or a skill directory root -> `skill`
- `~/.codex/skills/<name>/SKILL.md` -> `skill`
- files under `agents/` -> `agent`
- files under `.claude/agents/` -> `agent`
- `CLAUDE.md` -> `repo-policy`
- `AGENTS.md` -> `repo-policy`
- single prompt or instruction asset -> `instruction-file`
- grouped instruction set with related references -> `instruction-bundle`

If classification is ambiguous, prefer the narrower label and explain why.
