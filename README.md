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
- text and JSON reporting
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

## First Run

During development, run directly with Go:

```bash
go run ./cmd/fedora-state
```

Normal installed usage:

```bash
fedora-state
```

The first command only discovers your current system. It does not modify anything.

Example output:

```text
os=linux architecture=amd64 distribution=Fedora Linux objects=4233 diagnostics=0
```

## Command Usage

The program has one command with different output modes.

### Text report

```bash
go run ./cmd/fedora-state --format=text
```

or after installation:

```bash
fedora-state --format=text
```

### JSON export

Create a machine-readable snapshot:

```bash
go run ./cmd/fedora-state --format=json > state.json
```

### Manifest generation

Create the current system manifest:

```bash
go run ./cmd/fedora-state --format=manifest > manifest.json
```

### Plan generation

See what changes would be required:

```bash
go run ./cmd/fedora-state --format=plan > plan.json
```

Always review the plan before applying changes.

### Apply changes

Apply requires administrator privileges because it can call system package operations:

```bash
sudo fedora-state --format=apply
```

Recommended workflow:

```bash
1. Discover current state

go run ./cmd/fedora-state --format=json > state.json

2. Generate and inspect plan

go run ./cmd/fedora-state --format=plan > plan.json

3. Apply only after review

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

The project is written in Go.

Run tests:

```bash
go test ./...
```

Run static analysis:

```bash
go vet ./...
```
