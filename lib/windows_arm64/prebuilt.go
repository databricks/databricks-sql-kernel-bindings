//go:build cgo && windows && arm64

// Package kernellib_windows_arm64 carries the windows/arm64 kernel static archive
// and the cgo directive that links it. Imported for its link side-effect by the
// root module's prebuilt_windows_arm64.go under a matching build constraint.
//
// The committed libdatabricks_sql_kernel.a is a windows/arm64 GNU archive built
// for the aarch64-pc-windows-GNULLVM (llvm-mingw) target — NOT MSVC. Go cgo on
// Windows links through a GNU toolchain and consumes a GNU `.a`, never an MSVC
// `.lib`. It is produced by the peco-databricks-sql-kernel-go release workflow
// (cross-built on a Linux runner via llvm-mingw) and committed beside this file.
// See the repo README and RELEASING.md.
//
// The archive has no C++ symbols (no -lstdc++/-lc++), and embeds most of its
// Windows imports via raw-dylib (ws2_32/userenv/secur32/kernel32/...). The two
// load-time imports it does not embed — BCryptGenRandom (bcrypt) and
// NtCreateNamedPipeFile (ntdll) — are supplied explicitly below. Unwinding and
// compiler-rt builtins are auto-linked by the llvm-mingw clang driver.
package kernellib_windows_arm64

/*
#cgo LDFLAGS: -L${SRCDIR} -l:libdatabricks_sql_kernel.a -lbcrypt -lntdll
*/
import "C"
