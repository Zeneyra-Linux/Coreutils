package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// Build Locations
type BuildLocation string

const (
	ZigLocation  BuildLocation = "zig-out/bin/%s"
	RustLocation BuildLocation = "target/release/%s"
)

// Build Types
type BuildType string

const (
	ZigType  BuildType = "zig"
	RustType BuildType = "rust"
	CType    BuildType = "c"
)

// ----------------------------------------------------------------------------------------------------------------------- //
// ----------------------------------------------------------------------------------------------------------------------- //
// ----------------------------------------------------------------------------------------------------------------------- //

// Init processes
func new_cmd(dir string, name string, args ...string) error {
	fmt.Println(fmt.Sprintf("Building \u001b[36m%s\u001b[0m...", dir))
	cmd := exec.Command(name, args...)
	cmd.Dir = filepath.Join(cwd, dir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Build Wrapper
func build(name string, protype BuildType) error {
	switch protype {
	case ZigType:
		return new_cmd(name, "zig", "build")
	case RustType:
		return new_cmd(name, "cargo", "zigbuild", "--release")
	case CType:
		if runtime.GOOS == "windows" {
			return new_cmd(name, "zig", "cc", "main.c", "-o", fmt.Sprintf("../build/%s.exe", name), "-s")
		} else {
			return new_cmd(name, "zig", "cc", "main.c", "-o", fmt.Sprintf("../build/%s", name), "-s")
		}
	}
	return nil
}

// Wrapper for Build Wrapper
func build_init() {
	unsuccessful := 0
	os.Mkdir(builddir, 0755)
	for program, project := range projects {
		err := build(program, project)
		if err != nil {
			unsuccessful++
			println(err.Error())
		} else {
			if project == ZigType {
				os.Rename(fmt.Sprintf(string(ZigLocation), program), filepath.Join(builddir, program))
			}
			if project == RustType {
				os.Rename(fmt.Sprintf(string(RustLocation), program), filepath.Join(builddir, program))
			}
		}
	}
	if unsuccessful > 0 {
		println(fmt.Sprintf("\u001b[31m%d programs failed to build!\u001b[0m", unsuccessful))
		os.Exit(1)
	} else {
		fmt.Println("\u001b[32mAll programs built successfully!\u001b[0m")
	}
}

// ----------------------------------------------------------------------------------------------------------------------- //
// ----------------------------------------------------------------------------------------------------------------------- //
// ----------------------------------------------------------------------------------------------------------------------- //

// Clean Wrapper
func clean(name string, protype BuildType) {
	switch protype {
	case ZigType:
		fmt.Println("Removing \u001b[36" + name + "/zig-out/\u001b[0m...")
		os.RemoveAll(filepath.Join(cwd, name, "zig-out"))
		fmt.Println("Removing \u001b[36" + name + "/zig-cache/\u001b[0m...")
		os.RemoveAll(filepath.Join(cwd, name, "zig-cache"))
	case RustType:
		fmt.Println("Removing \u001b[36" + name + "/target/\u001b[0m...")
		os.RemoveAll(filepath.Join(cwd, name, "target"))
	}
}

// ----------------------------------------------------------------------------------------------------------------------- //
// ----------------------------------------------------------------------------------------------------------------------- //
// ----------------------------------------------------------------------------------------------------------------------- //

// Paths
var cwd string
var builddir string

// Map of Coreutils Programs (Key: program name; Value: project type)
var projects = map[string]BuildType{
	"true":  CType,
	"false": CType,
}

// Main Function
func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	cwd = path
	builddir = filepath.Join(cwd, "build")
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "build":
			build_init()
		case "clean":
			fmt.Println("Removing \u001b[36mbuild/\u001b[0m...")
			os.RemoveAll(builddir)
			for program, project := range projects {
				clean(program, project)
			}
		default:
			println(fmt.Sprintf("\u001b[31mUnknown command: %s\u001b[0m", os.Args[1]))
			os.Exit(1)
		}
	} else {
		build_init()
	}
	os.Exit(0)
}
