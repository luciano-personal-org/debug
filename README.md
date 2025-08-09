# Debug - Go Runtime Debugging Utility

A comprehensive Go package for runtime debugging and performance monitoring. This utility provides detailed insights into your application's memory usage, garbage collection statistics, stack traces, and build information.

## Features

- **Stack Trace Analysis**: Capture and display complete stack traces
- **Memory Statistics**: Detailed memory allocation and usage metrics
- **Garbage Collection Stats**: GC performance and timing information
- **Build Information**: Runtime build details and version info
- **Flexible Logging**: Support for both console output and structured logging
- **Configurable Levels**: Choose specific debug information or view all at once

## Installation

```bash
go get github.com/luciano-personal-org/debug
```

## Quick Start

### Basic Usage

```go
package main

import (
    "github.com/luciano-personal-org/debug"
)

func main() {
    options := debug.DebugOptions{
        Enabled: true,
        Level:   debug.INFO,
    }
    
    debug.PrintDebug("Application started", options)
}
```

### With Memory Statistics

```go
options := debug.DebugOptions{
    Enabled: true,
    Level:   debug.MEM,
}

debug.PrintDebug("Memory usage check", options)
```

### With Structured Logging

```go
import (
    "github.com/go-logr/logr"
    "github.com/luciano-personal-org/debug"
)

func debugWithLogger(logger logr.Logger) {
    options := debug.DebugOptions{
        Enabled: true,
        Level:   debug.ALL,
    }
    
    debug.PrintDebugWithLog("Comprehensive debug info", options, logger)
}
```

## Debug Levels

| Level | Description |
|-------|-------------|
| `INFO` | Basic information (default) |
| `STACK` | Stack trace information |
| `MEM` | Memory allocation statistics |
| `GC` | Garbage collection statistics |
| `BUILD` | Build and version information |
| `ALL` | All available debug information |

## API Reference

### Types

#### DebugOptions
```go
type DebugOptions struct {
    Enabled bool   // Enable/disable debugging
    Level   string // Debug level (INFO, STACK, MEM, GC, BUILD, ALL)
}
```

### Functions

#### PrintDebug
```go
func PrintDebug(message string, options DebugOptions) error
```
Prints debug information to standard output.

**Parameters:**
- `message`: Custom debug message
- `options`: Debug configuration

**Returns:**
- `error`: Validation error if invalid debug level provided

#### PrintDebugWithLog
```go
func PrintDebugWithLog(message string, options DebugOptions, logger logr.Logger) error
```
Outputs debug information using structured logging.

**Parameters:**
- `message`: Custom debug message
- `options`: Debug configuration
- `logger`: logr.Logger instance for structured output

**Returns:**
- `error`: Validation error if invalid debug level provided

## Examples

### Memory Monitoring

```go
func monitorMemoryUsage() {
    options := debug.DebugOptions{
        Enabled: true,
        Level:   debug.MEM,
    }
    
    // Before operation
    debug.PrintDebug("Before heavy operation", options)
    
    // Perform memory-intensive operation
    data := make([]byte, 1024*1024*100) // 100MB allocation
    
    // After operation
    debug.PrintDebug("After heavy operation", options)
    
    _ = data // Use data to prevent optimization
}
```

### Error Investigation with Stack Trace

```go
func investigateError() {
    defer func() {
        if r := recover(); r != nil {
            options := debug.DebugOptions{
                Enabled: true,
                Level:   debug.STACK,
            }
            debug.PrintDebug("Panic recovered", options)
        }
    }()
    
    // Code that might panic
    panic("Something went wrong")
}
```

### Comprehensive Performance Analysis

```go
func performanceAnalysis() {
    options := debug.DebugOptions{
        Enabled: true,
        Level:   debug.ALL,
    }
    
    debug.PrintDebug("Performance analysis checkpoint", options)
}
```

## Memory Statistics Explained

When using `MEM` or `ALL` levels, the following metrics are displayed:

- **Alloc**: Current heap memory in use
- **TotalAlloc**: Total heap memory allocated (cumulative)
- **HeapAlloc**: Current heap allocation
- **HeapSys**: Heap memory from OS
- **HeapIdle**: Idle heap memory
- **HeapInuse**: In-use heap memory
- **HeapReleased**: Heap memory returned to OS
- **HeapObjects**: Number of heap objects
- **StackInUse**: Stack memory in use
- **StackSys**: Stack memory from OS
- **NumGC**: Number of GC cycles completed

## GC Statistics Explained

When using `GC` or `ALL` levels:

- **LastGC**: Time of last garbage collection
- **NumGC**: Number of GC cycles
- **PauseTotal**: Total GC pause time
- **Pause**: Individual GC pause times
- **PauseEnd**: End times of GC pauses
- **PauseQuantiles**: GC pause time quantiles

## Development

### Prerequisites

- Go 1.23.3 or later
- Dependencies managed via `go.mod`

### Building

```bash
go build
```

### Testing

```bash
go test
```

### Hot Reload Development

Install AIR for development with hot reload:

```bash
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```

Start development:

```bash
air
```

## Dependencies

- [go-logr/logr](https://github.com/go-logr/logr) v1.4.2 - Structured logging interface

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the terms found in the LICENSE file.
