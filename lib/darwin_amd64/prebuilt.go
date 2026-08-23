//go:build cgo && darwin && amd64

// Package kernellib_darwin_amd64 carries the darwin/amd64 kernel static archive
// and the cgo directive that links it. Imported for its link side-effect by the
// root module's prebuilt_darwin_amd64.go under a matching build constraint.
//
// Darwin link specifics: Apple's ld64 does not accept the GNU `-l:<file>.a`
// form, so the archive is passed as a positional ${SRCDIR} input; -lc++ is the
// macOS C++ runtime and -lm the math library. No -rpath is set: the archive is
// linked statically, so there is no dynamic library to locate at run time, and
// -Wl,-rpath is rejected by cgo's default LDFLAGS allowlist (which would break a
// plain `go build -tags databricks_kernel`).
package kernellib_darwin_amd64

/*
#cgo LDFLAGS: ${SRCDIR}/libdatabricks_sql_kernel.a -lc++ -lm
*/
import "C"
