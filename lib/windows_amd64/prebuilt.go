//go:build cgo && windows && amd64

// Package kernellib_windows_amd64 carries the windows/amd64 kernel static
// archive and the cgo directive that links it. Imported for its link
// side-effect by the root module's prebuilt_windows_amd64.go under a matching
// build constraint.
//
// NOTE: the kernel static archive for this platform is NOT yet committed — it
// must be produced by the kernel repo's build-c-abi-libs workflow on a native
// windows/amd64 runner (the MSVC toolchain names it databricks_sql_kernel.lib;
// commit it here under the name the LDFLAGS below expect). Until the real
// archive is dropped in beside this file, a `-tags databricks_kernel` build for
// windows/amd64 fails at link with a missing-library error (by design). See the
// repo README and databricks-sql-kernel#244.
package kernellib_windows_amd64

/*
#cgo LDFLAGS: -L${SRCDIR} -l:libdatabricks_sql_kernel.a -lstdc++ -lm -lws2_32 -lbcrypt -lntdll -luserenv
*/
import "C"
