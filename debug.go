package debug

import (
	"errors"
	"fmt"
	"runtime"
	"runtime/debug"
	globaldebug "runtime/debug"

	"github.com/DataDog/appsec-internal-go/log"
	"github.com/go-logr/logr"
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

// Legacy Version
// PrintDebug prints debug information
func PrintDebug(message string, options DebugOptions) error {

	var level = options.Level

	// If Debug is disabled, return
	if !options.Enabled {
		return nil
	}

	// Default Level
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
		fmt.Printf("StackInUse: %d bytes\n", memStats.StackInuse)
		fmt.Printf("StackSys: %d bytes\n", memStats.StackSys)
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

// PrintDebug prints debug information
func PrintDebugWithLog(message string, options DebugOptions, logger logr.Logger) error {

	var level = options.Level

	// If Debug is disabled, return
	if !options.Enabled {
		return nil
	}

	// Default Level
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

	// // Initial Messages
	// fmt.Println("\nStart debug...")

	// Business Message
	logger.WithCallDepth(2).Info(message)

	// Stack Trace
	if level == STACK || level == ALL {
		stack := fmt.Sprintf("Stack Trace:%s", debug.Stack())
		log.Info(stack)
	}

	// Mem Stats
	if level == MEM || level == ALL {
		message_level := fmt.Sprintf("Alloc: %d bytes", memStats.Alloc)
		logger.WithCallDepth(2).Info(message_level)
		message_level = fmt.Sprintf("TotalAlloc: %d bytes", memStats.TotalAlloc)
		logger.WithCallDepth(2).Info(message_level)
		message_level = fmt.Sprintf("HeapAlloc: %d bytes", memStats.HeapAlloc)
		logger.WithCallDepth(2).Info(message_level)
		message_level = fmt.Sprintf("HeapSys: %d bytes", memStats.HeapSys)
		logger.WithCallDepth(2).Info(message_level)
		message_level = fmt.Sprintf("HeapIdle: %d bytes", memStats.HeapIdle)
		logger.WithCallDepth(2).Info(message_level)
		message_level = fmt.Sprintf("HeapInuse: %d bytes", memStats.HeapInuse)
		logger.WithCallDepth(2).Info(message_level)
		message_level = fmt.Sprintf("HeapReleased: %d bytes", memStats.HeapReleased)
		logger.WithCallDepth(2).Info(message_level)
		message_level = fmt.Sprintf("HeapObjects: %d", memStats.HeapObjects)
		logger.WithCallDepth(2).Info(message_level)
		message_level = fmt.Sprintf("StackInUse: %d bytes", memStats.StackInuse)
		logger.WithCallDepth(2).Info(message_level)
		message_level = fmt.Sprintf("StackSys: %d bytes", memStats.StackSys)
		logger.WithCallDepth(2).Info(message_level)
		message_level = fmt.Sprintf("NumGC: %d", memStats.NumGC)
		logger.WithCallDepth(2).Info(message_level)
	}

	// GC Stats
	if level == GC || level == ALL {
		message_gc := fmt.Sprintf("LastGC: %v", gcStats.LastGC)
		logger.WithCallDepth(2).Info(message_gc)
		message_gc = fmt.Sprintf("NumGC: %v", gcStats.NumGC)
		logger.WithCallDepth(2).Info(message_gc)
		message_gc = fmt.Sprintf("PauseTotal: %v", gcStats.PauseTotal)
		logger.WithCallDepth(2).Info(message_gc)
		message_gc = fmt.Sprintf("Pause: %v", gcStats.Pause)
		logger.WithCallDepth(2).Info(message_gc)
		message_gc = fmt.Sprintf("PauseEnd: %v", gcStats.PauseEnd)
		logger.WithCallDepth(2).Info(message_gc)
		message_gc = fmt.Sprintf("PauseQuantiles: %v", gcStats.PauseQuantiles)
		logger.WithCallDepth(2).Info(message_gc)
	}

	// Build Info
	if level == BUILD || level == ALL {
		if ok {
			message_build := fmt.Sprintf("Build Info: %s", buildInfo)
			logger.WithCallDepth(2).Info(message_build)
		}
	}

	return nil
}
