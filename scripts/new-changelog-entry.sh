#!/bin/bash
#
# Scaffold a new .changelog entry for the current branch/PR.
#
# Usage: ./scripts/new-changelog-entry.sh [name] [type]
#   name: Optional explicit entry name (used as .changelog/<name>.txt).
#         If omitted, the script tries, in order:
#           1. A JIRA-style ticket in the current branch name (e.g. feature/CDS-130810-foo -> CDS-130810)
#           2. The open GitHub PR number for the current branch (requires `gh`)
#   type: Optional release-note type (enhancement|bug|feature|new-resource|new-data-source|breaking-change|note)
#         Defaults to "enhancement".
#
# Examples:
#   ./scripts/new-changelog-entry.sh                       # auto-detect from branch/PR
#   ./scripts/new-changelog-entry.sh CDS-130810 bug         # explicit name + type
#

set -o errexit
set -o nounset

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'
BOLD='\033[1m'

info() { echo -e "${BLUE}→${NC} $1"; }
success() { echo -e "${GREEN}✓${NC} $1"; }
warn() { echo -e "${YELLOW}⚠${NC} $1"; }
error() { echo -e "${RED}✗${NC} $1"; }

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
__parent="$(dirname "$__dir")"
CHANGELOG_DIR="$__parent/.changelog"

NAME="${1:-}"
TYPE="${2:-enhancement}"

VALID_TYPES="enhancement bug feature new-resource new-data-source breaking-change note"
if ! echo " $VALID_TYPES " | grep -q " $TYPE "; then
    error "Invalid type '$TYPE'. Must be one of: $VALID_TYPES"
    exit 1
fi

if [ -z "$NAME" ]; then
    BRANCH=$(git rev-parse --abbrev-ref HEAD 2>/dev/null || echo "")

    # Look for a JIRA-style ticket (e.g. CDS-130810, PL-72902) anywhere in the branch name
    TICKET=$(echo "$BRANCH" | grep -Eio '[A-Za-z]+-[0-9]+' | head -n1 || true)

    if [ -n "$TICKET" ]; then
        NAME=$(echo "$TICKET" | tr '[:lower:]' '[:upper:]')
        info "Detected ticket ${BOLD}$NAME${NC} from branch '$BRANCH'"
    elif command -v gh >/dev/null 2>&1; then
        PR_NUMBER=$(gh pr view --json number --jq '.number' 2>/dev/null || true)
        if [ -n "$PR_NUMBER" ]; then
            NAME="$PR_NUMBER"
            info "Detected PR ${BOLD}#$NAME${NC} for branch '$BRANCH'"
        fi
    fi
fi

if [ -z "$NAME" ]; then
    error "Could not determine a changelog entry name."
    info "No JIRA ticket found in branch name and no open PR found for this branch."
    info "Pass one explicitly: make changelog-entry ENTRY=CDS-12345 [TYPE=bug]"
    exit 1
fi

ENTRY_FILE="$CHANGELOG_DIR/$NAME.txt"

if [ -f "$ENTRY_FILE" ]; then
    warn "Changelog entry already exists: .changelog/$NAME.txt"
    info "Edit it directly, or pass a different NAME to create another one."
    exit 0
fi

mkdir -p "$CHANGELOG_DIR"

cat > "$ENTRY_FILE" <<EOF
\`\`\`release-note:$TYPE
resource/harness_platform_example: Describe the user-facing change here
\`\`\`
EOF

success "Created .changelog/$NAME.txt"
info "Edit it to describe the change, then run 'make changelog' to regenerate CHANGELOG.md"
