# Changelog — Databricks SQL Kernel Go bindings

All notable changes to the packaged Go bindings are documented here. The format
is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this
project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Release notes for the prebuilt, per-platform C-ABI static archives
(`libdatabricks_sql_kernel.a`) packaged here as Go modules and consumed by
[`databricks-sql-go`](https://github.com/databricks/databricks-sql-go).

Each release is **versioned to match the `databricks-sql-kernel` release it is
built from**, so the C-ABI and behaviour changes for a given version are the
kernel's — see the kernel's own changelog:
<https://github.com/databricks/databricks-sql-kernel/blob/main/CHANGELOG.md>.

This file records **binding-repo-specific** notes: which kernel revision each
release is built from, the platforms shipped, and any packaging / link changes.
Tags are path-prefixed per module (`lib/<platform>/vX.Y.Z`) plus a root
`vX.Y.Z`; see [`RELEASING.md`](./RELEASING.md).

## [Unreleased]

## [v1.0.0] - 2026-08-31

- Built from **`databricks-sql-kernel` commit `36e14e81`**, tagged `v1.0.0`.
  All archives were CI-built and trivy-scanned by
  [`peco-databricks-sql-kernel-go` run 33436087078](https://github.com/databricks/secure-public-registry-releases-eng/actions/runs/33436087078).
- Platforms (7): `darwin_amd64`, `darwin_arm64`, `linux_amd64`, `linux_arm64`,
  `linux_arm`, `windows_amd64`, `windows_arm64`. The workflow's cgo link checks
  passed for `linux_arm` and `windows_arm64`.

## [v0.3.0] - 2026-08-28

- Built from **`databricks-sql-kernel` commit `46fffe9e`**, tagged `v0.3.0`.
  All archives were CI-built and trivy-scanned by
  [`peco-databricks-sql-kernel-go` run 33220006730](https://github.com/databricks/secure-public-registry-releases-eng/actions/runs/33220006730).
- Platforms (7): `darwin_amd64`, `darwin_arm64`, `linux_amd64`, `linux_arm64`,
  `linux_arm`, `windows_amd64`, `windows_arm64`. The workflow's cgo link checks
  passed for `linux_arm` and `windows_arm64`.
- Remove the committed `include/databricks_kernel.h`. Nothing in this repo
  compiles against it — the `lib/<platform>` modules carry only the archive +
  `#cgo LDFLAGS`. The C-ABI header is owned by the driver's cgo layer
  (`databricks-sql-go`'s `internal/backend/kernel/include/`), which is the only
  place that `#include`s it. Drops a redundant copy; the header still ships in the
  release artifact for the driver to sync from. (Root-module-only change; the
  `lib/<platform>` modules and their tags are unaffected.)

## [v0.2.3] - 2026-08-27

- **CI-built (scanned) rebuild of v0.2.2** — identical kernel source
  (`databricks-sql-kernel` commit `dd810d6d`, tagged `v0.2.2`), rebuilt through
  Databricks' internal CI release pipeline so the archives are CI-built and
  trivy-scanned rather than cross-built on a laptop. **Supersedes v0.2.2**, which
  was a local dev build of the same source.
- No API/behaviour change vs v0.2.2 — same C ABI, same `dd810d6d` header. The
  binaries differ only in build provenance (Rust builds are not bit-reproducible).
- v0.2.2 remains published but is immutable and superseded; new consumers should
  pin **v0.2.3**. (v0.2.1/v0.2.2 cannot be re-pointed — Go module versions are
  immutable once in the checksum DB.)
- **Two platforms added at `v0.2.3`** (later, from the same `dd810d6d` source):
  `lib/linux_arm` (`armv7-unknown-linux-gnueabihf`) and `lib/windows_arm64`
  (`aarch64-pc-windows-gnullvm`, built with llvm-mingw — Go cgo links a GNU `.a`,
  not MSVC). These are **new module paths**, so they take their first tag at
  `v0.2.3` without re-tagging the original five. Both are CI-built + trivy-scanned
  (cross-compiled on the Linux runner) and their `#cgo LDFLAGS` are validated by
  an in-CI on-demand link-check against every exported `kernel_*` symbol:
  `linux_arm` → `-lm -ldl` (no `-lstdc++`; the archive has no C++), `windows_arm64`
  → `-lbcrypt -lntdll` (its other Windows imports are embedded via raw-dylib).
- Platforms (7): `darwin_amd64`, `darwin_arm64`, `linux_amd64`, `linux_arm64`,
  `linux_arm`, `windows_amd64`, `windows_arm64`.

## [v0.2.2] - 2026-08-27

- Built from **`databricks-sql-kernel` commit `dd810d6d`** — the post-`v0.2.1`
  revision the driver's `KERNEL_REV` points at. This is **not yet a tagged
  kernel release**; it carries three **additive** (backward-compatible) C-ABI
  commits on top of `v0.2.1`:
  - `feat(c-abi): expose mTLS client identity` — new
    `kernel_session_config_set_tls_client_certificate`.
  - `feat(bindings): expose request timeout to Python and C` — new
    `kernel_session_config_set_request_timeout`.
  - `fix(metadata): align getTypeInfo with Thrift`.
- The bump was required because `databricks-sql-go` `main` calls the two new
  symbols above, which are absent from the `v0.2.1` archives — so `v0.2.1`
  would fail to link. The C-ABI header is `dd810d6d`'s (a strict superset of
  `v0.2.1`'s: 19 lines added, none removed/changed).
- Platforms: `darwin_amd64`, `darwin_arm64`, `linux_amd64`, `linux_arm64`,
  `windows_amd64` (Windows is the **`-gnu`** / MinGW archive, for Go cgo).
- Provenance note: these archives were **cross-built locally** (not via
  Databricks' internal CI release pipeline, whose reviewer gate requires a
  human). Re-cutting via CI for scanned provenance is a recommended follow-up.

## [v0.2.1] - 2026-08-24

- Built from **`databricks-sql-kernel` v0.2.1**; `KERNEL_REV` in the driver
  points at the same commit and the committed `include/databricks_kernel.h`
  is that revision's header.
- Platforms: `darwin_amd64`, `darwin_arm64`, `linux_amd64`, `linux_arm64`,
  `windows_amd64` (Windows is the **`-gnu`** / MinGW archive, for Go cgo).
- Built + security-scanned by Databricks' internal CI release pipeline and
  published here.
- Kernel C-ABI changes picked up in this release — identity federation /
  mandatory OIDC token exchange, JWT private-key M2M, Azure SP M2M, generic
  OAuth token-endpoint / scopes overrides, `kernel_session_test()`, retry-config
  and CloudFetch-chunk-cap setters, the callback logging bridge, and `Session`
  transaction control — are documented in the kernel v0.2.1 changelog.
- darwin modules link the archive without `-Wl,-rpath` (rejected by cgo's default
  LDFLAGS allowlist, and meaningless for a static archive).
