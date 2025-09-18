# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview
Murmur is a Go CLI that generates poetry using an OpenAI model. It features an interactive Bubble Tea UI and a command‑line mode.

## Common Commands
```sh
# Development
make help            # Show all make targets
make run             # Start the CLI interactively
make build            # Build the binary
make test             # Run all tests
make lint             # Run golangci‑lint

# Docker (optional)
make docker-build     # Build Docker image
make docker-run       # Run container locally
```

## Architecture
- `cmd/murmur/main.go` – entry point, parses flags, loads config, sets up logger, AI client, and Bubble Tea program.
- `internal/logger` – wrapper around `slog` that supports text (debug) and JSON (production).
- `internal/ai` – thin wrapper around the OpenAI Chat API. Exposes `GeneratePoem`.
- `internal/config` – loads configuration from environment variables with defaults.
- UI uses Bubble Tea to collect the poem prompt and then calls the AI client.

## Running Tests
Tests are written with Go’s testing package and can be run via `make test`.

## Building for Multiple Platforms
Use the provided make targets: `make build-linux`, `make build-windows`, `make build-darwin`.

## Notes
- The project uses `go.mod` for dependency management.
- All build artifacts are generated in the `build/` folder (default with `make build`).
