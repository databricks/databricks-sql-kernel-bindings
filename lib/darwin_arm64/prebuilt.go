//go:build cgo && darwin && arm64

// Package kernellib_darwin_arm64 carries the darwin/arm64 kernel static archive
// and the cgo directive that links it. It exports nothing; it is imported for
// its link side-effect by the root module's prebuilt_darwin_arm64.go under a
// matching build constraint. See the repo README.
//
// Darwin link specifics: Apple's ld64 does not accept the GNU `-l:<file>.a`
// form, so the archive is passed as a positional ${SRCDIR} input; -lc++ is the
// macOS C++ runtime; @loader_path keeps any dynamic reference resolvable.
package kernellib_darwin_arm64

/*
#cgo LDFLAGS: ${SRCDIR}/libdatabricks_sql_kernel.a -lc++ -lm -Wl,-rpath,@loader_path
*/
import "C"
