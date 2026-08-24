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
2. **The committed C header** — `include/databricks_kernel.h` here, and the copy
   the driver commits at `internal/backend/kernel/include/` — both synced to the
   same kernel `vX.Y.Z`.
3. **The driver's pin** — `databricks-sql-go` requires the five `lib/<platform>`
   modules at `vX.Y.Z` in its `go.mod`, and records the kernel commit in its
   `KERNEL_REV` file. Upgrading the driver's `go.mod` is what moves the kernel
   version for consumers.

If the header and the archive disagree (e.g. header declares a symbol the archive
doesn't export), the driver fails to link — so never bump one without the others.

## Cutting a release

1. **Kernel tag.** Confirm/cut `databricks-sql-kernel` tag `vX.Y.Z`. The tagged
   commit MUST contain the committed C-ABI header (tags ≤ `v0.2.0` predate it).
2. **Build.** Run the `peco-databricks-sql-kernel-go` workflow (in
   `databricks/secure-public-registry-releases-eng`) with `ref=vX.Y.Z`. It builds
   all five archives (GitHub-hosted macOS/Windows + protected Linux; Windows as
   `x86_64-pc-windows-gnu`), security-scans them, and uploads a
   `kernel-go-bindings-vX.Y.Z` artifact.
3. **Stage.** Download the artifact; copy each
   `dist/<platform>/libdatabricks_sql_kernel.a` into `lib/<platform>/`, and copy
   `dist/<any>/databricks_kernel.h` to `include/databricks_kernel.h`.
4. **Changelog.** Add a `## [vX.Y.Z]` entry to [`CHANGELOG.md`](./CHANGELOG.md)
   (kernel rev, platforms, any packaging/link notes).
5. **Commit + tag.** Commit, then push the path-prefixed module tags plus the
   root tag, all at that commit:
   ```bash
   for p in linux_amd64 linux_arm64 darwin_amd64 darwin_arm64 windows_amd64; do
     git tag "lib/$p/vX.Y.Z"
   done
   git tag "vX.Y.Z"
   git push origin main --tags
   ```
6. **Driver.** In `databricks-sql-go`, bump the five `go.mod` requires to
   `vX.Y.Z`, refresh `go.sum`, set `KERNEL_REV` to the `vX.Y.Z` commit, and sync
   the committed header. See that repo's `docs/RELEASING.md`.

## Immutability

Go module versions are **immutable** once a proxy has served them. Never move a
released tag — cut a new version instead. (While the repo is private and
un-proxied a tag can still be re-cut; once public, bump.)

## Adding a platform

Create `lib/<os>_<arch>/` with its own `go.mod`, a build-tag-gated `prebuilt.go`
(matching `//go:build` + `#cgo LDFLAGS`), and the committed archive; add a root
`prebuilt_<os>_<arch>.go` shim; add the platform to the release workflow's build
matrix and to the driver's `go.mod` + `cgo_<os>_<arch>.go` shim; tag it alongside
the others.
