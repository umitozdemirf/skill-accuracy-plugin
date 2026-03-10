# Skill Accuracy for Codex

Use the `skill-accuracy` skill when you need to evaluate, scan, compare, or improve instruction assets instead of directly solving the target task.

## Where the skill lives

The Codex-ready skill pack is bundled in:

- `codex/skills/skill-accuracy/`

If this repository is vendored into another project, keep that path available or copy the skill into `~/.codex/skills/skill-accuracy/`.

## Recommended usage

- For a fast pass, ask for `quick-analysis`
- For a deeper pass with prompt sets and richer findings, ask for `analyze`
- For repo work, expect the skill to check repo-local instruction assets first and then user-global Codex skills if requested

## Codex surfaces checked first

- `AGENTS.md`
- `SKILL.md`
- `skills/`
- `~/.codex/skills/`

## Example prompts

- `Use skill-accuracy to run a quick-analysis on this project`
- `Use skill-accuracy to analyze my Codex skills under ~/.codex/skills`
- `Use skill-accuracy to compare this skill before and after my edits`
