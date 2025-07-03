package main

import "fmt"

func main() {
	fmt.Printf("Example %s\n", version)
}

func Version() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		return info.Main.Version
	}
	return "(devel)"
}
