#!/bin/bash
#
# Generate changelog from .changelog entries
# Requires: github.com/hashicorp/go-changelog/cmd/changelog-build
#
# Usage: ./scripts/generate-changelog.sh [version]
#   version: Optional version string (e.g., "0.40.0"). If not provided, uses next minor version.
#

set -o errexit
set -o nounset

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color
BOLD='\033[1m'

info() { echo -e "${BLUE}→${NC} $1"; }
success() { echo -e "${GREEN}✓${NC} $1"; }
warn() { echo -e "${YELLOW}⚠${NC} $1"; }
error() { echo -e "${RED}✗${NC} $1"; }

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
__parent="$(dirname "$__dir")"

CHANGELOG_FILE_NAME="CHANGELOG.md"
CHANGELOG_TMP_FILE_NAME="CHANGELOG.tmp"

# Check if changelog-build is installed
CHANGELOG_BUILD_PATH="$(go env GOPATH)/bin/changelog-build"
if [ ! -f "$CHANGELOG_BUILD_PATH" ]; then
    error "changelog-build not found at $CHANGELOG_BUILD_PATH"
    echo ""
    info "Install it with:"
    echo "  go install github.com/hashicorp/go-changelog/cmd/changelog-build@latest"
    echo ""
    exit 1
fi

TARGET_SHA=$(git rev-parse HEAD)
PREVIOUS_RELEASE_TAG=$(git describe --abbrev=0 --match='v*.*.*' --tags 2>/dev/null || git describe --abbrev=0 --match='*.*.*' --tags 2>/dev/null || echo "")

if [ -z "$PREVIOUS_RELEASE_TAG" ]; then
    error "No previous release tag found"
    exit 1
fi

PREVIOUS_RELEASE_SHA=$(git rev-list -n 1 "$PREVIOUS_RELEASE_TAG")

info "Previous release: ${BOLD}$PREVIOUS_RELEASE_TAG${NC} ($PREVIOUS_RELEASE_SHA)"
info "Current HEAD: $TARGET_SHA"

if [ "$TARGET_SHA" == "$PREVIOUS_RELEASE_SHA" ]; then
    warn "HEAD is at the previous release tag. Nothing to do."
    exit 0
fi

# Extract version without 'v' prefix for searching
VERSION_WITHOUT_V="${PREVIOUS_RELEASE_TAG#v}"

# Try to find the previous changelog section - handle both with and without 'v' prefix
PREVIOUS_CHANGELOG=$(sed -n -e "/^# v\{0,1\}${VERSION_WITHOUT_V}/,\$p" "$__parent/$CHANGELOG_FILE_NAME" 2>/dev/null || echo "")

if [ -z "$PREVIOUS_CHANGELOG" ]; then
    warn "Could not find version $VERSION_WITHOUT_V in changelog. Using entire file as base."
    PREVIOUS_CHANGELOG=$(cat "$__parent/$CHANGELOG_FILE_NAME")
fi

info "Generating changelog from .changelog entries..."

# Create a temp file for error output
ERROR_FILE=$(mktemp)
trap "rm -f $ERROR_FILE" EXIT

# Run changelog-build, capturing stdout and stderr separately
CHANGELOG=$("$CHANGELOG_BUILD_PATH" -this-release "$TARGET_SHA" \
                      -last-release "$PREVIOUS_RELEASE_SHA" \
                      -git-dir "$__parent" \
                      -entries-dir .changelog \
                      -changelog-template "$__dir/changelog.tmpl" \
                      -note-template "$__dir/release-note.tmpl" \
                      -local-fs 2>"$ERROR_FILE") || true

# Check for errors
if [ -s "$ERROR_FILE" ]; then
    ERROR_MSG=$(cat "$ERROR_FILE")
    if echo "$ERROR_MSG" | grep -q "unstaged changes"; then
        error "Repository has uncommitted changes. Please commit or stash changes first."
        info "Run: git status"
        exit 1
    elif echo "$ERROR_MSG" | grep -q "error\|Error\|ERROR"; then
        error "changelog-build failed:"
        cat "$ERROR_FILE"
        exit 1
    fi
fi

if [ -z "$CHANGELOG" ]; then
    warn "No new changelog entries found in .changelog directory"
    info "To add entries, create files in .changelog/ with format:"
    echo "  .changelog/<PR_NUMBER>.txt"
    echo ""
    echo "  Example content:"
    echo "  \`\`\`release-note:enhancement"
    echo "  resource/harness_platform_connector: Added new feature X"
    echo "  \`\`\`"
    exit 0
fi

# Get version for the new entry
if [ -n "${1:-}" ]; then
    NEW_VERSION="$1"
else
    # Auto-increment patch version
    IFS='.' read -r major minor patch <<< "$VERSION_WITHOUT_V"
    NEW_VERSION="${major}.${minor}.$((patch + 1))"
fi

NEW_DATE=$(date "+%B %d, %Y")

info "Creating changelog entry for version ${BOLD}$NEW_VERSION${NC}"

rm -f "$CHANGELOG_TMP_FILE_NAME"

# Write new changelog
{
    echo "# v$NEW_VERSION ($NEW_DATE)"
    echo ""
    echo "$CHANGELOG"
    echo ""
    echo "$PREVIOUS_CHANGELOG"
} > "$CHANGELOG_TMP_FILE_NAME"

cp "$CHANGELOG_TMP_FILE_NAME" "$__parent/$CHANGELOG_FILE_NAME"
rm -f "$CHANGELOG_TMP_FILE_NAME"

success "Successfully generated changelog for v$NEW_VERSION"
info "Review changes in $CHANGELOG_FILE_NAME"

exit 0
