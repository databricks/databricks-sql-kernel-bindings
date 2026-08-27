# databricks-sql-kernel-bindings

Prebuilt, per-platform static libraries of the **Databricks SQL kernel**
(`libdatabricks_sql_kernel`) plus the cgo binding layer, packaged as Go modules
so the [Databricks SQL Driver for Go](https://github.com/databricks/databricks-sql-go)
can link the kernel via `go get` with **no Rust toolchain and no build step**.

> You almost certainly do **not** depend on this repo directly. It is an
> implementation detail of `github.com/databricks/databricks-sql-go`, which
> requires the right platform module for you automatically when you build with
> `-tags databricks_kernel`.

> **Issues and feature requests:** GitHub issues are disabled here — this repo
> only ships prebuilt artifacts. Please file them in the driver repo instead:
> [databricks/databricks-sql-go/issues](https://github.com/databricks/databricks-sql-go/issues).

## Why this repo exists

The kernel is a Rust library exposed over a C ABI. Go's module system can only
deliver a prebuilt archive to a `go get` consumer if that archive is committed in
the git tree of a Go module (release assets and Git LFS are invisible to the
module proxy). To keep the driver repo free of large committed binaries, those
archives live here instead — one nested Go module per platform, so a build
downloads **only** the archive for the platform it targets.

## Layout

```
databricks-sql-kernel-bindings/
├── go.mod                         # root module: cgo binding layer + per-platform import shims
├── include/databricks_kernel.h    # committed C ABI header (platform-independent)
├── prebuilt_<os>_<arch>.go         # build-tag-gated: imports the matching lib/<platform> module
└── lib/
    └── <os>_<arch>/
        ├── go.mod                  # its OWN module: github.com/databricks/databricks-sql-kernel-bindings/lib/<os>_<arch>
        ├── prebuilt.go             # //go:build <os> && <arch>  + the #cgo LDFLAGS that link the archive
        └── libdatabricks_sql_kernel.a
```

Each `lib/<platform>` is a **separate Go module**, and its `prebuilt.go` is
build-tag-gated. A build for, say, `darwin/arm64` downloads only
`.../lib/darwin_arm64`; a pure-Go Thrift build of the driver downloads none of
these modules.

## Versioning

Modules are released with **path-prefixed tags** so each can be versioned
independently:

```
lib/darwin_arm64/v1.2.3
lib/linux_amd64/v1.2.3
v1.2.3                     # the root module
```

The driver's `go.mod` pins the versions it needs; upgrading the driver is what
moves the kernel version. Each bindings version matches the
`databricks-sql-kernel` release it is built from. See [`RELEASING.md`](./RELEASING.md)
for the release process and the full pin chain, [`CHANGELOG.md`](./CHANGELOG.md)
for per-version release notes, and the driver repo's `docs/RELEASING.md` for the
consumer side.

## Cloning (contributors)

This repo commits large binaries, so a full clone accumulates their history. Use
a partial clone to fetch blobs lazily:

```bash
git clone --filter=blob:none https://github.com/databricks/databricks-sql-kernel-bindings
```

## License

Apache 2.0 — see [LICENSE](./LICENSE) and [NOTICE](./NOTICE). The committed
archives statically include third-party open source components; see NOTICE.
