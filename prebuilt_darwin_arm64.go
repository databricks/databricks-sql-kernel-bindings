//go:build cgo && databricks_kernel && darwin && arm64

// Package kernelbindings is the root binding module. Its per-platform files
// blank-import the matching lib/<platform> module so cgo collects that module's
// #cgo LDFLAGS at final link time, pulling libdatabricks_sql_kernel.a into the
// binary. The Databricks SQL Go driver imports this package (build-tag-gated) to
// link the kernel; the actual C ABI call layer lives in the driver.
package kernelbindings

import _ "github.com/databricks/databricks-sql-kernel-bindings/lib/darwin_arm64"
