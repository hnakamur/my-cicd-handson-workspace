package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	fmt.Printf("Example %s\n", version())
}

func version() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		return info.Main.Version
	}
	return "(devel)"
}
