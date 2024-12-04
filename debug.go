package debug

import (
	"fmt"
	"runtime"
	"runtime/debug"
	globaldebug "runtime/debug"
)

// DebugOptions is a type for debug options
const (
	INFO       = "INFO"  // Default
	STACKTRACE = "STACK" // Stack Trace
	MEMSTATS   = "MEM"   // Memory Stats
	GCSTATS    = "GC"    // GC Stats
	BUILDSTATS = "BUILD" // Build Info
	ALLSTATS   = "ALL"   // All Stats
)

// isValidOption validates the debug option
func isValidOption(option string) bool {
	switch option {
	case INFO, STACKTRACE, MEMSTATS, GCSTATS, BUILDSTATS, ALLSTATS:
		return true
	default:
		return false
	}
}

// PrintDebug prints debug information
func PrintDebug(message string, option string) {
	if !isValidOption(option) {
		fmt.Println("Invalid debug option: ", option)
		return
	}

	var gcStats globaldebug.GCStats
	var memStats runtime.MemStats

	globaldebug.ReadGCStats(&gcStats)
	runtime.ReadMemStats(&memStats)
	buildInfo, ok := globaldebug.ReadBuildInfo()

	// Initial Messages
	fmt.Println("Debugging...")

	// Business Message
	fmt.Println(message)

	// Stack Trace
	if option == STACKTRACE || option == ALLSTATS {
		fmt.Printf("Stack Trace:\n%s\n", debug.Stack())
	}

	// Mem Stats
	if option == MEMSTATS || option == ALLSTATS {
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
	if option == GCSTATS || option == ALLSTATS {
		fmt.Printf("LastGC: %v\n", gcStats.LastGC)
		fmt.Printf("NumGC: %v\n", gcStats.NumGC)
		fmt.Printf("PauseTotal: %v\n", gcStats.PauseTotal)
		fmt.Printf("Pause: %v\n", gcStats.Pause)
		fmt.Printf("PauseEnd: %v\n", gcStats.PauseEnd)
		fmt.Printf("PauseQuantiles: %v\n", gcStats.PauseQuantiles)
	}

	// Build Info
	if option == BUILDSTATS || option == ALLSTATS {
		if ok {
			fmt.Printf("Build Info:\n%s\n", buildInfo)
		}
	}
}
