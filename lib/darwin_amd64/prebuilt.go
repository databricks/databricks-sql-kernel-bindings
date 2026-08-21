//go:build cgo && darwin && amd64

// Package kernellib_darwin_amd64 carries the darwin/amd64 kernel static archive
// and the cgo directive that links it. Imported for its link side-effect by the
// root module's prebuilt_darwin_amd64.go under a matching build constraint.
//
// NOTE: libdatabricks_sql_kernel.a for this platform is NOT yet committed — it
// must be produced by the kernel repo's build-c-abi-libs workflow on a native
// darwin/amd64 (Intel) runner. Until the real archive is dropped in beside this
// file, a `-tags databricks_kernel` build for darwin/amd64 fails at link with a
// missing-library error (by design). See the repo README and
// databricks-sql-kernel#244.
package kernellib_darwin_amd64

/*
#cgo LDFLAGS: ${SRCDIR}/libdatabricks_sql_kernel.a -lc++ -lm -Wl,-rpath,@loader_path
*/
import "C"
