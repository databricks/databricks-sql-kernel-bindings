module github.com/databricks/databricks-sql-kernel-bindings

go 1.23

// The root module blank-imports the per-platform lib/<platform> submodules
// (build-tag-gated) so cgo links the right archive. Require them at their real
// released versions; the replace lets THIS repo build against the in-tree
// sources during development (a dependency's replace is ignored by consumers).
require github.com/databricks/databricks-sql-kernel-bindings/lib/darwin_arm64 v0.0.0

replace github.com/databricks/databricks-sql-kernel-bindings/lib/darwin_arm64 => ./lib/darwin_arm64
