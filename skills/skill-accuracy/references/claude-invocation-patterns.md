# Claude Invocation Patterns

This skill is meant to run inside Claude Code and analyze another instruction asset.

## Common User Intents

Interpret these as requests to evaluate or improve a target skill or agent:

- "analyze this skill"
- "rewrite this skill"
- "compare these versions"
- "test this agent"
- "measure consistency of this prompt"
- "review this installed Claude skill"
- "scan this repo for skills"

## Preferred Clarifications

If the target is not obvious, ask for one of:

- exact file path
- exact skill directory
- exact installed skill name

Do not ask for a user-provided test set unless needed. Generate one by default.

## Default Response Shape

Start with:

1. what was resolved as the target
2. requested mode
3. extracted rules
4. proposed prompt set or patch scope
5. whether execution happened or this is analysis-only
6. likely weak instructions
7. rewritten instruction suggestions or compare verdict

## Example Targets

### Local skill file

User:

`Use skill-accuracy to analyze ./skills/frontend/SKILL.md`

Expected behavior:

- read the target skill
- extract enforceable rules
- produce a compact test set
- critique ambiguity and weak wording
- produce a scorecard

### Installed Claude skill

User:

`Use skill-accuracy to evaluate the installed Claude skill "writing-helper"`

Expected behavior:

- locate the installed skill
- inspect `SKILL.md`
- infer what should be tested
- report whether execution is practical

### Rewrite request

User:

`Use skill-accuracy to apply the high-impact fixes to ./skills/frontend/SKILL.md`

Expected behavior:

- analyze first if needed
- back up the original before editing
- apply only the requested scope
- summarize what changed
- recommend re-evaluation

### Compare request

User:

`Use skill-accuracy to compare ./skills/frontend/SKILL.md against ./skills/frontend/SKILL.v2.md`

Expected behavior:

- read both versions
- compare rule quality and score deltas
- identify regressions
- give a verdict

### Agent prompt file

User:

`Use skill-accuracy to inspect ./agents/reviewer.md and tell me why it may be inconsistent`

Expected behavior:

- analyze the prompt structure
- identify conflicting or weak instructions
- suggest tighter rewrites
