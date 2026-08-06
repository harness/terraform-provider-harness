#!/bin/bash
#
# Generate provider documentation while preserving subcategories and only syncing changed docs.
#
# Usage: ./scripts/generate-docs.sh
# Env:
#   FORCE_DOCS=true  -> force docs generation regardless of changed file set
#

set -o errexit
set -o nounset
set -o pipefail

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
GENERATED_DOCS_DIR="$BACKUP_DIR/generated-docs"
FORCE_DOCS="${FORCE_DOCS:-false}"
TARGET_DOCS_FILE="$BACKUP_DIR/target-docs.txt"
FULL_REGEN_REQUIRED="false"
DOCS_SCOPE="targeted"

get_changed_files() {
    {
        git diff --name-only
        git diff --name-only --cached
        git ls-files --others --exclude-standard
        git diff --name-only "$(git merge-base HEAD origin/main 2>/dev/null || git merge-base HEAD main 2>/dev/null)" HEAD 2>/dev/null || true
    } | sort -u
}

should_generate_docs() {
    if [ "$FORCE_DOCS" = "true" ]; then
        return 0
    fi

    changed_files="$(get_changed_files)"

    if [ -z "$changed_files" ]; then
        return 1
    fi

    while IFS= read -r file; do
        case "$file" in
            internal/*|templates/*|examples/*|main.go|go.mod|go.sum)
                return 0
                ;;
        esac
    done <<< "$changed_files"

    return 1
}

collect_target_docs() {
    : > "$TARGET_DOCS_FILE"

    changed_files="$(get_changed_files)"

    if [ -z "$changed_files" ]; then
        return 0
    fi

    while IFS= read -r file; do
        [ -z "$file" ] && continue

        case "$file" in
            templates/*|main.go|go.mod|go.sum)
                FULL_REGEN_REQUIRED="true"
                DOCS_SCOPE="full"
                continue
                ;;
            docs/resources/*.md)
                echo "${file#docs/}" >> "$TARGET_DOCS_FILE"
                continue
                ;;
            docs/data-sources/*.md)
                echo "${file#docs/}" >> "$TARGET_DOCS_FILE"
                continue
                ;;
            examples/resources/harness_*.tf)
                resource_name="$(basename "$(dirname "$file")")"
                resource_name="${resource_name#harness_}"
                candidate="resources/${resource_name}.md"
                if [ -f "docs/$candidate" ]; then
                    echo "$candidate" >> "$TARGET_DOCS_FILE"
                fi
                continue
                ;;
            examples/data-sources/harness_*.tf)
                datasource_name="$(basename "$(dirname "$file")")"
                datasource_name="${datasource_name#harness_}"
                candidate="data-sources/${datasource_name}.md"
                if [ -f "docs/$candidate" ]; then
                    echo "$candidate" >> "$TARGET_DOCS_FILE"
                fi
                continue
                ;;
            internal/*.go|internal/*/*.go|internal/*/*/*.go|internal/*/*/*/*.go|internal/*/*/*/*/*.go)
                base="$(basename "$file" .go)"
                stem=""
                doc_type=""

                case "$base" in
                    resource_*)
                        stem="${base#resource_}"
                        doc_type="resources"
                        ;;
                    *_resource)
                        stem="${base%_resource}"
                        doc_type="resources"
                        ;;
                    data_source_*)
                        stem="${base#data_source_}"
                        doc_type="data-sources"
                        ;;
                    *_data_source)
                        stem="${base%_data_source}"
                        doc_type="data-sources"
                        ;;
                esac

                if [ -n "$stem" ] && [ -n "$doc_type" ]; then
                    stem_lower="$(printf '%s' "$stem" | tr '[:upper:]' '[:lower:]')"
                    for candidate in docs/"$doc_type"/*"${stem_lower}".md; do
                        if [ -f "$candidate" ]; then
                            echo "${candidate#docs/}" >> "$TARGET_DOCS_FILE"
                        fi
                    done
                fi
                continue
                ;;
        esac
    done <<< "$changed_files"

    if [ -f "$TARGET_DOCS_FILE" ]; then
        sort -u "$TARGET_DOCS_FILE" -o "$TARGET_DOCS_FILE"
    fi

    if [ "$FULL_REGEN_REQUIRED" = "true" ]; then
        info "Docs scope: full (global doc-impacting files changed)"
    elif [ -s "$TARGET_DOCS_FILE" ]; then
        target_count="$(wc -l < "$TARGET_DOCS_FILE" | tr -d '[:space:]')"
        info "Docs scope: targeted ($target_count files)"
    else
        info "Docs scope: targeted (0 mapped docs)"
    fi
}

sync_generated_docs() {
    local changed_count=0
    local created_count=0
    local removed_count=0

    info "Applying generated docs sync ($DOCS_SCOPE)..."

    process_doc_file() {
        local rel_file="$1"
        [ -z "$rel_file" ] && return 0
        local src="$GENERATED_DOCS_DIR/$rel_file"
        local dst="docs/$rel_file"
        local dst_dir
        dst_dir="$(dirname "$dst")"
        [ ! -f "$src" ] && return 0
        mkdir -p "$dst_dir"

        if [ ! -f "$dst" ]; then
            cp "$src" "$dst"
            created_count=$((created_count + 1))
            continue
        fi

        if ! cmp -s "$src" "$dst"; then
            cp "$src" "$dst"
            changed_count=$((changed_count + 1))
        fi
    }

    if [ "$FULL_REGEN_REQUIRED" = "true" ]; then
        while IFS= read -r rel_file; do
            process_doc_file "$rel_file"
        done < <(cd "$GENERATED_DOCS_DIR" && find . -type f -name "*.md" | sed 's|^\./||')
    elif [ -s "$TARGET_DOCS_FILE" ]; then
        while IFS= read -r rel_file; do
            process_doc_file "$rel_file"
        done < "$TARGET_DOCS_FILE"
    else
        warn "No target docs resolved from changed files; skipping docs sync"
        return 0
    fi

    if [ "$FULL_REGEN_REQUIRED" = "true" ]; then
        for dir in docs/resources docs/data-sources; do
            if [ -d "$dir" ]; then
                while IFS= read -r dst_file; do
                    rel_path="${dst_file#docs/}"
                    if [ ! -f "$GENERATED_DOCS_DIR/$rel_path" ]; then
                        rm -f "$dst_file"
                        removed_count=$((removed_count + 1))
                    fi
                done < <(find "$dir" -type f -name "*.md")
            fi
        done
    fi

    info "Docs sync summary: created=$created_count, changed=$changed_count, removed=$removed_count"
}

if ! should_generate_docs; then
    success "No doc-impacting file changes detected; skipping docs generation"
    exit 0
fi

# Clean up any previous backup
rm -rf "$BACKUP_DIR"
mkdir -p "$BACKUP_DIR"
collect_target_docs

# Save existing subcategories
info "Saving existing subcategories..."
subcategory_count=0

for dir in docs/resources docs/data-sources; do
    if [ -d "$dir" ]; then
        for f in "$dir"/*.md; do
            if [ -f "$f" ]; then
                subcat="$(grep -m1 '^subcategory:' "$f" 2>/dev/null | sed 's/^subcategory: *//' || echo "")"
                if [ -n "$subcat" ] && [ "$subcat" != '""' ]; then
                    echo "$f|$subcat" >> "$SUBCATEGORIES_FILE"
                    subcategory_count=$((subcategory_count + 1))
                fi
            fi
        done
    fi
done

info "Saved $subcategory_count subcategories"

# Generate docs into a temp directory first, then sync only changed files back.
info "Generating Terraform provider docs into temporary output..."
"$TFPLUGINDOCS_PATH" generate --rendered-website-dir "$GENERATED_DOCS_DIR"

# Restore subcategories in generated docs
if [ -f "$SUBCATEGORIES_FILE" ]; then
    info "Restoring subcategories..."
    restored_count=0

    while IFS='|' read -r file subcat; do
        generated_file="${file#docs/}"
        generated_file="$GENERATED_DOCS_DIR/$generated_file"

        if [ -f "$generated_file" ]; then
            current="$(grep -m1 '^subcategory:' "$generated_file" 2>/dev/null | sed 's/^subcategory: *//' || echo "")"
            if [ "$current" = '""' ]; then
                if [[ "$OSTYPE" == "darwin"* ]]; then
                    sed -i '' "s/^subcategory: \"\"$/subcategory: $subcat/" "$generated_file"
                else
                    sed -i "s/^subcategory: \"\"$/subcategory: $subcat/" "$generated_file"
                fi
                restored_count=$((restored_count + 1))
            fi
        fi
    done < "$SUBCATEGORIES_FILE"

    info "Restored $restored_count subcategories"
fi

sync_generated_docs

# Clean up
rm -rf "$BACKUP_DIR"

success "Documentation generation complete"
