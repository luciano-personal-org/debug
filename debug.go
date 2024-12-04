package debug

import (
	"fmt"
	"runtime"
	"runtime/debug"
	globaldebug "runtime/debug"
)

const (
	INFO  = "INFO"
	STACK = "STACK"
	MEM   = "MEM"
	GC    = "GC"
	BUILD = "BUILD"
	ALL   = "ALL"
)

// PrintDebug prints debug information
func PrintDebug(message string, option string) {

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
	if option == STACK || option == ALL {
		fmt.Printf("Stack Trace:\n%s\n", debug.Stack())
	}

	// Mem Stats
	if option == MEM || option == ALL {
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
	if option == GC || option == ALL {
		fmt.Printf("LastGC: %v\n", gcStats.LastGC)
		fmt.Printf("NumGC: %v\n", gcStats.NumGC)
		fmt.Printf("PauseTotal: %v\n", gcStats.PauseTotal)
		fmt.Printf("Pause: %v\n", gcStats.Pause)
		fmt.Printf("PauseEnd: %v\n", gcStats.PauseEnd)
		fmt.Printf("PauseQuantiles: %v\n", gcStats.PauseQuantiles)
	}

	// Build Info
	if option == BUILD || option == ALL {
		if ok {
			fmt.Printf("Build Info:\n%s\n", buildInfo)
		}
	}
}
