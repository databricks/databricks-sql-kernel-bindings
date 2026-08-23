//go:build cgo && linux && amd64

// Package kernellib_linux_amd64 carries the linux/amd64 kernel static archive
// and the cgo directive that links it. Imported for its link side-effect by the
// root module's prebuilt_linux_amd64.go under a matching build constraint.
//
// The committed libdatabricks_sql_kernel.a is a linux/amd64 (x86_64-unknown-linux-gnu)
// build of the kernel at the revision recorded in the driver's KERNEL_REV; it is
// produced by the kernel repo's build-c-abi-libs workflow (or an equivalent
// cross-build) and committed beside this file. See the repo README and
// databricks-sql-kernel#244.
//
// The -l:libdatabricks_sql_kernel.a form (GNU-ld extension) forces the static
// archive so the linker never prefers a same-named .so.
package kernellib_linux_amd64

/*
#cgo LDFLAGS: -L${SRCDIR} -l:libdatabricks_sql_kernel.a -lstdc++ -lm -ldl
*/
import "C"
