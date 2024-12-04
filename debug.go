package debug

import (
	"errors"
	"fmt"
	"runtime"
	"runtime/debug"
	globaldebug "runtime/debug"
)

// DebugOptions is a type for debug options
const (
	INFO  = "INFO"  // Default
	STACK = "STACK" // Stack Trace
	MEM   = "MEM"   // Memory Stats
	GC    = "GC"    // GC Stats
	BUILD = "BUILD" // Build Info
	ALL   = "ALL"   // All Stats
)

type DebugOptions struct {
	Enabled bool
	Level   string
}

// isValidOption validates the debug option
func isValidOption(level string) bool {
	switch level {
	case INFO, STACK, MEM, GC, BUILD, ALL:
		return true
	default:
		return false
	}
}

// PrintDebug prints debug information
func PrintDebug(message string, level string) error {
	if !isValidOption(level) {
		return errors.New("invalid debug option: " + level)
	}

	// Read GC Stats
	var gcStats globaldebug.GCStats
	globaldebug.ReadGCStats(&gcStats)

	// Read Mem Stats
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	// Read Build Info
	buildInfo, ok := globaldebug.ReadBuildInfo()

	// Initial Messages
	fmt.Println("\nStart debug...")

	// Business Message
	fmt.Println(message)

	// Stack Trace
	if level == STACK || level == ALL {
		fmt.Printf("Stack Trace:\n%s\n", debug.Stack())
	}

	// Mem Stats
	if level == MEM || level == ALL {
		fmt.Printf("Alloc: %d bytes\n", memStats.Alloc)
		fmt.Printf("TotalAlloc: %d bytes\n", memStats.TotalAlloc)
		fmt.Printf("HeapAlloc: %d bytes\n", memStats.HeapAlloc)
		fmt.Printf("HeapSys: %d bytes\n", memStats.HeapSys)
		fmt.Printf("HeapIdle: %d bytes\n", memStats.HeapIdle)
		fmt.Printf("HeapInuse: %d bytes\n", memStats.HeapInuse)
		fmt.Printf("HeapReleased: %d bytes\n", memStats.HeapReleased)
		fmt.Printf("HeapObjects: %d\n", memStats.HeapObjects)
		fmt.Printf("NumGC: %d\n", memStats.NumGC)
	}

	// GC Stats
	if level == GC || level == ALL {
		fmt.Printf("LastGC: %v\n", gcStats.LastGC)
		fmt.Printf("NumGC: %v\n", gcStats.NumGC)
		fmt.Printf("PauseTotal: %v\n", gcStats.PauseTotal)
		fmt.Printf("Pause: %v\n", gcStats.Pause)
		fmt.Printf("PauseEnd: %v\n", gcStats.PauseEnd)
		fmt.Printf("PauseQuantiles: %v\n", gcStats.PauseQuantiles)
	}

	// Build Info
	if level == BUILD || level == ALL {
		if ok {
			fmt.Printf("Build Info:\n%s\n", buildInfo)
		}
	}

	fmt.Println("\nFinish debug...")

	return nil
}
