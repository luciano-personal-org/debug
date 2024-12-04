package debug

import (
	"fmt"
	"runtime"
	"runtime/debug"
	globaldebug "runtime/debug"
)

// DebugOptions enum
type DebugOptions int

// DebugOptions enum values
const (
	STACKTRACE DebugOptions = iota
	MEMSTATS
	GCSTATS
	BUILDINFO
)

// PrintDebug prints debug information
func PrintDebug(message string, options ...DebugOptions) {

	var gcStats globaldebug.GCStats
	var memStats runtime.MemStats

	globaldebug.ReadGCStats(&gcStats)
	runtime.ReadMemStats(&memStats)
	buildInfo, ok := globaldebug.ReadBuildInfo()

	// Initial Message
	fmt.Println("Debugging...")

	// Business Message
	fmt.Println(message)

	// Stack Trace
	if containsOption(options, STACKTRACE) {
		fmt.Printf("Stack Trace:\n%s\n", debug.Stack())
	}

	// Mem Stats
	if containsOption(options, MEMSTATS) {
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
	if containsOption(options, GCSTATS) {
		fmt.Printf("LastGC: %v\n", gcStats.LastGC)
		fmt.Printf("NumGC: %v\n", gcStats.NumGC)
		fmt.Printf("PauseTotal: %v\n", gcStats.PauseTotal)
		fmt.Printf("Pause: %v\n", gcStats.Pause)
		fmt.Printf("PauseEnd: %v\n", gcStats.PauseEnd)
		fmt.Printf("PauseQuantiles: %v\n", gcStats.PauseQuantiles)
	}

	// Build Info
	if containsOption(options, BUILDINFO) {
		if ok {
			fmt.Printf("Build Info:\n%s\n", buildInfo)
		}
	}

}

func containsOption(options []DebugOptions, option DebugOptions) bool {
	for _, opt := range options {
		if opt == option {
			return true
		}
	}
	return false
}
