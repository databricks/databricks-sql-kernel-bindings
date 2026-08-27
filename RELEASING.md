# Releasing the Go bindings

This repo packages the kernel's prebuilt per-platform C-ABI archives as Go
modules. A release **wraps a specific `databricks-sql-kernel` version and is
versioned to match it** — bindings `vX.Y.Z` is built from kernel tag `vX.Y.Z`.
The C-ABI/behaviour changes for a release are the kernel's (see the kernel
[`CHANGELOG.md`](https://github.com/databricks/databricks-sql-kernel/blob/main/CHANGELOG.md));
this repo only packages the archives + records binding-repo notes in
[`CHANGELOG.md`](./CHANGELOG.md).

## The pin chain (all three must be the same kernel version)

A kernel opt-in build links three things that MUST come from one kernel commit:

1. **The archives here** — `lib/<platform>/libdatabricks_sql_kernel.a`, built
   from kernel `vX.Y.Z`.
2. **The committed C header** — the driver commits `databricks_kernel.h` at
   `internal/backend/kernel/include/` (its cgo layer `#include`s it), synced to
   the same kernel `vX.Y.Z`. It is **not** committed in this repo: the bindings
   modules carry only the archive + `#cgo LDFLAGS`, so nothing here compiles
   against the header.
3. **The driver's pin** — `databricks-sql-go` requires the five `lib/<platform>`
   modules at `vX.Y.Z` in its `go.mod`, and records the kernel commit in its
   `KERNEL_REV` file. Upgrading the driver's `go.mod` is what moves the kernel
   version for consumers.

If the header and the archive disagree (e.g. header declares a symbol the archive
doesn't export), the driver fails to link — so never bump one without the others.

## Cutting a release

1. **Kernel tag.** Confirm/cut `databricks-sql-kernel` tag `vX.Y.Z`. The tagged
   commit MUST contain the committed C-ABI header (tags ≤ `v0.2.0` predate it).
2. **Build.** Build the per-platform archives from the tagged kernel via
   Databricks' internal release pipeline (Databricks maintainers: see the
   `release-kernel-bindings` runbook). It security-scans the archives and produces
   a `kernel-go-bindings-vX.Y.Z` artifact holding each
   `dist/<platform>/libdatabricks_sql_kernel.a`. Windows archives use the GNU
   (`-gnu` / `-gnullvm`) triples — Go cgo links a GNU `.a`, never an MSVC `.lib`.
3. **Stage.** Download the artifact; copy each
   `dist/<platform>/libdatabricks_sql_kernel.a` into `lib/<platform>/`. (The
   artifact also contains `databricks_kernel.h`, but this repo does not commit it
   — the driver syncs the header into its own tree from the kernel at `KERNEL_REV`.)
4. **Changelog.** Add a `## [vX.Y.Z]` entry to [`CHANGELOG.md`](./CHANGELOG.md)
   (kernel rev, platforms, any packaging/link notes).
5. **Commit + tag.** Commit, then push the path-prefixed module tags plus the
   root tag, all at that commit:
   ```bash
   for p in linux_amd64 linux_arm64 linux_arm darwin_amd64 darwin_arm64 windows_amd64 windows_arm64; do
     git tag "lib/$p/vX.Y.Z"
   done
   git tag "vX.Y.Z"
   git push origin main --tags
   ```
6. **Driver.** In `databricks-sql-go`, bump the `go.mod` requires to `vX.Y.Z`,
   refresh `go.sum`, set `KERNEL_REV` to the `vX.Y.Z` commit, and sync the
   committed header (`internal/backend/kernel/include/`). Databricks maintainers:
   see the `release-kernel-bindings` runbook for the full driver-side flow.

## Immutability

Go module versions are **immutable** once the public proxy / checksum database
has served them — never move a released tag; always cut a new version. (A
brand-new module path may take its first tag at an existing version, since that
path was never published — that's how a new platform can be added at the current
version without re-tagging the others.)

## Adding a platform

Create `lib/<os>_<arch>/` with its own `go.mod`, a build-tag-gated `prebuilt.go`
(matching `//go:build` + `#cgo LDFLAGS`), and the committed archive; add a root
`prebuilt_<os>_<arch>.go` shim; add the platform to the release workflow's build
matrix and to the driver's `go.mod` + `cgo_<os>_<arch>.go` shim; tag it alongside
the others.
