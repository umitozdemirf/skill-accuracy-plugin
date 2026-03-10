# skill-accuracy

Claude-first plugin for analyzing repository skills, agents, and instruction files with automatic discovery, explainable scoring, and scoreboard-style reports.

## Install

### Claude Code Plugin Marketplace

```bash
/plugin marketplace add umitozdemirf/skill-accuracy-plugin
/plugin install skill-accuracy@umitozdemir
```

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

## What It Does

- discovers instruction assets in a repository
- classifies them conservatively
- analyzes likely weaknesses
- produces a human-readable scoreboard
- summarizes the highest-priority fixes

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

## Package Layout

- `.claude-plugin/marketplace.json`
- `skill-accuracy-plugin/.claude-plugin/plugin.json`
- `skill-accuracy-plugin/skills/skill-accuracy/SKILL.md`
- bundled `references/`, `scripts/`, and `examples/`

## Optional Helper Scripts

The packaged skill includes helper scripts for case generation and repeated-trial summarization:

- `scripts/generate_cases.go`
- `scripts/run_claude_trials.go`
- `scripts/summarize_runs.go`

The primary interface is still Claude invoking the installed `skill-accuracy` plugin.
