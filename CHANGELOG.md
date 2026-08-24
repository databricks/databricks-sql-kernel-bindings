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

## [v0.2.1] - 2026-08-24

- Built from **`databricks-sql-kernel` v0.2.1**; `KERNEL_REV` in the driver
  points at the same commit and the committed `include/databricks_kernel.h`
  is that revision's header.
- Platforms: `darwin_amd64`, `darwin_arm64`, `linux_amd64`, `linux_arm64`,
  `windows_amd64` (Windows is the **`-gnu`** / MinGW archive, for Go cgo).
- Built + security-scanned by the `peco-databricks-sql-kernel-go` release
  workflow (GitHub-hosted macOS/Windows runners) and published here.
- Kernel C-ABI changes picked up in this release — identity federation /
  mandatory OIDC token exchange, JWT private-key M2M, Azure SP M2M, generic
  OAuth token-endpoint / scopes overrides, `kernel_session_test()`, retry-config
  and CloudFetch-chunk-cap setters, the callback logging bridge, and `Session`
  transaction control — are documented in the kernel v0.2.1 changelog.
- darwin modules link the archive without `-Wl,-rpath` (rejected by cgo's default
  LDFLAGS allowlist, and meaningless for a static archive).
