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

## Package Layout

- `.claude-plugin/marketplace.json`
- `.claude-plugin/plugin.json`
- `skills/skill-accuracy/SKILL.md`
- bundled `references/`, `scripts/`, and `examples/`

## Optional Helper Scripts

The packaged skill includes helper scripts for case generation and repeated-trial summarization:

- `scripts/generate_cases.go`
- `scripts/run_claude_trials.go`
- `scripts/summarize_runs.go`

The primary interface is still Claude invoking the installed `skill-accuracy` plugin.
