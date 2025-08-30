#!/usr/bin/env bash
set -euo pipefail

git fetch origin main --depth=1 || true
CHANGED_FILES=$(git diff --name-only origin/main...HEAD || true)

echo "Changed files:"
echo "$CHANGED_FILES"

# If code changed, require docs update
if echo "$CHANGED_FILES" | grep -E '\.go$|\.js$|\.html$' >/dev/null; then
  # docs updated?
  if echo "$CHANGED_FILES" | grep -E '^docs/features/|^docs/|README.md' >/dev/null; then
    # validate new/changed docs have required header lines
    DOC_FILES=$(echo "$CHANGED_FILES" | grep -E '^docs/.*\.md$' || true)
    for f in $DOC_FILES; do
      # skip if file deleted
      if [ ! -f "$f" ]; then
        continue
      fi
      head -n 5 "$f" | grep -q '^Filepath:' || { echo "ERROR: $f missing 'Filepath:' header"; exit 1; }
      head -n 6 "$f" | grep -q '^Purpose:' || { echo "ERROR: $f missing 'Purpose:' header"; exit 1; }
    done
    echo "Documentation updated and headers validated. OK."
    exit 0
  else
    echo "ERROR: Code changes detected but no docs updated under docs/ or README.md."
    exit 1
  fi
else
  echo "No relevant code changes detected. Docs not required."
  exit 0
fi