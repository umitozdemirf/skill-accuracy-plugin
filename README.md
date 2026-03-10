# skill-accuracy

Claude-first plugin with Codex support for analyzing repository skills, agents, and instruction files with automatic discovery, explainable scoring, and scoreboard-style reports.

## Install

### Claude Code Plugin Marketplace

```bash
/plugin marketplace add umitozdemirf/skill-accuracy-plugin
/plugin install skill-accuracy@umitozdemirf
```

If the skill does not appear immediately after install, start a new Claude session or restart Claude. Some sessions do not reload installed plugins dynamically.

### Team Rollout

Add this to project-level `.claude/settings.json`:

```json
{
  "extraKnownMarketplaces": {
    "skill-accuracy": {
        "source": {
          "source": "github",
          "repo": "umitozdemirf/skill-accuracy-plugin"
        }
      }
    }
}
```

### Codex

Codex works best through the repository-level [AGENTS.md](/Users/umitozdemir/proj/skill-analysis/skill-accuracy-plugin/AGENTS.md) plus the bundled Codex skill tree under [codex/skills](/Users/umitozdemir/proj/skill-analysis/skill-accuracy-plugin/codex/skills).

Recommended install targets:

- project-local vendor path
- `~/.codex/skills/skill-accuracy/`

Minimal install flow:

```bash
bash scripts/install_codex.sh --project "$(pwd)"
```

## What It Does

- discovers instruction assets in a repository
- classifies them conservatively
- analyzes likely weaknesses
- produces a human-readable scoreboard
- summarizes the highest-priority fixes
- supports both Claude-style repo assets and Codex repo/user skill surfaces

## Usage Examples

### Load the skill

```text
/skill-accuracy:skill-accuracy
```

Expected interaction:

```text
The skill-accuracy skill is loaded. I can help you evaluate, scan, improve, or compare instruction assets.

Available modes:
1. Analyze
2. Repo Scan
3. Quick Analysis
4. Rewrite
5. Compare
6. Repeated Trials
7. Workspace Trials
```

### Full analysis

```text
analyze in this project
```

Expected behavior:

- scans the repo for instruction assets
- selects the primary scoreable targets
- produces a full analysis with scorecards, weak instructions, and test prompts

Use this when you want depth over speed.

### Quick analysis

```text
run a quick-analysis in this project
```

Expected behavior:

- checks standard Claude Code instruction surfaces first:
  - `CLAUDE.md`
  - `.claude/commands/`
  - `.claude/agents/`
  - `skills/`
- avoids broad discovery if those paths exist
- returns:
  - `Quick Scoreboard`
  - `Short Findings`
  - `Next Actions`

Use this when you want a faster structured pass.

### Codex quick analysis

```text
run a quick-analysis on this repo and my Codex skills
```

Expected behavior:

- checks repo-local Codex surfaces first:
  - `AGENTS.md`
  - `SKILL.md`
  - `skills/`
- also checks user-global Codex skills:
  - `~/.codex/skills/`
- groups results into:
  - `Repo Targets`
  - `User Skill Targets`

Use this when you want Codex-aware discovery instead of Claude-only repo scanning.

## Package Layout

- `.claude-plugin/marketplace.json`
- `skill-accuracy-plugin/.claude-plugin/plugin.json`
- `skill-accuracy-plugin/skills/skill-accuracy/SKILL.md`
- `codex/skills/skill-accuracy/SKILL.md`
- `AGENTS.md`
- bundled `references/`, `scripts/`, and `examples/`

## Optional Helper Scripts

The packaged skill includes helper scripts for case generation and repeated-trial summarization:

- `scripts/generate_cases.go`
- `scripts/run_claude_trials.go`
- `scripts/summarize_runs.go`

Repository-level helper scripts:

- `scripts/install_codex.sh`

The primary interfaces are Claude marketplace install and Codex `AGENTS.md` + skill-tree integration.
