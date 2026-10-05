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

Run:

```bash
go run ./cmd/fedora-state
```

## Usage

Basic system report:

```bash
fedora-state
```

Example output:

```text
os=linux architecture=amd64 distribution=Fedora Linux objects=4233 diagnostics=0
```

## Output formats

Text output:

```bash
fedora-state --format=text
```

JSON output:

```bash
fedora-state --format=json
```

Save JSON state:

```bash
fedora-state --format=json > state.json
```

Create manifest:

```bash
fedora-state --format=manifest
```

Generate execution plan:

```bash
fedora-state --format=plan
```

Apply planned changes:

```bash
sudo fedora-state --format=apply
```

Recommended workflow:

```bash
fedora-state --format=plan
```

Review the plan before applying:

```bash
sudo fedora-state --format=apply
```

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

## Scope

Included:

- package discovery
- source detection
- verification
- restore workflow
- system reports
- reproducible execution planning

Excluded:

- dotfiles
- personal configuration
- backups
- user data
- accounts

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
