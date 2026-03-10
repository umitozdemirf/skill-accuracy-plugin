#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
PLUGIN_DIR="$ROOT_DIR/skill-accuracy-plugin/skills/skill-accuracy"

mkdir -p "$ROOT_DIR/.claude-plugin"
mkdir -p "$ROOT_DIR/skill-accuracy-plugin/.claude-plugin"
mkdir -p "$PLUGIN_DIR"

cp "$ROOT_DIR/SKILL.md" "$PLUGIN_DIR/SKILL.md"

rm -rf "$PLUGIN_DIR/references" "$PLUGIN_DIR/scripts" "$PLUGIN_DIR/examples"
cp -R "$ROOT_DIR/references" "$PLUGIN_DIR/"
cp -R "$ROOT_DIR/scripts" "$PLUGIN_DIR/"
cp -R "$ROOT_DIR/examples" "$PLUGIN_DIR/"

echo "Synced marketplace plugin content into $PLUGIN_DIR"
