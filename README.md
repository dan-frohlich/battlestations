# battlestations

Go tools for the **Battlestations** cooperative sci-fi tabletop RPG: model,
validate, and print character sheets from YAML.

## Packages

| Package                  | Purpose                                                             |
|--------------------------|-------------------------------------------------------------------|
| `character/model`        | character, gear, pools, rank, skills, species, special abilities, plus a validator (`NewCharacterValidator().ValidateAll`) |
| `character/print`        | render a character to **PDF** (`gofpdf`) or an ASCII sheet; small (A5) and large (Letter) templates with background art |
| `character`              | `Manager` — loads a model character and prints it                   |

## Commands

| Command             | Does                                                        |
|---------------------|-----------------------------------------------------------|
| `cmd/climanager`    | validate a character file and print an ASCII sheet          |
| `cmd/printsheet`    | render a character file to a PDF                            |
| `cmd/prob`          | dice-probability helper                                     |

The repo-root `main.go` is an older equivalent of `cmd/climanager`.

## Usage

```bash
go run ./cmd/climanager -print -file character/model/sample01.yaml
go run ./cmd/printsheet -file character/model/sample01.yaml
```

Sample character files: `character/model/sample0{1..4}.yaml`.

> The `Makefile` targets reference `./cmd/cli` / binary `bs-printer`, which
> don't exist — build the `cmd/*` packages directly.
