//go:build cgo && windows && amd64

// Package kernellib_windows_amd64 carries the windows/amd64 kernel static
// archive and the cgo directive that links it. Imported for its link
// side-effect by the root module's prebuilt_windows_amd64.go under a matching
// build constraint.
//
// The committed libdatabricks_sql_kernel.a is a windows/amd64 GNU archive built
// for the x86_64-pc-windows-GNU (MinGW) target — NOT MSVC. Go cgo on Windows
// links through MinGW/gcc and consumes a GNU `.a`, never an MSVC `.lib`, so the
// kernel is built for the -gnu triple and committed under the .a name the
// LDFLAGS below expect. The archive is at the revision recorded in the driver's
// KERNEL_REV, produced by the kernel repo's build-c-abi-libs workflow (or an
// equivalent cross-build). See the repo README and databricks-sql-kernel#244.
package kernellib_windows_amd64

/*
#cgo LDFLAGS: -L${SRCDIR} -l:libdatabricks_sql_kernel.a -lstdc++ -lm -lws2_32 -lbcrypt -lntdll -luserenv
*/
import "C"
