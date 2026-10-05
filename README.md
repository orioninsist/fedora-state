# fedora-state

Fedora system state discovery and reproducibility engine.

## Purpose

Discover installed tools and packages from a Fedora system, generate a verified state report, and reproduce the environment on another Fedora installation.

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

## Usage

```bash
fedora-state
```

Example output:

```text
os=linux architecture=amd64 distribution=Fedora Linux objects=4233 diagnostics=0
```

Available formats:

```bash
fedora-state --format=text
fedora-state --format=json
fedora-state --format=manifest
fedora-state --format=plan
fedora-state --format=apply
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
