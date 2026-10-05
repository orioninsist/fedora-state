# fedora-state

Fedora system state discovery and reproducibility engine.

## Purpose

`fedora-state` discovers the software state of a Fedora system and creates a reproducible model that can be inspected and applied on another Fedora installation.

It answers:

- What is installed on this system?
- Which packages and tools exist?
- How can this environment be recreated?

## Current Status

The core architecture is complete and checkpointed.

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
- error propagation through execution layers

Validation:

```bash
go test ./...
go vet ./...
```

## Installation

### From source

Requirements:

- Fedora Linux
- Go 1.26+

Clone:

```bash
git clone https://github.com/orioninsist/fedora-state.git
cd fedora-state
```

Build the user-local binary:

```bash
mkdir -p ~/.local/bin

go build \
  -o ~/.local/bin/fedora-state \
  ./cmd/fedora-state

chmod 755 ~/.local/bin/fedora-state
```

Add `~/.local/bin` to PATH if needed:

```bash
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

Verify:

```bash
fedora-state --help
```

`~/.local/bin` is used intentionally because it keeps the application inside the user environment and makes rebuilding after a system reinstall simple.

## First Run

Discover the current system:

```bash
fedora-state
```

The discovery command does not modify the system.

Example output:

```text
os=linux architecture=amd64 distribution=Fedora Linux objects=4234 diagnostics=0
```

## Command Usage

The program has one command with different output modes.

### Text report

```bash
fedora-state --format=text
```

### JSON export

```bash
fedora-state --format=json > state.json
```

### Markdown report

```bash
fedora-state --format=markdown > report.md
```

### Manifest generation

```bash
fedora-state --format=manifest > manifest.json
```

### Plan generation

```bash
fedora-state --format=plan > plan.json
```

Always review the plan before applying changes.

### Apply changes

Apply requires administrator privileges because it can call system package operations:

```bash
sudo fedora-state --format=apply
```

## Recommended Workflow

```bash
# 1. Discover current state
fedora-state --format=json > state.json

# 2. Generate and inspect plan
fedora-state --format=plan > plan.json

# 3. Apply only after review
sudo fedora-state --format=apply
```

## Important

`fedora-state` is not a backup tool.

It does not manage:

- dotfiles
- personal configuration
- user data
- accounts

It focuses on reproducible software state.

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
