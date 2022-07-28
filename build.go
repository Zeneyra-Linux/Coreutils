package main

import (
	"fmt"
	"os"
)

// Build Locations
type BuildLocation string

const (
	ZigLocation  BuildLocation = "zig-out/bin/%s"
	RustLocation BuildLocation = "target/release/%s"
	CLocation    BuildLocation = "%s"
)

// Build Types
type BuildType string

const (
	ZigType  BuildType = "zig"
	RustType BuildType = "rust"
	CType    BuildType = "c"
)

// Build Commands
type BuildCmd string

const (
	ZigBuild  BuildCmd = "zig build"
	RustBuild BuildCmd = "cargo zigbuild --release"
	CBuild    BuildCmd = "zig cc main.c -o %s -s"
)

// Map of Coreutils Programs (Key: program name; Value: project type)
var projects = map[string]BuildType{
	"true":  CType,
	"false": CType,
}

// Build Wrapper
func build(name string, protype BuildType) {
	// Run build commands
}

// Clean Wrapper
func clean(name string, protype BuildType) {
	// Remove build files
}

// Main Function
func main() {
	for program, project := range projects {
		if len(os.Args) > 1 {
			switch os.Args[1] {
			case "build":
				build(program, project)
			case "clean":
				clean(program, project)
			default:
				println(fmt.Sprintf("Unknown command: %s", os.Args[1]))
				os.Exit(1)
			}
		} else {
			build(program, project)
		}
	}
}
