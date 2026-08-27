module github.com/databricks/databricks-sql-kernel-bindings

go 1.23

// Per-platform submodules, each carrying that platform's kernel archive + cgo
// link directive. The root blank-imports the matching one (build-tag-gated).
// Required at real versions; replace points at in-tree sources for local dev
// (a dependency's replace is ignored by consumers).
require (
	github.com/databricks/databricks-sql-kernel-bindings/lib/darwin_amd64 v0.0.0
	github.com/databricks/databricks-sql-kernel-bindings/lib/darwin_arm64 v0.0.0
	github.com/databricks/databricks-sql-kernel-bindings/lib/linux_amd64 v0.0.0
	github.com/databricks/databricks-sql-kernel-bindings/lib/linux_arm v0.0.0
	github.com/databricks/databricks-sql-kernel-bindings/lib/linux_arm64 v0.0.0
	github.com/databricks/databricks-sql-kernel-bindings/lib/windows_amd64 v0.0.0
	github.com/databricks/databricks-sql-kernel-bindings/lib/windows_arm64 v0.0.0
)

replace (
	github.com/databricks/databricks-sql-kernel-bindings/lib/darwin_amd64 => ./lib/darwin_amd64
	github.com/databricks/databricks-sql-kernel-bindings/lib/darwin_arm64 => ./lib/darwin_arm64
	github.com/databricks/databricks-sql-kernel-bindings/lib/linux_amd64 => ./lib/linux_amd64
	github.com/databricks/databricks-sql-kernel-bindings/lib/linux_arm => ./lib/linux_arm
	github.com/databricks/databricks-sql-kernel-bindings/lib/linux_arm64 => ./lib/linux_arm64
	github.com/databricks/databricks-sql-kernel-bindings/lib/windows_amd64 => ./lib/windows_amd64
	github.com/databricks/databricks-sql-kernel-bindings/lib/windows_arm64 => ./lib/windows_arm64
)
