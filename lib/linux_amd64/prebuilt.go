//go:build cgo && linux && amd64

// Package kernellib_linux_amd64 carries the linux/amd64 kernel static archive
// and the cgo directive that links it. Imported for its link side-effect by the
// root module's prebuilt_linux_amd64.go under a matching build constraint.
//
// NOTE: libdatabricks_sql_kernel.a for this platform is NOT yet committed — it
// must be produced by the kernel repo's build-c-abi-libs workflow on a native
// linux/amd64 runner. Until the real archive is dropped in beside this file, a
// `-tags databricks_kernel` build for linux/amd64 fails at link with a
// missing-library error (by design — better than shipping a placeholder). See
// the repo README and databricks-sql-kernel#244.
//
// The -l:libdatabricks_sql_kernel.a form (GNU-ld extension) forces the static
// archive so the linker never prefers a same-named .so.
package kernellib_linux_amd64

/*
#cgo LDFLAGS: -L${SRCDIR} -l:libdatabricks_sql_kernel.a -lstdc++ -lm -ldl
*/
import "C"
