# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go debugging utility package that provides runtime debugging capabilities including stack traces, memory statistics, GC stats, and build information. It's part of the larger "algo-platform" monorepo structure.

## Development Commands

### Building and Running
- `go build` - Build the package
- `go test` - Run tests
- `go mod tidy` - Clean up dependencies

### Hot Reload Development
- Install AIR for hot reload: `curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin`
- Start development with hot reload: `air`
- Run with Docker: `docker-compose up -d && air`

## Architecture

This is a simple Go package that exposes debugging utilities through two main functions:

### Core Components
- **debug.go**: Main package file containing the debug functionality
- **DebugOptions struct**: Configuration for debug level and enable/disable state
- **Debug levels**: INFO, STACK, MEM, GC, BUILD, ALL

### Key Functions
- `PrintDebug(message, options)` - Legacy console output version
- `PrintDebugWithLog(message, options, logger)` - Logger-based version using logr.Logger

### Dependencies
- `github.com/go-logr/logr v1.4.2` - Structured logging interface
- Standard Go runtime packages for memory/GC stats

## Development Notes

### Environment Setup
The README mentions GoFiber web framework and MongoDB configuration, but this specific debug package appears to be a standalone utility that can be used across different projects in the platform.

### Module Structure
- Module path: `github.com/luciano-personal-org/debug`
- Go version: 1.23.3
- Single dependency on go-logr for structured logging

This package provides runtime debugging capabilities that can be integrated into larger applications, with both console and structured logging output options.