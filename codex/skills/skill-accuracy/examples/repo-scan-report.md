# Repo Summary

- total targets: 4
- confidence distribution: 3 high, 1 medium
- highest-priority targets: `agents/openai.yaml`, `CLAUDE.md`
- strongest targets: `SKILL.md`

# Discovery Overview

| path | type | confidence | reason_signals |
| --- | --- | --- | --- |
| `SKILL.md` | skill | high | `filename=SKILL.md`, `content=workflow rules` |
| `agents/openai.yaml` | agent | high | `directory=agents/`, `content=behavior constraints` |
| `CLAUDE.md` | repo-policy | high | `filename=CLAUDE.md`, `content=repo rules` |
| `docs/assistant-guidance.md` | instruction-file | medium | `content=must/never rules`, `content=format contract` |

# Scoreboard

| path | type | confidence | overall_score | clarity | testability | consistency_risk | honesty_conflict_risk | format_control | priority | notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `SKILL.md` | skill | high | 84 | 86 | 82 | 24 | 18 | 80 | medium | Strong structure, minor ambiguity in edge cases |
| `agents/openai.yaml` | agent | high | 69 | 73 | 64 | 48 | 52 | 71 | high | Persona pressure competes with source-grounding |
| `CLAUDE.md` | repo-policy | high | 72 | 78 | 61 | 37 | 29 | 65 | high | Strong policy coverage, weak testability |
| `docs/assistant-guidance.md` | instruction-file | medium | 58 | 55 | 41 | 57 | 33 | 50 | medium | Looks instruction-like but classification remains tentative |

# Per-Target Findings

## `agents/openai.yaml`

- likely weak areas: honesty pressure, ambiguity under conflicting goals
- next action: tighten source-priority and refusal behavior

## `CLAUDE.md`

- likely weak areas: long-form policy density, unclear operational test points
- next action: extract enforceable rules and reduce overlap

# Cross-Repo Anti-Patterns

- persona language outweighs behavioral constraints
- source-use rules are present but not prioritized
- output format rules are inconsistent across files

# Recommended Next Actions

- tighten the highest-priority agent file first
- separate repo policy from task-specific behavior rules
- add more explicit format constraints where determinism matters
