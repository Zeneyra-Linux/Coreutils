# Coreutils
Rewrite of the GNU Coreutils

# Compiling
## SDKs
In order to compile all Coreutils, you need these things to be installed:
* Cargo (Rust), see [rustup](https://www.rust-lang.org/tools/install) for more.
* [Zig](https://ziglang.org/), used for Zig programs `zig build`, C programs `zig cc` and Rust programs (`cargo zigbuild`).
* [Cargo Zigbuild](https://github.com/messense/cargo-zigbuild)

## Optional
* [Go](https://go.dev/), to run the build script that automates everything, though this is optional.

The final build command is `go run build.go`

# Installing
## From Source
1. Clone the repository with `git clone https://github.com/Zeneyra-Linux/Coreutils.git`
2. `cd Coreutils`
3. See [above](#compiling)

## Via a package manager
WIP.