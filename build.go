package main

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
func build() {

}

// Clean Wrapper
func clean() {

}

// Main Function
func main() {
	for program, project := range projects {

	}
}
