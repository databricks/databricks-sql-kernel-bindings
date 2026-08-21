//go:build cgo && linux && arm64

// Package kernellib_linux_arm64 carries the linux/arm64 kernel static archive
// and the cgo directive that links it. Imported for its link side-effect by the
// root module's prebuilt_linux_arm64.go under a matching build constraint.
//
// NOTE: libdatabricks_sql_kernel.a for this platform is NOT yet committed — it
// must be produced by the kernel repo's build-c-abi-libs workflow on a native
// linux/arm64 runner. Until the real archive is dropped in beside this file, a
// `-tags databricks_kernel` build for linux/arm64 fails at link with a
// missing-library error (by design). See the repo README and
// databricks-sql-kernel#244.
package kernellib_linux_arm64

/*
#cgo LDFLAGS: -L${SRCDIR} -l:libdatabricks_sql_kernel.a -lstdc++ -lm -ldl
*/
import "C"
