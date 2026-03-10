# Rewrite Heuristics

When an instruction is weak, prefer rewrites that reduce ambiguity and clarify priority.

## Hard Requirement

Weak:

- Prefer JavaScript

Stronger:

- All code examples must use JavaScript. Do not use TypeScript unless the user explicitly overrides this rule.

## Priority Clarification

Weak:

- Be concise and thorough

Stronger:

- Be concise by default. If thoroughness is required, expand only after covering the direct answer in 3 to 5 sentences.

## Conflict Handling

Weak:

- Ask questions when needed

Stronger:

- If required input is missing, ask a clarifying question before proposing an implementation.

## Format Locking

Weak:

- Return structured output

Stronger:

- Return exactly these sections in order: Summary, Risks, Next Steps.

## Rewrite Rule

Prefer rewrites that make these explicit:

- trigger condition
- required action
- forbidden action
- priority relative to other rules
- allowed override condition
