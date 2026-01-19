#!/bin/bash
#
# Generate provider documentation while preserving subcategories
#
# Usage: ./scripts/generate-docs.sh
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

cd "$__parent"

# Check if tfplugindocs is installed
TFPLUGINDOCS_PATH="$(go env GOPATH)/bin/tfplugindocs"
if [ ! -f "$TFPLUGINDOCS_PATH" ]; then
    error "tfplugindocs not found at $TFPLUGINDOCS_PATH"
    echo ""
    info "Install it with:"
    echo "  go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@latest"
    echo "  or run: make tools"
    echo ""
    exit 1
fi

BACKUP_DIR=".docs-backup"
SUBCATEGORIES_FILE="$BACKUP_DIR/subcategories.txt"

# Clean up any previous backup
rm -rf "$BACKUP_DIR"
mkdir -p "$BACKUP_DIR"

# Save existing subcategories
info "Saving existing subcategories..."
subcategory_count=0

for dir in docs/resources docs/data-sources; do
    if [ -d "$dir" ]; then
        for f in "$dir"/*.md; do
            if [ -f "$f" ]; then
                # Extract subcategory value (with quotes)
                subcat=$(grep -m1 '^subcategory:' "$f" 2>/dev/null | sed 's/^subcategory: *//' || echo "")
                if [ -n "$subcat" ] && [ "$subcat" != '""' ]; then
                    echo "$f|$subcat" >> "$SUBCATEGORIES_FILE"
                    subcategory_count=$((subcategory_count + 1))
                fi
            fi
        done
    fi
done

info "Saved $subcategory_count subcategories"

# Generate documentation
info "Generating Terraform provider docs..."
"$TFPLUGINDOCS_PATH" generate

# Restore subcategories
if [ -f "$SUBCATEGORIES_FILE" ]; then
    info "Restoring subcategories..."
    restored_count=0

    while IFS='|' read -r file subcat; do
        if [ -f "$file" ]; then
            # Check if subcategory was cleared (is now "")
            current=$(grep -m1 '^subcategory:' "$file" 2>/dev/null | sed 's/^subcategory: *//' || echo "")
            if [ "$current" = '""' ]; then
                # Restore the original subcategory
                if [[ "$OSTYPE" == "darwin"* ]]; then
                    sed -i '' "s/^subcategory: \"\"$/subcategory: $subcat/" "$file"
                else
                    sed -i "s/^subcategory: \"\"$/subcategory: $subcat/" "$file"
                fi
                restored_count=$((restored_count + 1))
            fi
        fi
    done < "$SUBCATEGORIES_FILE"

    info "Restored $restored_count subcategories"
fi

# Clean up
rm -rf "$BACKUP_DIR"

success "Documentation generated in docs/ with subcategories preserved"
