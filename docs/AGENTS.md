# AGENTS.md - Documentation

## Purpose

This directory contains auto-generated Terraform provider documentation. The documentation is generated from resource and data source schema descriptions using `tfplugindocs`.

## Directory Structure

```
docs/
├── index.md              # Provider overview and configuration
├── guides/               # User guides and tutorials
├── resources/            # Resource documentation (auto-generated)
│   ├── platform_connector_*.md
│   ├── platform_organization.md
│   ├── platform_project.md
│   └── ...
└── data-sources/         # Data source documentation (auto-generated)
    ├── platform_connector_*.md
    └── ...
```

## Regenerating Documentation

Run from the repository root:
```sh
make docs
```

This executes `scripts/generate-docs.sh` which:
1. Runs `tfplugindocs generate`
2. Preserves subcategory frontmatter in existing docs
3. Formats the generated markdown

## Documentation Templates

Templates for documentation are in `templates/`:
- Resource templates define the structure of resource docs
- Data source templates define the structure of data source docs

## Frontmatter

Each documentation file has YAML frontmatter:
```yaml
---
page_title: "harness_platform_connector_github Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating a GitHub connector.
---
```

The `subcategory` field determines navigation grouping on the Terraform Registry.

## Schema Descriptions

Documentation is generated from schema descriptions in resource code:
```go
Schema: map[string]*schema.Schema{
    "identifier": {
        Description: "Unique identifier of the resource.",  // This becomes docs
        Type:        schema.TypeString,
        Required:    true,
    },
}
```

## Validation

Validate documentation:
```sh
make docs-validate
```

## DOs

- Always regenerate docs after modifying resource schemas
- Write clear, user-friendly descriptions in schema definitions
- Use proper markdown formatting in descriptions
- Keep frontmatter `subcategory` consistent

## DON'Ts

- Don't manually edit files in `docs/resources/` or `docs/data-sources/` - they will be overwritten
- Don't commit docs changes without running `make docs` first
- Don't remove frontmatter from generated files

## Updating index.md and guides/

The `docs/index.md` and files in `docs/guides/` are NOT auto-generated and can be manually edited. These provide:
- Provider configuration examples
- Authentication guidance
- Usage tutorials
