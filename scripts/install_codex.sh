#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
SOURCE_SKILL_DIR="$ROOT_DIR/codex/skills/skill-accuracy"

usage() {
  cat <<'EOF'
Usage:
  bash scripts/install_codex.sh --project /path/to/project
  bash scripts/install_codex.sh --home

Behavior:
  --project  Installs the skill under <project>/tools/skill-accuracy/codex/skills/skill-accuracy
             and creates a minimal AGENTS.md if the project does not already have one.
  --home     Installs the skill under ~/.codex/skills/skill-accuracy
EOF
}

install_project() {
  local project_dir="$1"
  local target_root="$project_dir/tools/skill-accuracy"
  local target_skill_dir="$target_root/codex/skills/skill-accuracy"

  mkdir -p "$target_root/codex/skills"
  rm -rf "$target_skill_dir"
  cp -R "$SOURCE_SKILL_DIR" "$target_root/codex/skills/"

  if [[ ! -f "$project_dir/AGENTS.md" ]]; then
    cp "$ROOT_DIR/AGENTS.md" "$project_dir/AGENTS.md"
  fi

  echo "Installed Codex skill to $target_skill_dir"
}

install_home() {
  local target_dir="$HOME/.codex/skills/skill-accuracy"

  mkdir -p "$HOME/.codex/skills"
  rm -rf "$target_dir"
  cp -R "$SOURCE_SKILL_DIR" "$HOME/.codex/skills/"

  echo "Installed Codex skill to $target_dir"
}

if [[ $# -eq 0 ]]; then
  usage
  exit 1
fi

case "${1:-}" in
  --project)
    if [[ $# -ne 2 ]]; then
      usage
      exit 1
    fi
    install_project "$2"
    ;;
  --home)
    if [[ $# -ne 1 ]]; then
      usage
      exit 1
    fi
    install_home
    ;;
  *)
    usage
    exit 1
    ;;
esac
