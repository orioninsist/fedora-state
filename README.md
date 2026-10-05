# fedora-state

Fedora system state discovery and reproducibility engine.

`fedora-state` discovers the software state of a Fedora machine and produces a reproducible model that can be inspected, exported, and applied on another Fedora installation.

## What it does

It answers:

- What software exists on this system?
- Which packages and tools are installed?
- How can this environment be recreated?

The goal is reproducible system state, not a traditional backup.

## Current Status

The core architecture is implemented and checkpointed.

Implemented:

- system discovery pipeline
- package and tool detection
- manifest generation
- plan generation
- apply workflow
- executor engine abstraction
- DNF backend integration
- dry-run execution path
- text, JSON, and Markdown reporting
- layered error propagation

Validation:

```bash
go test ./...
go vet ./...
```

## Installation

Requirements:

- Fedora Linux
- Go 1.26+

Clone:

```bash
git clone https://github.com/orioninsist/fedora-state.git
cd fedora-state
```

Build:

```bash
mkdir -p ~/.local/bin

go build -o ~/.local/bin/fedora-state ./cmd/fedora-state
chmod 755 ~/.local/bin/fedora-state
```

Verify:

```bash
fedora-state --help
```

Using `~/.local/bin` keeps the binary inside the user environment and makes rebuilding after a reinstall simple.

## Usage

Discover current system state:

```bash
fedora-state
```

The discovery operation does not modify the system.

Example:

```text
os=linux architecture=amd64 distribution=Fedora Linux objects=4234 diagnostics=0
```

## Output Modes

Text report:

```bash
fedora-state --format=text
```

JSON export:

```bash
fedora-state --format=json > state.json
```

Markdown report:

```bash
fedora-state --format=markdown > report.md
```

Manifest:

```bash
fedora-state --format=manifest > manifest.json
```

Plan:

```bash
fedora-state --format=plan > plan.json
```

Always review generated plans before applying changes.

Apply:

```bash
sudo fedora-state --format=apply
```

## Recommended Workflow

```bash
# Discover
fedora-state --format=json > state.json

# Create and review plan
fedora-state --format=plan > plan.json

# Apply after review
sudo fedora-state --format=apply
```

## Scope

`fedora-state` manages software state only.

It does not manage:

- dotfiles
- personal configuration
- user data
- accounts

## Architecture

```text
cmd/fedora-state
        |
        v
commands dispatcher
        |
        v
handlers
        |
        +-- discovery
        +-- manifest
        +-- plan
        +-- apply

internal/executor
        |
        +-- Engine
        |
        +-- Backend
             |
             +-- Real DNF backend
             +-- Dry run backend
```

## Development

Run tests:

```bash
go test ./...
```

Run static analysis:

```bash
go vet ./...
```
