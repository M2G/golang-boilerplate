#!/bin/bash

set -e

MODDWATCH_SRC="$(cd "$(dirname "$0")" && pwd)/../../moddwatch"
MODD_VENDOR="$(cd "$(dirname "$0")" && pwd)/../tools/modd/vendor/github.com/cortesi/moddwatch"

echo "Syncing moddwatch → modd vendor..."
echo "  from: $MODDWATCH_SRC"
echo "  to:   $MODD_VENDOR"

FILES=(
    "fswatch_bridge.c"
    "fswatch_bridge.h"
    "fswatch_darwin.go"
    "fswatch_linux.go"
    "watch.go"
    "filter"
)

for f in "${FILES[@]}"; do
    src="$MODDWATCH_SRC/$f"
    dst="$MODD_VENDOR/$f"
    if [ -e "$src" ]; then
        cp -r "$src" "$dst"
        echo "Done $f"
    else
        echo "Fail $f (not found, skipping)"
    fi
done

# Remove old fswatch.go without build tag if present
if [ -f "$MODD_VENDOR/fswatch.go" ]; then
    # Check if it has a build tag
    if ! head -1 "$MODD_VENDOR/fswatch.go" | grep -q "//go:build"; then
        rm "$MODD_VENDOR/fswatch.go"
        echo "done! removed old fswatch.go (no build tag)"
    fi
fi

echo ""
echo "Done. Run 'docker-compose build --no-cache' to rebuild."
