//go:build cgo && linux && arm

// Package kernellib_linux_arm carries the linux/arm (armv7) kernel static archive
// and the cgo directive that links it. Imported for its link side-effect by the
// root module's prebuilt_linux_arm.go under a matching build constraint.
//
// The committed libdatabricks_sql_kernel.a is a linux/arm (armv7-unknown-linux-gnueabihf)
// build of the kernel at the revision recorded in the driver's KERNEL_REV; it is
// produced by the peco-databricks-sql-kernel-go release workflow (cross-built on a
// Linux runner via gcc-arm-linux-gnueabihf) and committed beside this file. See
// the repo README and RELEASING.md.
//
// The -l:libdatabricks_sql_kernel.a form (GNU-ld extension) forces the static
// archive so the linker never prefers a same-named .so. The archive has no C++
// symbols, so unlike some platforms it needs no -lstdc++ (unwinding is provided
// by libgcc, auto-linked by the C driver); it needs only libm and libdl.
package kernellib_linux_arm

/*
#cgo LDFLAGS: -L${SRCDIR} -l:libdatabricks_sql_kernel.a -lm -ldl
*/
import "C"
