# gogeul

[![Go Reference](https://pkg.go.dev/badge/github.com/geul-org/gogeul.svg)](https://pkg.go.dev/github.com/geul-org/gogeul)

A Go implementation of GEUL (General Encoding Unified Language) and the `geul` command-line tool.

GEUL is an artificial language and binary format for exchanging information between people and AI without ambiguity. A GEUL document (`.geul`) is a sequence of records made of 16-bit words: References point to things, Values hold values, Edges make statements, and Meta records describe the document. Documents are published by an Issuer, and each published version is pinned by a content hash.

- Website: <https://geul.org>
- Grammar: <https://geul.org/grammar/> (Korean: <https://geul.org/ko/grammar/>)

## Install

CLI:

```sh
go install github.com/geul-org/gogeul/cmd/geul@latest
```

Library:

```sh
go get github.com/geul-org/gogeul
```

## Quick start

### CLI

`geul` works on a *data folder*: a directory that holds `.well-known/geul/`, laid out like the public addresses (`https://geul.org/.well-known/geul/...`). It is found from `--data`, then `GEULDATA`, then by walking up from the current directory (as git finds `.git`).

Browse an archive:

```sh
geul ls                                  # Issuers in the data folder
geul ls geul.org                         # documents of one Issuer
geul show geul.org/triple                # one document, latest version, one record per line
geul show geul.org/triple@2               # a given version
geul show ./d.geul                       # a local file
geul show versions geul.org/triple       # versions: number, time, content hash
```

Write and publish a document under a self-issued Issuer:

```sh
geul issuer new --name me                # Ed25519 key, kept outside the data folder
geul new me/0x0000000000000001 -o d.geul
geul add d.geul LABEL Q42 '"Douglas Adams"@en'
geul add d.geul TRIPLE P31 Q42 Q5 normal
geul check d.geul
geul publish d.geul --name adams
geul show me/adams
```

Codec names and fields (`LABEL`, `TRIPLE`, `normal`, ...) come from the codec spec documents of the GEUL Issuer (`0x0000`, geul.org) in the data folder. Without them, `add` cannot read such lines and `show` prints those records as raw words.

View it in a browser, or bake it for a static host:

```sh
geul serve                               # prints a local URL
geul build -o site                       # writes site/.well-known/geul/
```

Convert between the two storage formats:

```sh
geul pack build -o packed ./tree         # directory archive -> pack archive
geul pack unpack -o tree ./packed        # pack archive -> directory archive
geul pack verify
```

### Library

Check a document, decode it with the codec specs of a directory archive, and print it as `geul show` does:

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/geul-org/gogeul/pkg/inspect"
	"github.com/geul-org/gogeul/pkg/notation"
	"github.com/geul-org/gogeul/pkg/resolve"
	"github.com/geul-org/gogeul/pkg/schema"
)

func main() {
	b, err := os.ReadFile("d.geul")
	if err != nil {
		log.Fatal(err)
	}

	// Format check: framing, document structure, Kind rules.
	rep := schema.Check(b, schema.Options{})
	fmt.Println("pass:", rep.Verdict == schema.Pass, "facts:", len(rep.Facts))

	// Codec specs published in the archive give records their names and fields.
	spec, err := resolve.Dir{Root: "data"}.CodecSpec()
	if err != nil {
		log.Fatal(err)
	}
	doc, err := inspect.DecodeWith(b, spec)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range append(doc.Head, doc.Body...) {
		fmt.Println(notation.FormatEntry(e))
	}
}
```

`resolve.Dir{Root: ...}` reads a directory archive. For a pack archive, open it with `pack.OpenDir` and pass it as `Dir.Store`.

## Packages

| Package | Purpose |
|---|---|
| [`pkg/core`](pkg/core) | Engine-embedded rules: framing (TAG WORD, Length), record and document encoding, canonical form, content hash. No external dependencies, no network I/O. |
| [`pkg/codecspec`](pkg/codecspec) | Gathers codec spec documents (DECLARE, FIELD, ENUM, CODEC, BIND) into one field list per codec. |
| [`pkg/delta`](pkg/delta) | Version diffs (DELTA-BASE, DELTA-REMOVE): create and apply, list changes by Edge Id, blame, carry LOCAL IDs over to the next version. |
| [`pkg/inspect`](pkg/inspect) | Decodes a document into interpreted entries (spec name, content, Edge Id, references); content is decoded only where re-encoding gives back the original words. |
| [`pkg/notation`](pkg/notation) | The one-line text form of a record, used by `geul show`, `add` and `meta`; lines round-trip. |
| [`pkg/pack`](pkg/pack) | Pack archive format: immutable content-named objects plus `latest.txt`, read locally (pread) or over HTTP Range. |
| [`pkg/resolve`](pkg/resolve) | Archives (publication, version lists, names, deprecation) and IMPORT dependency resolution through an injected source list. Decoding never fetches. |
| [`pkg/schema`](pkg/schema) | Format checks: zero-layer rules (framing, document structure, Kind rules), then codec layouts. Results are Facts and a Verdict. |
| [`pkg/toulmin`](pkg/toulmin) | The TOULMIN codec: records of judgment and execution graphs, their code tables and values. |

The CLI lives in `cmd/geul` and `internal/`.

## CLI commands

| Group | Commands |
|---|---|
| Read and check | `ls`, `show`, `show versions`, `cat`, `hash`, `diff`, `blame`, `grep`, `check [--archive]`, `resolve` |
| Edit a draft file | `new`, `open`, `add`, `def`, `set`, `rm [--force]`, `meta [--unset]` |
| Publish | `publish [--amend] [--name]`, `name`, `delta`, `deprecate`, `tombstone` |
| Issuers and settings | `issuer new\|ls\|show`, `env [-w\|-u]` |
| Viewer | `serve`, `build` |
| Pack archives | `pack build [--single]`, `pack unpack`, `pack verify`, `pack compact` |

Changes are compared by Edge Id, not by line. Run `geul <command> --help` for details. Global flags: `--data`, `--config`, `--format text|json`.

## Archives and storage

A data folder's `.well-known/geul/` is stored in one of two formats. Both are read the same way (`resolve.Store`).

- **Directory archive**: one file per public path (`<Issuer>/<doc>/versions.geul`, `<Issuer>/<doc>/<version>.geul`, `h/<content hash>.geul`, name indexes). This is the layout a web host serves.
- **Pack archive**: immutable content-named objects (`m/<sha256>.gpm` manifests, `o/<sha256>.gpx` indexes, `o/<sha256>.gpk` data packs) and `latest.txt`, the only mutable file. There is one set per Issuer (`<Issuer>/latest.txt`), or one set at the top (`pack build --single`). The same objects can be read from a local folder or a CDN with one HTTP Range request per read.

A new, empty data folder is written as a pack archive; a directory archive that already has files stays a directory archive. Commands that write to a pack archive add a new generation per changed set after the command succeeds.

### Settings

Each setting is read from a command-line flag, then the environment variable, then the settings file (`GEULENV`, default `<user config dir>/geul/env`), then a default. `geul env` prints the values and where each came from; `geul env -w` and `-u` edit the settings file.

| Variable | Meaning |
|---|---|
| `GEULDATA` | Data folder (the directory holding `.well-known/geul/`). Default: found by walking up from the current directory. |
| `GEULENV` | Settings file. |
| `GEULFORMAT` | Output format: `text` (default) or `json`. |
| `GEULISSUER` | Issuer prefixed to the document ID given to `geul new` without one. |
| `GEULSTORE` | `dir` makes a new data folder a directory archive instead of a pack archive. |
| `GEULPROXY` | Ordered source list for resolving IMPORTs, separated by `,` or `\|`. Items: `https://...`, `file://<folder>`, `registry`, and `off` (stop here). A source that lacks a document passes to the next; an item followed by `\|` also passes on any other error. Default: `https://geul.org,registry`. A `file://` folder holding a pack archive is read as one. Plain `http://` is rejected. |
| `GEULPACK` | Pack sources read before `GEULPROXY`, comma-separated: `https://` set addresses, local set folders, or data folders holding a pack archive. |
| `GEULPRIVATE` | Comma-separated `<Issuer>/<doc>` patterns (`*` allowed) of documents never requested from geul.org or `registry`. |

`geul resolve` uses the data folder first, then `GEULPACK`, then `GEULPROXY`. Private keys of self-issued Issuers are kept in `<user config dir>/geul/keys/`, outside the data folder. `geul env` also lists names reserved for features not implemented yet (`GEULCACHE`, `GEULREMOTE`, `GEULLOG`, `GEULNOLOG`, `GEULINSECURE`, `GEULWORK`).

### Compression

Records in a pack archive are stored raw or as zstd frames ([klauspost/compress](https://github.com/klauspost/compress)). The write codec comes from `GEULCOMPRESS` (environment variable or settings file key); the `--compress` flag of `geul pack build` and `geul pack compact` takes precedence over it.

| Value | Meaning |
|---|---|
| `none` | Store records raw. |
| `zstd` | zstd at level 19, with a trained dictionary. |
| `zstd:LEVEL` | zstd at level 1–22. |
| `zstd:nodict`, `zstd:LEVEL:nodict` | zstd without a dictionary. |

- Unset, a new set is written raw, and an existing set keeps its codec: it is written with zstd if it already has a dictionary or a zstd record.
- Reading does not depend on the setting; compressed records are always read.
- klauspost/compress maps zstd levels to four speeds; 11 and above run at its best setting.
- Each set gets at most one dictionary (about 110 KiB), trained deterministically from sampled records (the same samples give the same bytes). It is kept only when it saves more than its own size; later generations reuse it, and `pack compact` trains a new one.
- Very large sets (Wikidata) are written in one pass with `pack.Stream`: records in key order go straight into data packs, P leaves are written as they fill, H keys are sorted in spilled runs and merged at the end, and `Checkpoint`/`ResumeStream` survive a crash at any point, including inside `Close`, with the same bytes as an uninterrupted run (`pkg/pack/README.md`). Memory does not grow with the number of runs.
- Up to 4 records are compressed at once. Inputs shorter than 256 bytes are not compressed.
- Measured on Wikidata Q1–Q10000, zstd stores 0.336 of the raw size with a dictionary and 0.393 without.

## Viewer

```
geul serve [--addr host:port] [FILE.geul|DIR]
geul build -o OUT [DIR]
```

Without an argument both use the data folder.

- `serve` listens on `127.0.0.1` with a free port unless `--addr` is given, and re-reads files on every request. Documents are served at their public paths under `/.well-known/geul/`. `<path>.geul` is always the bytes (`application/vnd.geul`); `<path>` is the bytes when `Accept` asks for `application/vnd.geul`, and the viewer page otherwise (`Vary: Accept`). Directories are listed, with `index.json` for the same listing as JSON. Name paths (`geul.org/triple`) redirect to ID paths.
- `build` writes the same pages as static files to `OUT/.well-known/geul/`: `<path>.geul`, the page at `<path>`, `index.html` and `index.json` per directory, and the scripts under `_static/`. The web host must serve `<path>.geul` for `<path>` when `Accept` asks for `application/vnd.geul`, and send `Vary: Accept`. `build` refuses an output inside the source tree and stops at a document that does not decode.
- The viewer makes no external network requests. The graph view embeds Cytoscape.js, dagre and cytoscape-dagre (MIT).

## Requirements

Go 1.25 or later (the version in [`go.mod`](go.mod)).

## Contributing

Issues and pull requests are welcome. Run `go build ./...`, `go vet ./...` and `go test ./...` before sending a change.

## License

The project's own work in this repository is dedicated to the public domain under [CC0 1.0](LICENSE). Third-party components (Go modules, embedded JavaScript) keep their own licenses; see [NOTICE](NOTICE).
