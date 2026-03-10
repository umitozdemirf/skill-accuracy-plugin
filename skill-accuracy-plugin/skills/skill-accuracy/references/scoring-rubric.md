# Scoring Rubric

Use lightweight metrics first.

In repo-scan mode, treat `overall_score`, `priority`, and `confidence` as separate signals:

- `overall_score` estimates instruction quality
- `priority` estimates what needs attention first
- `confidence` estimates how trustworthy the classification and discovery are

## Adherence Rate

Fraction of runs that satisfy the targeted instruction.

Examples:

- used JavaScript when required
- stayed under the required length
- asked a clarifying question first

## Consistency Rate

Fraction of runs that preserve the same essential behavior under repeated trials.

Consistency is not exact string matching only. Evaluate:

- same decision
- same constraint compliance
- same output structure

## Format Compliance

Whether the output shape is stable enough to be usable.

Examples:

- valid JSON
- correct headings
- bullet list preserved

## Failure Pattern Frequency

Count repeated failure modes instead of isolated misses.

Examples:

- often ignores the "ask first" rule
- occasionally drifts into TypeScript
- becomes verbose on ambiguous prompts

## Interpretation

- high adherence + low consistency: wording may be fragile
- low adherence across strong agents: instruction likely weak
- high variance only on one agent: likely agent limitation

## Scorecard Template

Every evaluation should define a scorecard.

Recommended template:

- `source_selection_score`: 20%
- `citation_score`: 20%
- `honesty_score`: 25%
- `instruction_adherence_score`: 20%
- `format_stability_score`: 5%
- `consistency_score`: 10%

Example critical-fail rules:

- fabricated unsupported facts
- ignored a hard language constraint
- answered without using any relevant source when source use is mandatory

For each prompt, define:

- targeted rules
- pass condition
- fail condition
- weight or severity

## Repo Scan Dimensions

For repository audits, expose a scoreboard with these dimensions:

- `overall_score`
- `clarity`
- `testability`
- `consistency_risk`
- `honesty_conflict_risk`
- `format_control`
- `priority`
- `confidence`

Interpretation guidance:

- high `clarity` + low `testability`: the instructions read cleanly but are hard to verify
- low `clarity` + high `consistency_risk`: wording is likely brittle
- high `honesty_conflict_risk`: persona or output pressure may incentivize fabricated claims
- high `priority`: the file is a strong candidate for the next rewrite pass

`priority` should not be a copy of `overall_score`. Derive it from a combination of low instruction quality, apparent repo importance, and risk severity.
